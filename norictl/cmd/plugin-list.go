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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of plugins",
	Run: func(cmd *cobra.Command, args []string) {
		cli, closeCh := client.NewClient(viper.GetString("grpc-address"))

		reply, err := cli.ListCommand(context.Background(), &commands.ListRequest{})
		close(closeCh)
		if err != nil {
			logrus.Fatal(err)
		}

		for i, resp := range reply.Data {
			fmt.Printf(": %3d : %15s : %10s : %10s :\n", i+1, resp.Namespace, resp.Name, resp.Author)
		}
	},
}

func init() {
	pluginCmd.AddCommand(listCmd)
}
