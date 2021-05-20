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

	"github.com/nori-io/nori-grpc/pkg/api/proto"
	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
)

var (
	uninstallAll        bool
	uninstallDependents bool
)

var uninstallAllCmd = &cobra.Command{

	Use:   "uninstall [OPTIONS]",
	Short: "Uninstall plugins.",
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := connection.CurrentConnection()
		if err != nil {
			fmt.Println(err)
			return
		}
		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)

		reply, err := client.PluginUninstall(context.Background(), &proto.PluginUninstallRequest{
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
			common.UI.PluginUninstallAllFailure()
			return
		}
		common.UI.PluginUninstallAllSuccess()

	},
}

func init() {
	uninstallAllCmd.Flags().BoolVarP(&uninstallAll, "all", "a", false, "Uninstall all installed plugins")
	uninstallAllCmd.Flags().BoolVarP(&uninstallDependent, "dependent", "d", false, "Uninstall plugin and depend plugins")
}
