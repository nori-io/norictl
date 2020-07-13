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

func interfaceCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "interface [InterfaceName]",
		Short: "Shows list of plugins that implement specify interface.",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsGet()
			conn, err := connection.CurrentConnection()
			if err != nil {
				fmt.Println("%s", err)
				return
			}

			interfaceName := args[0]

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)
			defer close(closeCh)

			reply, err := client.PluginInterfaceCommand(context.Background(), &protoGenerated.PluginInterfaceRequest{
				Interface: interfaceName,
			})
			if err != nil {
				fmt.Println("%s", err)
				if reply != nil {
					fmt.Println("%s", protoGenerated.Error{
						Code:    reply.Error.GetCode(),
						Message: reply.Error.GetMessage(),
					})
				}
				return
			} else {
				common.UI.InterfacePluginList(fmt.Sprintf("%s", reply))
			}
		},
	}
}
