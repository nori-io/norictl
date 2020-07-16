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
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	protoGenerated "github.com/nori-io/norictl/pkg/proto"
)

var (
	listError       bool
	listInstallable bool
	listInstalled   bool
	listRunning     bool
	listStopped     bool
)

var lsCmd = &cobra.Command{

	Use:     "ls [OPTIONS]",
	Aliases: []string{"list"},
	Short:   "Shows list of plugins on remote Nori node.",
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Flags().BoolVarP(&listError, "error", "e", false, "Show plugins with errors (not implement)")
		cmd.Flags().BoolVarP(&listInstallable, "installable", "", false, "Show plugins that need to install") // TODO
		cmd.Flags().BoolVarP(&listInstalled, "installed", "i", false, "Show only installed plugins")
		cmd.Flags().BoolVarP(&listRunning, "running", "r", false, "Show only running plugins")
		cmd.Flags().BoolVarP(&listStopped, "stopped", "s", false, "Show plugins that are not running")

		conn, err := connection.CurrentConnection()
		if err != nil {
			fmt.Println("%s", err)
		}

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)
		defer close(closeCh)

		reply, err := client.PluginListCommand(context.Background(), &protoGenerated.PluginListRequest{
			FlagError:       listError,
			FlagInstalled:   listInstalled,
			FlagInstallable: listInstallable,
			FlagRunning:     listRunning,
			FlagStopped:     listStopped,
		})

		if err != nil {
			if reply != nil {
				fmt.Println("%s", protoGenerated.Error{
					Code:    reply.Error.GetCode(),
					Message: reply.Error.GetMessage(),
				})
			}
			fmt.Println("%s", err)
		}

		list := reply.Plugin
		var plugins [][]string
		for _, l := range list {
			var licenses, dependecies string
			for _, license := range l.Meta.Licenses {
				licenses = licenses + license.String() + "\n"
			}
			for _, dependency := range l.Meta.Dependencies {
				dependecies = dependecies + dependency.String() + "\n"
			}
			plugins = append(plugins, []string{l.Meta.Id.PluginId + ":" + l.Meta.Id.Version, l.Meta.Author.String(), l.Meta.Interface, licenses, dependecies})
		}
		common.UI.PluginsList(plugins)

	},
}
