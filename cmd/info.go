package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var infoCmd = NewServerCommand(
	"info",
	func(cmd *cobra.Command, args []string) {
		fmt.Printf("args: %#v", args)
	},
)

func init() {
	infoCmd.AppendTo(RootCmd)
}
