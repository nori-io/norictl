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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of plugins",
	Run: func(cmd *cobra.Command, args []string) {
		cli, closeCh := client.NewClient(
			viper.GetString("grpc-address"),
			viper.GetString("ca"),
			viper.GetString("ServerHostOverride"),
		)

		reply, err := cli.ListCommand(context.Background(), &commands.ListRequest{})
		close(closeCh)
		if err != nil {
			if reply != nil {
				logrus.Fatal(reply.Error)
			}
			logrus.Fatal(err)
		}

		for i, resp := range reply.Data {
			fmt.Printf(": %3d : %15s : %10s : %10s :\n", i+1, resp.Id, resp.Name, resp.Author)
		}
	},
}

func init() {
	pluginCmd.AddCommand(listCmd)
}
