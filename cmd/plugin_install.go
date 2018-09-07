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

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install plugin",
	Run: func(cmd *cobra.Command, args []string) {
		id := viper.GetString("plugin-id")
		if len(id) == 0 && len(args) > 0 {
			id = args[0]
		}

		cli, closeCh := client.NewClient(
			viper.GetString("grpc-address"),
			viper.GetString("ca"),
			viper.GetString("ServerHostOverride"),
		)

		reply, err := cli.InstallCommand(context.Background(), &commands.InstallRequest{Id: id})
		defer close(closeCh)
		if err != nil {
			if reply != nil {
				logrus.Fatal(reply.Error)
			}
			logrus.Fatal(err)
		}

		fmt.Printf("Plugin %s installed, %3d :\n", id, resp.Int)
	},
}

func init() {
	pluginCmd.AddCommand(installCmd)
}
