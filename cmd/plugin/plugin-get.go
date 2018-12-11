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

	"github.com/secure2work/nori/proto"
	"github.com/secure2work/norictl/client"
	"github.com/secure2work/norictl/client/connection"
	"github.com/secure2work/norictl/client/utils"
)

var (
	getDownload func() bool
	getUpdate   func() bool
)

var getCmd = &cobra.Command{
	Use:   "get [OPTIONS] PLUGIN_ID",
	Short: "downloading plugin",
	Long: `Get downloads the plugin, along with its dependencies.
	It then installs the plugin, like norictl plugin install.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := connection.CurrentConnection()
		if err != nil {
			log.Fatal(err)
		}

		if len(args) == 0 {
			log.Fatal("PLUGIN_ID required!")
		}

		pluginId := args[0]

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)

		reply, err := client.PluginGetCommand(context.Background(), &commands.PluginGetRequest{
			Uri:                 pluginId,
			InstallDependencies: true,
		})

		close(closeCh)

		if err != nil {
			log.Fatal(err)
			if reply != nil {
				log.Fatal(reply.Error)
			}
		} else {
			fmt.Printf("Plugin %q successfully installed\n", pluginId)
		}
	},
}

func init() {
	flags := utils.NewFlagBuilder(PluginCmd, getCmd)
	flags.Bool(&getDownload, "download", "d", false, "Stop after downloading the plugin, do not install it")
	flags.Bool(&getUpdate, "update", "u", false, "Use the network to update plugin and plugin dependencies")
}
