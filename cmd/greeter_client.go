package cmd

import (
	"net"
	"strconv"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing"

	"github.com/hg2c/hellogrpc/greeter"
)

// greeterClientCmd represents the greeter command
var greeterClientCmd = &cobra.Command{
	Use:   "greeter_client",
	Short: "Starts Greeter service",
	Long:  `Starts Greeter service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := log.NewFactory(logger.With(zap.String("client", "greeter_client")))
		client := greeter.NewClient(
			net.JoinHostPort(greeterClientOptions.serverInterface, strconv.Itoa(greeterClientOptions.serverPort)),
			tracing.Init("greeter", metricsFactory.Namespace("greeter", nil), logger),
			logger,
		)
		client.Hello("luo")
		return nil
	},
}

var (
	greeterClientOptions struct {
		serverInterface string
		serverPort      int
	}
)

func init() {
	RootCmd.AddCommand(greeterClientCmd)

	greeterClientCmd.Flags().StringVarP(&greeterClientOptions.serverInterface, "bind", "", "127.0.0.1", "interface to which the Greeter server will bind")
	greeterClientCmd.Flags().IntVarP(&greeterClientOptions.serverPort, "port", "p", 50051, "port on which the Greeter server will listen")
}
