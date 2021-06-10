package grpc

import (
"fmt"
"net"
"time"

"github.com/spf13/pflag"
)

type GRPCServerOptions struct {
	BindAddress net.IP
	BindPort    int

	// Enable security
	TlsCertFile string
	TlsKeyFile  string
	EnableTLS   bool

	// Request timeout, count by seconds
	RequestTimeout int

	// Skip authentication and authorization check(only for local debug!!!)
	SkipAuthentication bool
	SkipAuthorization  bool
}

func NewGRPCServerOptions() *GRPCServerOptions {
	op := GRPCServerOptions{
		BindAddress:        net.ParseIP("0.0.0.0/24"),
		BindPort:           443,
		TlsCertFile:        "",
		TlsKeyFile:         "",
		RequestTimeout:     30,
		SkipAuthentication: false,
		SkipAuthorization:  false,
	}

	return &op
}

func (s *GRPCServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.IPVar(&s.BindAddress, "grpc-bind-ip", s.BindAddress, ""+
		"The IP address on which to listen for the api-routing server. This "+
		"address must be reachable by the rest of the cluster. If blank, "+
		"the host's default interface will be used.")

	fs.IntVar(&s.BindPort, "grpc-bind-port", s.BindPort, ""+
		"The bind port is for http server listen. If not set, the default port will be used.")

	fs.StringVar(&s.TlsCertFile, "grpc-tls-certificate", s.TlsCertFile,
		"Path to a cert file for grpc TLS.")

	fs.StringVar(&s.TlsKeyFile, "grpc-tls-key", s.TlsKeyFile,
		"Path to a key file for grpc TLS.")

	fs.BoolVar(&s.EnableTLS, "grpc-tls-enable", s.EnableTLS, "Enable grpc tls support")

	fs.IntVar(&s.RequestTimeout, "grpc-request-timeout", s.RequestTimeout, ""+
		"An optional field indicating the duration a handler must keep a request open before timing "+
		"it out. This is the default request timeout for requests")

	fs.BoolVar(
		&s.SkipAuthorization, "grpc-skip-authorization",
		s.SkipAuthorization, "skip authorization check for grpc server(set as True only for local debug)")
	fs.BoolVar(
		&s.SkipAuthentication, "grpc-skip-authentication",
		s.SkipAuthentication, "skip authorization check for grpc server(set as True only for local debug)")
}

func (s *GRPCServerOptions) Validate() []error {
	var errs []error
	if s.BindPort > 65535 || s.BindPort < 1 {
		errs = append(errs, fmt.Errorf("--grpc-bind-port should be in the range [1-65535]: %d", s.BindPort))
	}

	if s.EnableTLS && (len(s.TlsKeyFile) == 0 || len(s.TlsCertFile) == 0) {
		errs = append(errs, fmt.Errorf("when enable tls, --grpc-tls-certificate/--grpc-tls-key cannot be empty"))
	}

	if s.RequestTimeout < 0 {
		errs = append(errs, fmt.Errorf("--grpc-request-timeout cannot be negative: %d", s.RequestTimeout))
	}

	return errs
}

func (s *GRPCServerOptions) ApplyTo(c *Config) {
	c.IP = s.BindAddress
	c.Port = s.BindPort
	c.RequestTimeout = time.Duration(s.RequestTimeout) * time.Second
	c.CertPath = s.TlsCertFile
	c.KeyPath = s.TlsKeyFile
	c.EnableTLS = s.EnableTLS
	c.SkipAuthorization = s.SkipAuthorization
	c.SkipAuthentication = s.SkipAuthentication
}

