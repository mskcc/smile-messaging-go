module github.com/smile_messaging_go

go 1.17

replace github.com/mskcc/smile_messaging_go => ../../cmo/n1zea_smile_messaging_go

require (
	github.com/mskcc/smile_messaging_go v0.0.0-00010101000000-000000000000
	github.com/nats-io/nats.go v1.16.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/nats-io/nats-server/v2 v2.8.4 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
