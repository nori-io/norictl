package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/secure2work/nori/proto"
	"github.com/secure2work/norictl/client"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "downloading and installing plugin",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("plugin-path")
		if len(path) == 0 && len(args) > 0 {
			path = args[0]
		}

		client, closeCh := client.NewClient(
			viper.GetString("grpc-address"),
			viper.GetString("ca"),
			viper.GetString("ServerHostOverride"),
		)

		reply, err := client.GetCommand(context.Background(), &commands.GetRequest{
			Uri:                 path,
			InstallDependencies: viper.GetBool("dependencies"),
		})

		close(closeCh)

		if err != nil {
			logrus.Fatal(err)
			if reply != nil {
				logrus.Fatal(reply.Error)
			}
		} else {
			fmt.Printf("Plugin %q successfully installed\n", path)
		}
	},
}

func init() {
	pluginCmd.AddCommand(getCmd)
}
