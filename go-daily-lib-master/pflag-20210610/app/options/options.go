package options

import (
	"fmt"
	grpc2 "pflag-test/app/grpc"
	http2 "pflag-test/app/http"
	utils2 "pflag-test/app/utils"
)

type ServerRunOptions struct {
	// http server options
	HTTPServerOptions *http2.HTTPServerOptions

	// grpc server options
	GRPCServerOptions *grpc2.GRPCServerOptions


	// Delay after server shutdown, count by seconds
	ShutdownDelay int
}

// NewServerRunOptions creates a ServerRunOptions
func NewServerRunOptions() *ServerRunOptions {
	op := ServerRunOptions{
		HTTPServerOptions: http2.NewHTTPServerOptions(),
		GRPCServerOptions: grpc2.NewGRPCServerOptions(),
		ShutdownDelay:     5,
	}

	return &op
}

// Flags returns flags for a specific APIRouting by section name
func (s *ServerRunOptions) Flags() *utils2.NamedFlagSets {
	flagSets := &utils2.NamedFlagSets{}
	s.HTTPServerOptions.AddFlags(flagSets.FlagSet("http"))
	s.GRPCServerOptions.AddFlags(flagSets.FlagSet("grpc"))
	//s.CommunicationOptions.AddFlags(flagSets.FlagSet("communication"))
	//s.DatabaseConfig.AddFlags(flagSets.FlagSet("database"))

	fs := flagSets.FlagSet("common")
	fs.IntVar(&s.ShutdownDelay, "shutdown-delay-seconds", s.ShutdownDelay, ""+
		"Time to delay the termination. During that time the server keeps serving requests normally and /healthz "+
		"returns success, but /readyz immediately returns failure. Graceful termination starts after this delay "+
		"has elapsed. This can be used to allow load balancer to stop sending traffic to this server.")

	return flagSets
}

// Validate will check the options validation
func (s *ServerRunOptions) Validate() []error {
	var errs []error
	if s.ShutdownDelay > 60 || s.ShutdownDelay < 0 {
		errs = append(errs, fmt.Errorf("--shutdown-delay-seconds should be in the range [0-60]: %d", s.ShutdownDelay))
	}

	errs = append(errs, s.HTTPServerOptions.Validate()...)
	errs = append(errs, s.GRPCServerOptions.Validate()...)
	//errs = append(errs, s.CommunicationOptions.Validate()...)
	//errs = append(errs, s.DatabaseConfig.Validate()...)
	return errs
}

//// ApplyToHTTPServerConfig convert options to the HTTPServerConfig
//func (s *ServerRunOptions) ApplyToHTTPServerConfig(middlewareConfig server.MiddlewareConfig) *http.Config {
//	communicationConfig := communication.NewCommunicationConfig()
//	s.CommunicationOptions.ApplyTo(&communicationConfig)
//	c := http.NewConfig(middlewareConfig, communicationConfig)
//	s.HTTPServerOptions.ApplyTo(c)
//	c.ShutdownDelayDuration = time.Duration(s.ShutdownDelay) * time.Second
//
//	return c
//}
//
//// ApplyToGRPCServerConfig convert options to the GRPCServerConfig
//func (s *ServerRunOptions) ApplyToGRPCServerConfig(middlewareConfig server.MiddlewareConfig) *grpc.Config {
//	communicationConfig := communication.NewCommunicationConfig()
//	s.CommunicationOptions.ApplyTo(&communicationConfig)
//	c := grpc.NewConfig(middlewareConfig, communicationConfig)
//	s.GRPCServerOptions.ApplyTo(c)
//	c.ShutdownDelayDuration = time.Duration(s.ShutdownDelay) * time.Second
//
//	return c
//}
