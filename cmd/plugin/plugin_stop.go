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
	"github.com/nori-io/norictl/internal/client/utils"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var (
	stopAll func() bool
)

func stopCmd(log logger.Logger) *cobra.Command {

	return &cobra.Command{
		Use:   "norictl plugin stop [PLUGIN_ID] [OPTIONS]",
		Short: "Stop plugin's or plugins' execution",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsStop(log)
			conn, err := connection.CurrentConnection()
			if err != nil {
				log.Fatal("%s ", err)
			}

			if len(args) == 0 {
				log.Fatal("PLUGIN_ID required!")
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

			reply, err := client.PluginStopCommand(context.Background(), &protoNori.PluginStopRequest{
				Id: &protoNori.ID{
					Id:                   pluginIdSplit[0],
					Version:              pluginIdSplit[1],
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				FlagAll:              stopAll(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			})
			defer close(closeCh)
			if err != nil {
				common.UI.StopFailure(pluginId)
				if reply != nil {
					log.Fatal("%s", protoNori.ErrorReply{
						Status:               false,
						Error:                err.Error(),
						XXX_NoUnkeyedLiteral: struct{}{},
						XXX_unrecognized:     nil,
						XXX_sizecache:        0,
					})
				}
				log.Fatal("%s", err)
			}

			common.UI.StopFailure(pluginId)
		},
	}
}

func init() {
}

func setFlagsStop(log logger.Logger) {
	flags := utils.NewFlagBuilder(PluginCmd(log), stopCmd(log))
	flags.Bool(&stopAll, "all", "--all", false, "Stop all plugins") // TODO
}
