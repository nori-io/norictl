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
	//commands "github.com/nori-io/nori/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/client/utils"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
	"github.com/nori-io/norictl/internal/ui"
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
			FlagDeps:             pullDeps(),
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})

		close(closeCh)
		uiPull := ui.NewUI()
		if err != nil {
			log.Fatal(err)
			uiPull.PullFailure(pluginId)
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
			uiPull.PullFailure(pluginId)
		}
	},
}

func init() {
	PluginCmd.AddCommand(pullCmd)
	flags := utils.NewFlagBuilder(PluginCmd, pullCmd)
	flags.Bool(&pullDeps, "pull", "-d", false, "Pull downloads the plugin, with or without it's dependencies.")
}
