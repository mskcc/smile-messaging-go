package nats

import (
	smsg "github.com/mskcc/smile-messaging-go/messaging"
	"github.com/nats-io/nats.go"
)

// This implements smile_messaging_go.Messaging
type Messaging struct {
	nc *nats.Conn
	js nats.JetStream
}

func NewMessaging(url string, opts ...Option) (*Messaging, error) {
	var options Options
	for _, o := range opts {
		o(&options)
	}

	var nc *nats.Conn
	var err error
	if options.UseTLS {
		cert := nats.ClientCert(options.TLSCertPath, options.TLSKeyPath)
		nc, err = nats.Connect(url, cert, nats.UserInfo(options.UserId, options.Password))
		if err != nil {
			return nil, err
		}
	} else {
		nc, err = nats.Connect(url)
		if err != nil {
			return nil, err
		}
	}
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	return &Messaging{
		nc: nc,
		js: js,
	}, nil
}

func (m *Messaging) Publish(subj string, data []byte) error {
	msg := &nats.Msg{Subject: subj, Data: data}
	msg.Header = make(map[string][]string)
	msg.Header.Add("Nats-Msg-Subject", subj)

	_, err := m.js.PublishMsg(msg)
	if err != nil {
		return err
	}
	select {
	case <-m.js.PublishAsyncComplete():
	}
	return err
}

func (m *Messaging) Subscribe(con string, subj string, mh smsg.MsgHandler) error {

	// lets create a nats messsage handler
	// that passes the nats message content
	// to the smile message handler
	nmh := func(m *nats.Msg) {
		sm := &smsg.Msg{
			Subject: m.Subject,
			Data:    m.Data,
		}
		mh(sm)
	}

	// subscribe to the Nats subject & register the nats message handler
	_, err := m.js.Subscribe(subj, nmh, nats.Durable(con))
	return err
}

func (m *Messaging) Shutdown() {
	m.nc.Flush()
	m.nc.Close()
	m.nc = nil
	m.js = nil
}
