package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/secure2work/nori/proto"
	"github.com/secure2work/norictl/client"
	"github.com/olekukonko/tablewriter"
	"os"
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

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "ID", "Name", "Author"})

		var i int
		for _, v := range reply.Data {
			i += 1
			table.Append([]string{
				string(i), v.Id, v.Name, v.Author,
			})
		}
		table.Render()
	},
}

func init() {
	pluginCmd.AddCommand(listCmd)
}
