package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/secure2work/nori/proto"
	"github.com/secure2work/norictl/client"
	"github.com/fzzy/radix/redis/resp"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "install plugin",
	Run: func(cmd *cobra.Command, args []string) {
		id := viper.GetString("id")
		if len(id) == 0 && len(args) > 0 {
			id = args[0]
		}

		cli, closeCh := client.NewClient(
			viper.GetString("grpc-address"),
			viper.GetString("ca"),
			viper.GetString("ServerHostOverride"),
		)

		reply, err := cli.PluginUninstallCommand(context.Background(), &commands.PluginUninstallRequest{Id: id})
		defer close(closeCh)
		if err != nil {
			if reply != nil {
				logrus.Fatal(reply.Error)
			}
			logrus.Fatal(err)
		}

		fmt.Printf("Plugin %s uninstalled, %3d :\n", id, resp.Int)
	},
}

func init() {
	pluginCmd.AddCommand(uninstallCmd)
}

