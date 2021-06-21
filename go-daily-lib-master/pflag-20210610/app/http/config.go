package http

import (
"net"
"time"

//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/apirouting/server"
//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/apirouting/server/http/filters"
//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/apirouting/service/communication"
//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/utils"
//
//"github.com/gin-gonic/gin"
)

// Config is a structure used to configure a APIServer.
type Config struct {
	IP                    net.IP
	Port                  int
	RequestTimeout        time.Duration
	//HandlerChainWaitGroup *utils.SafeWaitGroup
	ShutdownDelayDuration time.Duration
	CorsAllowedOriginList []string

	//BuildHandlerChainFunc func(*gin.Engine, *Config)
	//MiddlewareConfig      server.MiddlewareConfig
	//CommunicationConfig   communication.Config

	EnableHttps bool
	CertPath    string
	KeyPath     string
	CaPath      string

	// Skip authentication and authorization check(only for local debug!!!)
	SkipAuthentication bool
	SkipAuthorization  bool
}
//
//// NewConfig returns a Config struct with the default values
//func NewConfig(middlewareConfig server.MiddlewareConfig, communicationConfig communication.Config) *Config {
//	return &Config{
//		IP:                    []byte("0.0.0.0"),
//		Port:                  443,
//		RequestTimeout:        time.Duration(60) * time.Second,
//		HandlerChainWaitGroup: new(utils.SafeWaitGroup),
//		ShutdownDelayDuration: time.Duration(5) * time.Second,
//
//		BuildHandlerChainFunc: DefaultBuildHandlerChain,
//		MiddlewareConfig:      middlewareConfig,
//		CommunicationConfig:   communicationConfig,
//	}
//}
//
//// DefaultBuildHandlerChain defines the http request handler chain
//func DefaultBuildHandlerChain(engine *gin.Engine, c *Config) {
//	engine.Use(filters.WithPanicRecovery())
//	engine.Use(filters.OpenTracing())
//	engine.Use(filters.WithRequestInfo(c.MiddlewareConfig.RequestInfoRegister))
//	if !c.SkipAuthentication {
//		engine.Use(filters.WithAuthentication(c.MiddlewareConfig.AuthenticatorRegister))
//	}
//	engine.Use(filters.WithWaitGroup(c.HandlerChainWaitGroup))
//	engine.Use(filters.WithTimeout(c.RequestTimeout))
//	engine.Use(filters.WithCORS(
//		c.CorsAllowedOriginList, nil, nil, nil, "true"))
//	if !c.SkipAuthorization {
//		engine.Use(filters.WithAuthorization(c.MiddlewareConfig.Authorizator))
//	}
//	engine.Use(filters.WithMaxInFlightLimit())
//	engine.Use(filters.WithAdmission(c.MiddlewareConfig.AdmissionRegister))
//}

