package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"net"
	"strconv"
)

type ServerCommand struct {
	name string
	cmd  *cobra.Command
}

func NewServerCommand(
	name string,
	run func(cmd *cobra.Command, args []string),
) *ServerCommand {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Show Info.",
		Long:  `Show Info.`,
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, append(args, GetEndpoint(name)))
		},
	}

	return &ServerCommand{
		name: name,
		cmd:  cmd,
	}
}

func (c *ServerCommand) AppendTo(root *cobra.Command) {
	var (
		bind string
		port int
	)

	c.cmd.Flags().StringVarP(&bind, "bind", "", "127.0.0.1", "interface to which the Info server will bind")
	c.cmd.Flags().IntVarP(&port, "port", "p", 50052, "port on which the Info info will listen")

	c.SetOption("info", "bind")
	c.SetOption("info", "port")

	root.AddCommand(c.cmd)
}

func (c *ServerCommand) SetOption(prefix, key string) {
	opt := prefix + "." + key
	viper.BindEnv(opt)
	viper.BindPFlag(opt, c.cmd.Flags().Lookup(key))
}

func GetEndpoint(prefix string) string {
	bind := viper.Get(prefix + ".bind")
	port := viper.Get(prefix + ".port")
	return net.JoinHostPort(bind.(string), strconv.Itoa(port.(int)))
}
