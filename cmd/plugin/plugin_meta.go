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
	metaDeps       func() bool
	metaDepsStatus func() bool
)

var metaCmd = &cobra.Command{
	Use:   "meta [OPTIONS] PLUGIN_ID",
	Short: "Show plugin meta data.",
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

		meta := &commands.PluginMetaRequest{
			Id:                 pluginId,
			Dependencies:       metaDeps(),
			DependenciesStatus: metaDepsStatus(),
		}

		reply, err := cli.PluginMetaCommand(context.Background(), meta)
		defer close(closeCh)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Meta info for plugin \"%s\":\n%s\n", pluginId, reply.Json)
	},
}

func init() {
	flags := utils.NewFlagBuilder(PluginCmd, metaCmd)
	flags.Bool(&metaDeps, "deps", "d", false, "Show only plugin dependencies")
	flags.Bool(&metaDepsStatus, "deps-status", "s", false, "Show plugin dependencies with dependent plugin status (downloaded, installed, not found etc)")
}
