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
	metaDeps            bool
	metaDepsStatus      bool
	metaDependent       bool
	metaDependentStatus bool
)

var metaCmd = &cobra.Command{

	Use:   "meta [PLUGIN_ID] [OPTIONS]",
	Short: "Show plugin meta data.",
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := connection.CurrentConnection()
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) == 0 {
			fmt.Println("PLUGIN_ID required!")
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
		defer close(closeCh)

		meta := &protoGenerated.PluginMetaRequest{
			Id: &protoGenerated.ID{
				PluginId: pluginId,
				Version:  versionPlugin,
			},
			FlagDeps:            metaDeps,
			FlagDepsStatus:      metaDepsStatus,
			FlagDependent:       metaDependent,
			FlagDependentStatus: metaDependentStatus,
		}

		reply, err := client.PluginMeta(context.Background(), meta)
		if err != nil {
			fmt.Println(err)
			return
		}
		common.UI.PluginMetaExist(fmt.Sprintf("%s", reply))

	},
}

func init() {
	metaCmd.Flags().BoolVarP(&metaDeps, "deps", "d", false, "Show only plugin dependencies")
	metaCmd.Flags().BoolVarP(&metaDepsStatus, "deps-status", "", false, "Show plugin dependencies with dependent plugin status (downloaded, installed, not found etc, with errors, running, installable,inactive)")
	metaCmd.Flags().BoolVarP(&metaDependent, "dependent", "", false, "Show only plugins, that depend on specified plugin")
	metaCmd.Flags().BoolVarP(&metaDependentStatus, "dependent-status", "", false, "Show plugins, that depend on specified plugin with their status (downloaded, installed, not found etc, with errors, running, installable,inactive)")
}
