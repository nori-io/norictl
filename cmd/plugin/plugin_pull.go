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

	//commands "github.com/nori-io/nori/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/client"
	"github.com/nori-io/norictl/client/connection"
	"github.com/nori-io/norictl/client/utils"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var (
	pullDeps func() bool
)

var pullCmd = &cobra.Command{
	Use:   "norictl plugin pull [PLUGIN_ID] [OPTIONS]",
	Short: "downloading plugin",
	Long:  `Pull downloads the plugin, with or without it's dependencies.`,
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

		reply, err := client.PluginPullCommand(context.Background(), &protoNori.PluginPullRequest{
			Id: &protoNori.ID{
				Id:                   pluginId,
				Version:              "",
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			FlagDeps:             false,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})

		close(closeCh)

		if err != nil {
			log.Fatal(err)
			if reply != nil {
				log.Fatal(protoNori.ErrorReply{
					Status:               false,
					Error:                err.Error(),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				})
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
