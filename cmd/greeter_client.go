package cmd

import (
	"net"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/hg2c/hellogrpc/greeter"
)

// greeterClientCmd represents the greeter command
var greeterClientCmd = &cobra.Command{
	Use:   "greeter_client",
	Short: "Starts Greeter service",
	Long:  `Starts Greeter service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := greeter.NewClient(
			"greeter_client",
			net.JoinHostPort(greeterClientOptions.serverInterface, strconv.Itoa(greeterClientOptions.serverPort)),
		)
		defer client.Close()

		client.Hello("tao")
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
