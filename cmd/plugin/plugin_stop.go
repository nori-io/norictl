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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/client/utils"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var (
	stopAll func() bool
)

var stopCmd = &cobra.Command{
	Use:   "norictl plugin stop [PLUGIN_ID] [OPTIONS]",
	Short: "Stop plugin's or plugins' execution",
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

		reply, err := client.PluginStopCommand(context.Background(), &protoNori.PluginStopRequest{
			Id: &protoNori.ID{
				Id:                   pluginId,
				Version:              "",
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			FlagAll:              false,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})
		defer close(closeCh)
		if err != nil {
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

		fmt.Printf("Plugin %s stopped, %3d :\n", pluginId, resp.Int)
	},
}

func init() {
	PluginCmd.AddCommand(startCmd)
	flags := utils.NewFlagBuilder(PluginCmd, startCmd)
	flags.Bool(&stopAll, "all", "--all", false, "Stop all plugins") // TODO

}
