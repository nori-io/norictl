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
	metaDeps            func() bool
	metaDepsStatus      func() bool
	metaDependent       func() bool
	metaDependentStatus func() bool
)

func metaCmd(log logger.FieldLogger) *cobra.Command {

	return &cobra.Command{
		Use:   "meta [PLUGIN_ID] [OPTIONS]",
		Short: "Show plugin meta data.",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsMeta(log)
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

			meta := &protoNori.PluginMetaRequest{
				ID: &protoNori.PluginID{
					MetaId:               pluginId,
				},
				FlagDeps:             metaDeps(),
				FlagDepsStatus:       metaDepsStatus(),
				FlagDependent:        metaDependent(),
				FlagDependentStatus:  metaDependentStatus(),
			}

			reply, err := client.PluginMetaCommand(context.Background(), meta)
			defer close(closeCh)
			if err != nil {
				log.Error("%s", err)
			}
			common.UI.PluginMetaExist(fmt.Sprintf("%s", reply))
		},
	}
}

func init() {
}

func setFlagsMeta(log logger.FieldLogger) {
	flags := utils.NewFlagBuilder(PluginCmd(log), metaCmd(log))
	flags.Bool(&metaDeps, "deps", "--deps", false, "Show only plugin dependencies")
	flags.Bool(&metaDepsStatus, "deps-status", "--deps-status", false, "Show plugin dependencies with dependent plugin status (downloaded, installed, not found etc, with errors, running, installable,inactive)")
	flags.Bool(&metaDependent, "dependent", "--dependent", false, "Show only plugins, that depend on specified plugin")
	flags.Bool(&metaDependentStatus, "dependent-status", "--dependent-status", false, "Show plugins, that depend on specified plugin with their status (downloaded, installed, not found etc, with errors, running, installable,inactive)")
}
