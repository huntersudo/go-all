// @owner:沈若禺
// @dest: sfc核心组件，各模块通信的API层
package main

import (
	"os"

	//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/cmd/sfc-apirouting/app"
	//_ "gerrit.cmss.com/BC-PaaS/ecloud-sfc/cmd/sfc-apirouting/docs"
	"pflag-test/app"
)

// @title ESFC-API
// @version 1.0
// @description  API for ESFC
// @contact.name 移动云-云原生团队
// @license.name 1.0
// @BasePath /api/sfc/apirouting/

func main() {
	command := app.NewAPIRoutingCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

/*
./pflag-test.exe \
--http-bind-ip=0.0.0.0 \
--http-bind-port=11433 \
--https-tls-certificate=./certs/server.crt \
--https-tls-key=./certs/server.key \
--https-enable=true \
--http-request-timeout=30 \
--grpc-bind-ip=0.1.1.1 \
--grpc-bind-port=11434  \
--grpc-tls-certificate=./certs/server.crt \
--grpc-tls-key=./certs/server.key \
 --grpc-tls-enable=true \
 --grpc-request-timeout=30 \
 --shutdown-delay-seconds=10

 */