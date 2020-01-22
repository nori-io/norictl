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
	"strings"

	"github.com/nori-io/nori-common/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/client/utils"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var (
	startAll func() bool
)

var startCmd = &cobra.Command{
	Use:   "norictl plugin start [PLUGIN_ID] [OPTIONS]",
	Short: "Start one plugin or all plugins.",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := connection.CurrentConnection()
		if err != nil {
			log.Fatal(err)
		}

		if len(args) == 0 {
			log.Fatal("PLUGIN_ID required!")
		}

		pluginId := args[0]

		pluginIdSplit := strings.Split(pluginId, ":")
		versionPlugin := pluginIdSplit[1]
		_, err= version.NewVersion(versionPlugin)
		if err != nil {
			fmt.Println("Format of plugin's version is incorrect:", err)
		}

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)

		reply, err := client.PluginStartCommand(context.Background(), &protoNori.PluginStartRequest{
			Id: &protoNori.ID{
				Id:                   pluginId,
				Version:              "",
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			FlagAll:              startAll(),
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})
		defer close(closeCh)
		if err != nil {
			UI.StartFailure(pluginId)
			if reply != nil {
				log.Fatal(protoNori.ErrorReply{
					Status:               false,
					Error:                err.Error(),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				})
			}
			log.Fatal(err)
		}
		UI.StartSuccess(pluginId)

	},
}

func init() {
	PluginCmd.AddCommand(startCmd)
	flags := utils.NewFlagBuilder(PluginCmd, startCmd)
	flags.Bool(&startAll, "all", "--all", false, "Start all plugins") // TODO

}
