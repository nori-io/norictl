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
	"github.com/nori-io/nori-common/v2/version"
	"github.com/nori-io/norictl/internal/errors"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"strings"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	protoGenerated "github.com/nori-io/norictl/pkg/proto"
)

var getCmd = &cobra.Command{
	Use:   "get [PLUGIN_ID] [OPTIONS]",
	Short: "downloading plugin",
	Long: `Get downloads the plugin, along with its dependencies.
	It then installs the plugin, like norictl plugin install.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := connection.CurrentConnection()
		if err != nil {
			fmt.Println("%s", err)
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

		flagVerbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer close(closeCh)
		reply, err := client.PluginGetCommand(context.Background(), &protoGenerated.PluginGetRequest{
			Id: &protoGenerated.ID{
				PluginId: pluginIdSplit[0],
				Version:  pluginIdSplit[1],
			},
			FlagVerbose: flagVerbose,
		})

		if err != nil {
			fmt.Println("%s", err)
			common.UI.PluginGetFailure(pluginId)
			if reply != nil {
				fmt.Println("%s", protoGenerated.Error{
					Code:    reply.GetMessage(),
					Message: reply.GetCode(),
				})
			}
			return
		} else {
			common.UI.PluginGetSuccess(pluginId)
		}
	},
}
