package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hg2c/hellogrpc/greeter"
	"github.com/hg2c/hellogrpc/rpc"
)

// greeterClientCmd represents the greeter command
var greeterClientCmd = &cobra.Command{
	Use:   "greeter_client",
	Short: "Starts Greeter service",
	Long:  `Starts Greeter service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := greeter.NewClient(
			"greeter_client",
			greeterClientOptions.Host,
			greeterClientOptions.Port,
		)
		defer client.Close()

		client.Hello("tao")
		return nil
	},
}

var (
	greeterClientOptions rpc.ListenAddress
)

func init() {
	RootCmd.AddCommand(greeterClientCmd)

	greeterClientCmd.Flags().StringVarP(&greeterClientOptions.Host, "bind", "", "127.0.0.1", "interface to which the Greeter server will bind")
	greeterClientCmd.Flags().IntVarP(&greeterClientOptions.Port, "port", "p", 50051, "port on which the Greeter server will listen")
}
