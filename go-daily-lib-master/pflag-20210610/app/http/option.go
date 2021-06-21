package http


import (
"fmt"
"net"
"time"


"github.com/spf13/pflag"
)

type HTTPServerOptions struct {
	BindAddress net.IP
	BindPort    int

	// List of allowed origins for CORS
	CorsAllowedOrigins []string

	// Enable security
	TlsCertFile string
	TlsKeyFile  string
	TlsCaFile   string
	EnableHttps bool
	Tracing     bool

	// Request timeout, count by seconds
	RequestTimeout int

	// Skip authentication and authorization check(only for local debug!!!)
	SkipAuthentication bool
	SkipAuthorization  bool
}

func NewHTTPServerOptions() *HTTPServerOptions {
	op := HTTPServerOptions{
		BindAddress:        net.ParseIP("0.0.0.0/24"),
		BindPort:           443,
		CorsAllowedOrigins: nil,
		TlsCertFile:        "",
		TlsKeyFile:         "",
		RequestTimeout:     30,
		SkipAuthentication: false,
		SkipAuthorization:  false,
	}

	return &op
}

func (s *HTTPServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.IPVar(&s.BindAddress, "http-bind-ip", s.BindAddress, ""+
		"The IP address on which to listen for the api-routing server. This "+
		"address must be reachable by the rest of the cluster. If blank, "+
		"the host's default interface will be used.")

	fs.IntVar(&s.BindPort, "http-bind-port", s.BindPort, ""+
		"The bind port is for http server listen. If not set, the default port will be used.")

	fs.StringSliceVar(&s.CorsAllowedOrigins, "http-cors-allowed-origins", s.CorsAllowedOrigins, ""+
		"List of allowed origins for CORS, comma separated.  An allowed origin can be a regular "+
		"expression to support subdomain matching. If this list is empty CORS will not be enabled.")

	fs.StringVar(&s.TlsCertFile, "https-tls-certificate", s.TlsCertFile,
		"Path to a cert file for https TLS.")

	fs.StringVar(&s.TlsKeyFile, "https-tls-key", s.TlsKeyFile,
		"Path to a key file for https TLS.")

	fs.StringVar(&s.TlsCaFile, "https-ca-key", s.TlsCaFile, "Path to a ca file for https TLS.")

	fs.BoolVar(&s.EnableHttps, "https-enable", s.EnableHttps, "Enable https tls support")

	fs.IntVar(&s.RequestTimeout, "http-request-timeout", s.RequestTimeout, ""+
		"An optional field indicating the duration a handler must keep a request open before timing "+
		"it out. This is the default request timeout for requests")

	fs.BoolVar(
		&s.SkipAuthorization, "http-skip-authorization",
		s.SkipAuthorization, "skip authorization check for http server(set as True only for local debug)")
	fs.BoolVar(
		&s.SkipAuthentication, "http-skip-authentication",
		s.SkipAuthentication, "skip authorization check for http server(set as True only for local debug)")
	fs.BoolVar(&s.Tracing, "tracing", false, "Enable tracing")
}

func (s *HTTPServerOptions) Validate() []error {
	var errs []error
	if s.BindPort > 65535 || s.BindPort < 1 {
		errs = append(errs, fmt.Errorf("--http-bind-port should be in the range [1-65535]: %d", s.BindPort))
	}

	if s.EnableHttps && (len(s.TlsKeyFile) == 0 || len(s.TlsCertFile) == 0) {
		errs = append(errs, fmt.Errorf("when enable https, --https-tls-certificate/--https-tls-key cannot be empty"))
	}

	if s.RequestTimeout < 0 {
		errs = append(errs, fmt.Errorf("--http-request-timeout cannot be negative: %d", s.RequestTimeout))
	}

	return errs
}

func (s *HTTPServerOptions) ApplyTo(c *Config) {
	c.IP = s.BindAddress
	c.Port = s.BindPort
	c.RequestTimeout = time.Duration(s.RequestTimeout) * time.Second
	c.CorsAllowedOriginList = s.CorsAllowedOrigins
	c.CertPath = s.TlsCertFile
	c.KeyPath = s.TlsKeyFile
	c.EnableHttps = s.EnableHttps
	c.CaPath = s.TlsCaFile
	c.SkipAuthentication = s.SkipAuthentication
	c.SkipAuthorization = s.SkipAuthorization
}

