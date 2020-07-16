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
	"github.com/nori-io/norictl/internal/errors"
	"strings"

	"github.com/nori-io/nori-common/v2/version"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	protoGenerated "github.com/nori-io/norictl/pkg/proto"
)

var (
	installAll      bool
)

func installCmd() *cobra.Command {

	cmd:= &cobra.Command{
		Use:   "install [PLUGIN_ID] [OPTIONS]",
		Short: "Install downloaded plugin or plugins.",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := connection.CurrentConnection()
			if err != nil {
				fmt.Println("%s", err)
			}

			if len(args) == 0 {
				errors.ErrorEmptyPluginId()
				return
			}

			pluginId := args[0]
			pluginIdSplit := strings.Split(pluginId, ":")
			if len(pluginIdSplit) != 2 {
				errors.ErrorFormatPluginId()
				return
			}
			versionPlugin := pluginIdSplit[1]
			_, err = version.NewVersion(versionPlugin)
			if err != nil {
				errors.ErrorFormatPluginVersion(err)
				return
			}

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)
			defer close(closeCh)

			reply, err := client.PluginInstallCommand(context.Background(), &protoGenerated.PluginInstallRequest{
				Id: &protoGenerated.ID{
					PluginId: pluginIdSplit[0],
					Version:  pluginIdSplit[1],
				},
				FlagAll:     installAll,
			})

			if err != nil {
				fmt.Println("%s", err)
				if reply != nil {
					fmt.Println("%s", protoGenerated.Error{
						Code:    reply.GetCode(),
						Message: reply.GetMessage(),
					})
				}
				common.UI.PluginInstallFailure(pluginId)
			}
			common.UI.PluginInstallSuccess(pluginId)
		},
	}
	cmd.Flags().BoolVarP(&installAll, "--all", "a", true, "Install all installable plugins")
	return cmd
}


