package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show Info.",
	Long:  `Show Info.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Using config file: " + viper.ConfigFileUsed())
		fmt.Println(viper.Get("CUSTOMER.ENDPOINT"))
		fmt.Println(viper.Get("metrics"))
	},
}

func init() {
	RootCmd.AddCommand(infoCmd)
}
