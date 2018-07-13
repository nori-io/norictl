package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/secure2work/nori-cli/norictl/client"
	"github.com/secure2work/nori-cli/proto"
)

var installCmd = &cobra.Command{
	Use:   "get",
	Short: "downloading and installing plugin",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("plugin-path")
		if len(path) == 0 && len(args) > 0 {
			path = args[0]
		}

		client, closeCh := client.NewClient(viper.GetString("grpc-address"))

		reply, err := client.GetCommand(context.Background(), &commands.GetRequest{
			Uri:                 path,
			InstallDependencies: viper.GetBool("dep-install"),
		})

		close(closeCh)

		if err != nil {
			logrus.Fatal(err, reply.Error)
		} else {
			fmt.Printf("Plugin %q successfully installed\n", path)
		}
	},
}

func init() {
	pluginCmd.AddCommand(installCmd)

	installCmd.Flags().Bool("dep-install", true, "install dependencies")
	viper.BindPFlag("dep-install", installCmd.Flags().Lookup("dep-install"))
}
