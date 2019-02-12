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
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nori-io/nori/core/grpc"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build plugin",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("plugin-path")

		if len(path) == 0 && len(args) > 0 {
			logrus.Fatal("plugin-path is required")
		}

		toolchain, err := grpc.SetupToolChain()
		if err != nil {
			logrus.Fatal(err)
		}

		toolchain.InstallDependencies = viper.GetBool("dependencies")

		err = toolchain.Do(path)
		if err != nil {
			logrus.Fatal(err)
		} else {
			fmt.Printf("Plugin %q successfully built\n", path)
		}
	},
}

func init() {
	PluginCmd.AddCommand(buildCmd)
}
