package nats

// Options for Messaging provider
type Options struct {
	UseTLS      bool
	TLSCertPath string
	TLSKeyPath  string
	UserId      string
	Password    string
}

// An Option is a function operating on the Messaging Options
type Option func(*Options)

// WithTLS is an Option to enable a TLS channel
func WithTLS(certPath, keyPath, userId, pw string) Option {
	return func(o *Options) {
		o.UseTLS = true
		o.TLSCertPath = certPath
		o.TLSKeyPath = keyPath
		o.UserId = userId
		o.Password = pw
	}
}
