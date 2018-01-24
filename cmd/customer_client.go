package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hg2c/hellogrpc/customer"
	"github.com/hwgo/pher/wgrpc"
)

// customerClientCmd represents the customer command
var customerClientCmd = &cobra.Command{
	Use:   "customer_client",
	Short: "Starts Customer service",
	Long:  `Starts Customer service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := customer.NewClient2(
			"customer_client",
			customerClientOptions.Host,
			customerClientOptions.Port,
		)
		defer client.Close()

		client.Get()
		return nil
	},
}

var (
	customerClientOptions wgrpc.ListenAddress
)

func init() {
	RootCmd.AddCommand(customerClientCmd)

	customerClientCmd.Flags().StringVarP(&customerClientOptions.Host, "bind", "", "127.0.0.1", "interface to which the Customer server will bind")
	customerClientCmd.Flags().IntVarP(&customerClientOptions.Port, "port", "p", 50052, "port on which the Customer server will listen")
}
