package cmd

import (
	"net"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/hg2c/hellogrpc/greeter"
)

// greeterCmd represents the greeter command
var greeterCmd = &cobra.Command{
	Use:   "greeter",
	Short: "Starts Greeter service",
	Long:  `Starts Greeter service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server := greeter.NewServer(
			"greeter",
			net.JoinHostPort(greeterOptions.serverInterface, strconv.Itoa(greeterOptions.serverPort)),
		)
		return server.Run()
	},
}

var (
	greeterOptions struct {
		serverInterface string
		serverPort      int
	}
)

func init() {
	RootCmd.AddCommand(greeterCmd)

	greeterCmd.Flags().StringVarP(&greeterOptions.serverInterface, "bind", "", "127.0.0.1", "interface to which the Greeter server will bind")
	greeterCmd.Flags().IntVarP(&greeterOptions.serverPort, "port", "p", 50051, "port on which the Greeter server will listen")
}
