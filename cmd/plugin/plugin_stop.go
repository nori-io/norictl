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

	"github.com/nori-io/nori-common/v2/version"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/client/utils"
	protoGenerated "github.com/nori-io/norictl/internal/generated/protobuf"
)

var (
	stopAll func() bool
)

func stopCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "stop [PLUGIN_ID] [OPTIONS]",
		Short: "Stop plugin's or plugins' execution",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsStop()
			conn, err := connection.CurrentConnection()
			if err != nil {
				fmt.Println("%s ", err)
				return
			}

			if len(args) == 0 {
				fmt.Println("PLUGIN_ID required!")
				return
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

			reply, err := client.PluginStopCommand(context.Background(), &protoGenerated.PluginStopRequest{
				Id: &protoGenerated.ID{
					Id:                   pluginIdSplit[0],
					Version:              pluginIdSplit[1],
				},
				FlagAll:              stopAll(),
			})
			defer close(closeCh)
			if err != nil {
				fmt.Println("%s", err)
				common.UI.PluginStopFailure(pluginId)
				if reply != nil {
					fmt.Println("%s", protoGenerated.ErrorReply{
						Status:               false,
						Error:                err.Error(),
					})
				}
			}

			common.UI.PluginStopFailure(pluginId)
		},
	}
}

func init() {
}

func setFlagsStop() {
	flags := utils.NewFlagBuilder(PluginCmd(), stopCmd())
	flags.Bool(&stopAll, "all", "--all", false, "Stop all plugins") // TODO
}
