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

	"github.com/nori-io/nori-common/v2/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

func interfaceCmd(log logger.Logger) *cobra.Command {

	return &cobra.Command{
		Use:   "norictl plugin interface [InterfaceName]",
		Short: "Shows list of plugins that implement specify interface.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("InterfaceName required!!!")
			}

			interfaceName := args[0]

			client, closeCh := client.NewClient(
				viper.GetString("grpc-address"),
				viper.GetString("ca"),
				viper.GetString("ServerHostOverride"),
			)
			defer close(closeCh)

			reply, err := client.PluginInterfaceCommand(context.Background(), &protoNori.PluginInterfaceRequest{
				InterfaceName:        interfaceName,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			})
			if err != nil {
				log.Fatal("%s", err)
				if reply != nil {
					log.Fatal("%s", protoNori.ErrorReply{
						Status:               false,
						Error:                err.Error(),
						XXX_NoUnkeyedLiteral: struct{}{},
						XXX_unrecognized:     nil,
						XXX_sizecache:        0,
					})
				}
			} else {
				common.UI.InterfacePluginList(fmt.Sprintf("%s", reply))
			}
		},
	}
}

func init() {
}
