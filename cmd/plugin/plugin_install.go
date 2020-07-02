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
	commonProtoGenerated "github.com/nori-io/norictl/internal/generated/protobuf/common"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var (
	installVerbose func() bool
	installDeps    func() bool
	installAll     func() bool
)

func installCmd(log logger.Logger) *cobra.Command {

	return &cobra.Command{
		Use:   "norictl plugin install [PLUGIN_ID] [OPTIONS]",
		Short: "Install downloaded plugin or plugins.",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsInstall(log)
			conn, err := connection.CurrentConnection()
			if err != nil {
				log.Fatal("%s", err)
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

			reply, err := client.PluginInstallCommand(context.Background(), &protoNori.PluginInstallRequest{
				Id: &commonProtoGenerated.ID{
					Id:                   pluginIdSplit[0],
					Version:              pluginIdSplit[1],
				},
				FlagVerbose:          installVerbose(),
				FlagDeps:             installDeps(),
				FlagAll:              installAll(),
			})
			defer close(closeCh)
			if err != nil {
				log.Fatal("%s", err)
				if reply != nil {
					log.Fatal("%s", commonProtoGenerated.ErrorReply{
						Status:               false,
						Error:                err.Error(),
					})
				}
				common.UI.PluginInstallFailure(pluginId)
			}
			common.UI.PluginInstallSuccess(pluginId)
		},
	}
}

func setFlagsInstall(log logger.Logger) {
	flags := utils.NewFlagBuilder(PluginCmd(log), installCmd(log))
	flags.Bool(&installVerbose, "--verbose", "-v", false, "Verbose progress and debug output")
	flags.Bool(&installDeps, "--deps", "-d", false, "Install plugin with dependencies")
	flags.Bool(&installAll, "--all", "-all", false, "Install all installable plugins")
}

func init() {
}
