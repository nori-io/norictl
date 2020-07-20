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
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	protoGenerated "github.com/nori-io/norictl/pkg/proto"
)

var interfaceCmd = &cobra.Command{

	Use:   "interface [InterfaceName]",
	Short: "Shows list of plugins that implement specify interface.",
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := connection.CurrentConnection()
		if err != nil {
			fmt.Println(err)
			return
		}

		interfaceName := args[0]

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)
		defer close(closeCh)

		reply, err := client.PluginInterface(context.Background(), &protoGenerated.PluginInterfaceRequest{
			Interface: interfaceName,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		if reply.Plugin == nil {
			fmt.Println("Plugins not found")
			return
		} else {
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
		}
	},
}
