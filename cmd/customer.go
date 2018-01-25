package cmd

import (
	"net"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hg2c/hellogrpc/customer"
)

// customerCmd represents the customer command
var customerCmd = &cobra.Command{
	Use:   "customer",
	Short: "Starts Customer service",
	Long:  `Starts Customer service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server := customer.NewServer(
			"customer",
			net.JoinHostPort(customerOptions.serverInterface, strconv.Itoa(customerOptions.serverPort)),
		)
		return server.Run()
	},
}

var (
	customerOptions struct {
		serverInterface string
		serverPort      int
	}
)

func init() {
	RootCmd.AddCommand(customerCmd)
	viper.BindEnv("customer.endpoint")
	customerCmd.Flags().StringVarP(&customerOptions.serverInterface, "bind", "", "127.0.0.1", "interface to which the Customer server will bind")
	customerCmd.Flags().IntVarP(&customerOptions.serverPort, "port", "p", 50052, "port on which the Customer server will listen")
}
