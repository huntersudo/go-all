package grpc

import (
	"net"
	"time"
)



// Config is a structure used to configure a grpc APIServer.
type Config struct {
	IP   net.IP
	Port int

	RequestTimeout        time.Duration
	//HandlerChainWaitGroup *utils.SafeWaitGroup
	ShutdownDelayDuration time.Duration

	//MiddlewareConfig        server.MiddlewareConfig
	//CommunicationConfig     communication.Config
	//BuildUnaryHandlerChain  func(*Config) []grpc.UnaryServerInterceptor
	//BuildStreamHandlerChain func(*Config) []grpc.StreamServerInterceptor

	EnableTLS bool
	CertPath  string
	KeyPath   string

	// Skip authentication and authorization check(only for local debug!!!)
	SkipAuthentication bool
	SkipAuthorization  bool
}

//// NewConfig returns a Config struct with the default values
//func NewConfig(middlewareConfig server.MiddlewareConfig, communicationConfig communication.Config) *Config {
//	return &Config{
//		IP:                    net.ParseIP("0.0.0.0"),
//		Port:                  10443,
//		RequestTimeout:        time.Duration(60) * time.Second,
//		HandlerChainWaitGroup: new(utils.SafeWaitGroup),
//		ShutdownDelayDuration: time.Duration(5) * time.Second,
//
//		BuildUnaryHandlerChain:  DefaultBuildUnaryHandlerChain,
//		BuildStreamHandlerChain: DefaultBuildStreamHandlerChain,
//		MiddlewareConfig:        middlewareConfig,
//		CommunicationConfig:     communicationConfig,
//	}
//}
//
//func DefaultBuildUnaryHandlerChain(c *Config) []grpc.UnaryServerInterceptor {
//	return []grpc.UnaryServerInterceptor{
//		interceptor.PanicUnaryServerInterceptor(),
//		interceptor.RequestInfoUnaryServerInterceptor(),
//		interceptor.AuthenticationUnaryServerInterceptor(c.MiddlewareConfig.AuthenticatorRegister, c.SkipAuthentication),
//		interceptor.WaitGroupUnaryServerInterceptor(c.HandlerChainWaitGroup),
//		interceptor.TimeoutUnaryServerInterceptor(c.RequestTimeout),
//		interceptor.AuthorizationUnaryServerInterceptor(c.MiddlewareConfig.Authorizator, c.SkipAuthorization),
//		interceptor.MaxInFlightLimitUnaryServerInterceptor(),
//		interceptor.AdmissionUnaryServerInterceptor(c.MiddlewareConfig.AdmissionRegister),
//	}
//}
//
//func DefaultBuildStreamHandlerChain(c *Config) []grpc.StreamServerInterceptor {
//	return []grpc.StreamServerInterceptor{
//		interceptor.PanicStreamServerInterceptor(),
//		interceptor.RequestInfoStreamServerInterceptor(),
//		interceptor.AuthenticationStreamServerInterceptor(c.MiddlewareConfig.AuthenticatorRegister, c.SkipAuthentication),
//		interceptor.WaitGroupStreamServerInterceptor(c.HandlerChainWaitGroup),
//		interceptor.TimeoutStreamServerInterceptor(c.RequestTimeout),
//		interceptor.AuthorizationStreamServerInterceptor(c.MiddlewareConfig.Authorizator, c.SkipAuthorization),
//		interceptor.MaxInFlightLimitStreamServerInterceptor(),
//		interceptor.AdmissionStreamServerInterceptor(c.MiddlewareConfig.AdmissionRegister),
//	}
//}
//
