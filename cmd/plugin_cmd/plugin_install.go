// Copyright © 2018 Secure2Work info@secure2work.com
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

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/fzzy/radix/redis/resp"
	"github.com/secure2work/nori/proto"
	"github.com/secure2work/norictl/client"
)

var installCmd = &cobra.Command{
	Use:   "install",
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

		reply, err := cli.PluginInstallCommand(context.Background(), &commands.PluginInstallRequest{Id: id})
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
	PluginCmd.AddCommand(installCmd)
}
