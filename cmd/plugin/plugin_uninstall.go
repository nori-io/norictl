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
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/errors"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/nori-grpc/pkg/api/proto"
	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
)

var (
	uninstallAll       bool
	uninstallDependent bool
)

var uninstallCmd = &cobra.Command{

	Use:   "uninstall [PLUGIN_ID] [OPTIONS]",
	Short: "Uninstall plugin or plugins.",
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := connection.CurrentConnection()
		if err != nil {
			fmt.Println(err)
			return
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
		/* @todo versionPlugin := pluginIdSplit[1]
		_, err = version.NewVersion(versionPlugin)
		if err != nil {
			errors.ErrorFormatPluginVersion(err)
			return
		}*/

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)



		reply, err := client.PluginUninstall(context.Background(), &proto.PluginUninstallRequest{
			Id: &proto.ID{
				PluginId: pluginIdSplit[0],
				Version:  pluginIdSplit[1],
			},
			FlagAll:       uninstallAll,
			FlagDependent: uninstallDependent,
		})
		defer close(closeCh)
		if (err != nil) || (reply.Error.GetCode() != "") {
			if err != nil {
				fmt.Println(err)
			}
			if reply.Error.GetCode() != "" {
				fmt.Println(proto.Error{
					Code:    reply.Error.GetCode(),
					Message: reply.Error.GetMessage(),
				})
			}
			common.UI.PluginGetFailure(pluginId)
			return
		}
		common.UI.PluginUninstallSuccess(pluginId)

	},
}

func init() {
	uninstallCmd.Flags().BoolVarP(&uninstallAll, "all", "a", false, "Uninstall all installed plugins")
	uninstallCmd.Flags().BoolVarP(&uninstallDependent, "dependent", "d", false, "Uninstall plugin and depend plugins")
}
