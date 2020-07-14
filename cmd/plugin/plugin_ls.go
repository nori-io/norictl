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
	"github.com/nori-io/norictl/internal/client/utils"
	protoGenerated "github.com/nori-io/norictl/pkg/proto"
)

var (
	listError       func() bool
	listInstallable func() bool
	listInstalled   func() bool
	listRunning     func() bool
	listStopped    func() bool

)

func lsCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "ls [OPTIONS]",
		Aliases: []string{"list"},
		Short:   "Shows list of plugins on remote Nori node.",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsLs()
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
				FlagError:       listError(),
				FlagInstalled:   listInstalled(),
				FlagRunning:     listRunning(),
				FlagInstallable: listInstallable(),
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

			filter := func(list []*protoGenerated.Plugin, f func(p protoGenerated.Plugin) bool) []*protoGenerated.Plugin{
				newList := make([]*protoGenerated.Plugin, 0)
				plugins := make([][]string, len(list))
				for _, l := range list {
					if f(*l) {
						newList = append(newList, l)
						var  licenses, dependecies string
						for _, license:=range l.Meta.Licenses{
							licenses=licenses+license.String()+"\n"
						}
						for _, dependency:=range l.Meta.Dependencies{
							dependecies=dependecies+dependency.String()+"\n"
						}
						plugins = append(plugins, []string{l.Meta.Id.String(), l.Meta.Author.String(),l.Meta.Interface, licenses, dependecies})
					}
				}
				common.UI.PluginsList(plugins)

				return newList
			}

			if listError() {
				list = filter(list, func(p protoGenerated.Plugin) bool {
					return p.Status.Enum().String()=="WithError"
				})
			}

/*			if listInactive() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})
					}
					common.UI.PluginsInactive(plugins)
					return p.FlagInactive
				})
			}
*/
/*			if listInstallable() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})

					}
					common.UI.PluginsInstallable(plugins)
					return p.FlagInstallable
				})
			}
*/
/*			if listInstalled() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})

					}
					common.UI.PluginsInstalled(plugins)
					return p.FlagInstalled
				})
			}
*/
/*			if listRunning() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})
					}
					common.UI.PluginsRunning(plugins)
					return p.FlagRunning
				})
			}
*/		},
	}
}
func setFlagsLs() {
	flags := utils.NewFlagBuilder(PluginCmd(), lsCmd())
	flags.Bool(&listError, "error", "e", true, "Show plugins with errors (not implement)")                 // TODO
	flags.Bool(&listInstallable, "installable", "", true, "Show plugins that need to install") // TODO
	flags.Bool(&listInstalled, "installed", "i", true, "Show only installed plugins")
	flags.Bool(&listRunning, "running", "r", true, "Show only running plugins")
	flags.Bool(&listRunning, "stopped", "s", true, "Show plugins that are not running")

}
