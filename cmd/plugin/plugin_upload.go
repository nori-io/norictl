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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/utils"
	protoGenerated "github.com/nori-io/norictl/pkg/proto"
)

var (
	uploadFile func() string
)

func uploadCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "upload [OPTIONS]",
		Short: "Upload the plugin from local machine.",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsUpload()
			conn, err := connection.CurrentConnection()
			if err != nil {
				fmt.Println("%s", err)
				return
			}

			//path := viper.GetString("file")
			path:=args[1]
			if len(path) == 0 && len(args) > 1 {
				path = args[0]
			}

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)
			defer close(closeCh)

			fmt.Println("Current folder", path)
			f, err := os.Open(path)
			if err != nil {
				fmt.Println("%s", err)
				return
			}

			defer f.Close()

			_, err = ioutil.ReadAll(f)
			if err != nil {
				fmt.Println("%s", err)
				return
			}
			path = filepath.Base(path)

			reply, err := client.PluginUploadCommand(context.Background())
			if err != nil {
				common.UI.PluginUploadFailure(path)
				fmt.Println("%s", err)
				if reply != nil {
					fmt.Println("%s", protoGenerated.Error{
						Code:    "",
						Message: "",
					})
					return
				}
				return
			} else {
				common.UI.PluginUploadSuccess(path)
			}
		},
	}
}

func setFlagsUpload() {
	flags := utils.NewFlagBuilder(PluginCmd(), uploadCmd())
	flags.String(&uploadFile, "file", "f", "", "Specify path to plugin") // TODO
}
