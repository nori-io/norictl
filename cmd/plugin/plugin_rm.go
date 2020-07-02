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

// Package plugin_cmd implements commands for work with plugins
//by command prompt*/
package plugin_cmd

import (
	"fmt"
	"strings"

	"github.com/nori-io/nori-common/v2/logger"
	"github.com/nori-io/nori-common/version"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	commonProtoGenerated "github.com/nori-io/norictl/internal/generated/protobuf/common"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

func rmCmd(log logger.FieldLogger) *cobra.Command {

	return &cobra.Command{
		Use:   "rm [PLUGIN_ID] [OPTIONS]",
		Short: "remove plugin",
		Long:  `Remove plugin from local machine.`,
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := connection.CurrentConnection()
			if err != nil {
				log.Error("%s", err)
			}

			if len(args) == 0 {
				log.Error("PLUGIN_ID required!")
			}

			pluginId := args[0]

			pluginIdSplit := strings.Split(pluginId, ":")
			versionPlugin := pluginIdSplit[1]
			_, err = version.NewVersion(versionPlugin)
			if err != nil {
				fmt.Println("Format of plugin's version is incorrect:", err)
			}

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)

			reply, err := client.PluginRemoveCommand(context.Background(), &protoNori.PluginRemoveRequest{
				Id: &commonProtoGenerated.ID{
					Id:                   pluginIdSplit[0],
					Version:              pluginIdSplit[1],
				},
			})

			close(closeCh)
			if err != nil {
				common.UI.PluginRmFailure(pluginId)
				log.Error("%s", err)
				if reply != nil {
					log.Error("%s", commonProtoGenerated.ErrorReply{
						Status:               false,
						Error:                err.Error(),
					})
				}
			} else {
				common.UI.PluginRmSuccess(pluginId)
			}
		},
	}
}

func init() {
}
