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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/internal/client"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var uploadCmd = &cobra.Command{
	Use:   "norictl plugin upload [OPTIONS]",
	Short: "Upload the plugin from local machine.",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("file")

		if len(path) == 0 && len(args) > 0 {
			path = args[0]
		}

		client, closeCh := client.NewClient(
			viper.GetString("grpc-address"),
			viper.GetString("ca"),
			viper.GetString("ServerHostOverride"),
		)
		defer close(closeCh)

		f, err := os.Open(path)
		if err != nil {
			logrus.Fatal(err)
		}

		defer f.Close()

		so, err := ioutil.ReadAll(f)
		if err != nil {
			logrus.Fatal(err)
		}
		path = filepath.Base(path)

		reply, err := client.PluginUploadCommand(context.Background(), &protoNori.PluginUploadRequest{
			Filepath:             path,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})

		if err != nil {
			logrus.Fatal(err)
			if reply != nil {
				logrus.Fatal(protoNori.ErrorReply{
					Status:               false,
					Error:                err.Error(),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				})
			}
		} else {
			fmt.Printf("Plugin %q successfully uploaded\n", path)
		}
	},
}

func init() {
	PluginCmd.AddCommand(uploadCmd)
}
