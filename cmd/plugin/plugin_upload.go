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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/nori-io/nori-common/v2/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/utils"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var (
	uploadFile func() string
)

func uploadCmd(log logger.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "norictl plugin upload [OPTIONS]",
		Short: "Upload the plugin from local machine.",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsUpload(log)
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
				log.Fatal("%s", err)
			}

			defer f.Close()

			_, err = ioutil.ReadAll(f)
			if err != nil {
				log.Fatal("%s", err)
			}
			path = filepath.Base(path)

			reply, err := client.PluginUploadCommand(context.Background(), &protoNori.PluginUploadRequest{
				Filepath:             path,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			})
			if err != nil {
				common.UI.UploadFailure(path)
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
				common.UI.UploadSuccess(path)
			}
		},
	}
}

func init() {

}

func setFlagsUpload(log logger.Logger) {
	flags := utils.NewFlagBuilder(PluginCmd(log), uploadCmd(log))
	flags.String(&uploadFile, "file", "--file", "", "Specify path to plugin") // TODO
}
