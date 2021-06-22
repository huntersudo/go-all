package app

import (
	"fmt"
	options2 "pflag-test/app/options"
	utils2 "pflag-test/app/utils"
	//"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	//"github.com/uber/jaeger-client-go"
	//"github.com/uber/jaeger-client-go/log"
	//"net"
	"os"
	//"sync"
	//"time"

	//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/cmd/sfc-apirouting/app/options"
	//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/apirouting/server"
	//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/apirouting/server/grpc"
	//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/apirouting/server/http"
	//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/apirouting/storage/persistence"
	//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/logger"
	//"gerrit.cmss.com/BC-PaaS/ecloud-sfc/pkg/utils"
	//
	//"github.com/opentracing/opentracing-go"
	//"github.com/spf13/cobra"
	//"github.com/uber/jaeger-client-go"
	//jaegerconfig "github.com/uber/jaeger-client-go/config"
	utilerrors "k8s.io--/apimachinery/pkg/util/errors"
)

// NewAPIRoutingCommand creates a *cobra.Command object with default parameters
func NewAPIRoutingCommand() *cobra.Command {
	serverOptions := options2.NewServerRunOptions()
	//var logConfig *logger.LogConfig
	cmd := &cobra.Command{
		Use: "esfc-apirouting",
		Long: `The API Routing services REST/GRPC operations and
provides the frontend to the third part's call'`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// validate log config
			//if errs := logConfig.Validate(); len(errs) != 0 {
			//	return utilerrors.NewAggregate(errs)
			//}
			//
			//if err := logger.InitialLogger(logConfig); err != nil {
			//	return fmt.Errorf("logger init error: %s\n", err)
			//}

			PrintFlags(cmd.Flags())

			// validate options
			if errs := serverOptions.Validate(); len(errs) != 0 {
				return utilerrors.NewAggregate(errs)
			}

			return Run(serverOptions, SetupSignalHandler())
		},
	}

	fs := cmd.Flags()
	namedFlagSets := serverOptions.Flags()
	//logConfig = logger.AddFlags(namedFlagSets.FlagSet("log"))
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	fmt.Println(serverOptions.HTTPServerOptions)
	fmt.Println("============")

	fmt.Println(serverOptions.GRPCServerOptions)


	usageFmt := "Usage:\n  %s\n"
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		utils2.PrintSections(cmd.OutOrStderr(), *namedFlagSets, 4)
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		utils2.PrintSections(cmd.OutOrStdout(), *namedFlagSets, 4)
	})

	return cmd
}

func Run(serverOptions *options2.ServerRunOptions, stopCh <-chan struct{}) error {
	fmt.Println("API Routing starting...")

	//tableStore, err := persistence.NewTableStore(*serverOptions.DatabaseConfig)
	//if err != nil {
	//	return err
	//}
	//
	//var middlewareConfig server.MiddlewareConfig
	//middlewareConfig, err = server.NewMiddlewareConfig(tableStore)
	//if err != nil {
	//	return err
	//}

	//hc := serverOptions.ApplyToHTTPServerConfig(middlewareConfig)
	//fmt.Println(hc)
	//httpServer := http.NewAPIServer(hc, tableStore)

	//gc := serverOptions.ApplyToGRPCServerConfig(middlewareConfig)
	//fmt.Println(gc)
	//grpcServer := grpc.NewAPIServer(gc, tableStore)

	if serverOptions.HTTPServerOptions.Tracing {
		// 本地测试
		// 启动命令行配置--tracing true，默认为false
		// os.Setenv("JAEGER_AGENT_HOST", "10.142.114.35")
		// os.Setenv("JAEGER_AGENT_PORT", "31155")
		// 设置环境变量JAEGER_AGENT_HOST与JAEGER_AGENT_PORT
		host := os.Getenv("JAEGER_AGENT_HOST")
		port := os.Getenv("JAEGER_AGENT_PORT")
		if host == "" || port == "" {
			host = "sfc-simplest-agent.observability"
			port = "6831"
		}
		//cfg := jaegerconfig.Configuration{
		//	ServiceName: "sfc-apirouting",
		//	Sampler: &jaegerconfig.SamplerConfig{
		//		// SamplingServerURL: "http://10.142.114.35:32083/sampling",
		//		Type:  jaeger.SamplerTypeConst,
		//		Param: 1,
		//	},
		//	Reporter: &jaegerconfig.ReporterConfig{
		//		LogSpans:            true,
		//		BufferFlushInterval: 1 * time.Second,
		//		LocalAgentHostPort:  net.JoinHostPort(host, port),
		//	},
		//}
		// 以下参数后续调整
		// jLogger := jaegerlog.StdLogger
		// jMetricsFactory := metrics.NullFactory
		// 本地测试，可开启日志
		//tracer, closer, err := cfg.NewTracer(
		// jaegerconfig.Logger(jLogger),
		// jaegerconfig.Metrics(jMetricsFactory)
		//)
		//if err != nil {
		//	return fmt.Errorf("failed to instantiate Jaeger tracer with %v", err)
		//}
		//opentracing.SetGlobalTracer(tracer)
		//defer closer.Close()
	}

	var errors []error
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	var err error
	//	httpServer, err = httpServer.PreRun()
	//	if err != nil {
	//		errors = append(errors, err)
	//		return
	//	}
	//
	//	err = httpServer.Run(stopCh)
	//	if err != nil {
	//		errors = append(errors, err)
	//	}
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	var err error
	//	grpcServer, err = grpcServer.PreRun()
	//	if err != nil {
	//		errors = append(errors, err)
	//		return
	//	}
	//
	//	err = grpcServer.Run(stopCh)
	//	if err != nil {
	//		errors = append(errors, err)
	//	}
	//}()
	//
	//wg.Wait()
	return utilerrors.NewAggregate(errors)
}


// PrintFlags logs the flags in the flagset
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
		fmt.Println()
	})
}