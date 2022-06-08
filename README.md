# smile_messaging_go

A wrapper around [NATS - Go Client](https://github.com/nats-io/nats.go).

## Installation

```bash
go get github.com/mskcc/smile_messaging_go
go get github.com/mskcc/smile_messaging_go/mom/nats
```

## Basic Usage

```go

import (
	"github.com/mskcc/smile_messaging_go"
	"github.com/mskcc/smile_messaging_go/mom/nats"
)

// without TLS
m, err := nats.NewMessage("localhost:4222")

// with TLS (see https://github.com/mskcc/smile_messaging_go/mom/nats/options.go)
m, err := nats.NewMessaging("localhost:4222", nats.WithTLS(certPath, keyPath, userId, pw))
if err != nil {
	// do something
}

// publish a message (see https://github.com/mskcc/smile_messaging_go/messaging.go)
err = m.Publish("foo subject", []byte("Hello World"))
if err != nil {
	// do something	
}

// subscribe to a subject(see https://github.com/mskcc/smile_messaging_go/messaging.go)
// consumer id much match an authorized id setup in NATS/Jetstream configuration 
m.Subscribe("consumer id", "foo subject", func(m *smile_messaging_go.Msg) {
	fmt.Println("Subscriber received an message via NATS:", string(m.Data))
})

m.Shutdown()
```
