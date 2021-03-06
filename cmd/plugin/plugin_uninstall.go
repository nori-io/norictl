// Copyright © 2018 Nori info@nori.io
//
// This program is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package plugin_cmd

import (
	"fmt"

	"github.com/fzzy/radix/redis/resp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/nori-io/nori/proto"
	"github.com/nori-io/norictl/client"
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
	PluginCmd.AddCommand(uninstallCmd)
}
