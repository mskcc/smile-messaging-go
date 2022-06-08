package smile_messaging_go

// The type passed into a message handler
// to the Subscribe routine when a message is
// received on the subject
type Msg struct {
	Subject string
	Data    []byte
}

// Clients to subscribe "register"
// a MsgHandler to process the message
type MsgHandler func(msg *Msg)

// An implementor of Messaging provides
// publish and subscribe capabilities
type Messaging interface {
	// publish a message to the given subject
	Publish(subj string, data []byte) error
	// the given consumer (con) subscribes to the given subject with the given message handler
	Subscribe(con string, subj string, mh MsgHandler) error
	// Release message broker resources
	Shutdown()
}
