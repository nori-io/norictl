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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/fzzy/radix/redis/resp"
	"github.com/secure2work/nori/proto"
	"github.com/secure2work/norictl/client"
	"github.com/secure2work/norictl/client/connection"
)

var installCmd = &cobra.Command{
	Use:   "install [OPTIONS] PLUGIN_ID",
	Short: "Install downloaded plugin.",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := connection.CurrentConnection()
		if err != nil {
			log.Fatal(err)
		}

		if len(args) == 0 {
			log.Fatal("PLUGIN_ID required!")
		}

		pluginId := args[0]

		cli, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)

		reply, err := cli.PluginInstallCommand(context.Background(), &commands.PluginInstallRequest{Id: pluginId})
		defer close(closeCh)
		if err != nil {
			if reply != nil {
				log.Fatal(reply.Error)
			}
			log.Fatal(err)
		}

		fmt.Printf("Plugin %s installed, %3d :\n", pluginId, resp.Int)
	},
}

func init() {
	PluginCmd.AddCommand(installCmd)
}
