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

	"github.com/nori-io/nori-grpc/pkg/api/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
)

var (
	installAll bool
)

var installAllCmd = &cobra.Command{
	Use:   "install-all [OPTIONS]",
	Short: "Install all downloaded plugins.",
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := connection.CurrentConnection()
		if err != nil {
			fmt.Println(err)
		}

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)
		defer close(closeCh)

		flagVerbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			fmt.Println(err)
			return
		}

		reply, err := client.PluginInstall(context.Background(), &proto.PluginInstallRequest{
			FlagAll:     installAll,
			FlagVerbose: flagVerbose,
		})

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
			common.UI.PluginInstallAllFailure(err)
			return
		}
		common.UI.PluginInstallAllSuccess()

	},
}

func init() {
	installAllCmd.Flags().BoolVarP(&installAll, "all", "a", false, "Install all installable plugins")
}
