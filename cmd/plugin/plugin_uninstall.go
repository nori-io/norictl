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

package plugin_cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/utils"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
	"github.com/nori-io/norictl/internal/ui"
)

var (
	uninstallAll       func() bool
	uninstallDependent func() bool
)

var uninstallCmd = &cobra.Command{

	Use:   "norictl plugin uninstall [PLUGIN_ID] [OPTIONS]",
	Short: "Uninstall plugin or plugins.",
	Run: func(cmd *cobra.Command, args []string) {
		id := viper.GetString("id")
		if len(id) == 0 && len(args) > 0 {
			id = args[0]
		}

		cli, closeCh := client.NewClient(
			viper.GetString("grpc-address"),
			viper.GetString("ca"),
			viper.GetString("ServerHostOverride"),
		)

		reply, err := cli.PluginUninstallCommand(context.Background(), &protoNori.PluginUninstallRequest{
			Id: &protoNori.ID{
				Id:                   id,
				Version:              "",
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			FlagAll:              false,
			FlagDependent:        false,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})
		defer close(closeCh)
		uiIninstall := ui.NewUI()
		if err != nil {
			if reply != nil {
				uiIninstall.UninstallFailure(id)
				logrus.Fatal(protoNori.ErrorReply{
					Status:               false,
					Error:                err.Error(),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				})
			}
			logrus.Fatal(err)
		}

		uiIninstall.UninstallSuccess(id)
	},
}

func init() {
	PluginCmd.AddCommand(uninstallCmd)
	flags := utils.NewFlagBuilder(PluginCmd, uninstallCmd)
	flags.Bool(&uninstallAll, "all", "--all", false, "Uninstall all installed plugins")                       // TODO
	flags.Bool(&uninstallDependent, "dependent", "--dependent", false, "Uninstall plugin and depend plugins") // TODO
}
