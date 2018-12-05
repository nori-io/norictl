// Copyright © 2018 Secure2Work info@secure2work.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package connection_cmd

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/secure2work/norictl/client/connection"
	"github.com/secure2work/norictl/client/consts"
	"github.com/secure2work/norictl/client/utils"
)

var (
	cert  string
	name  string
	force bool
)

var createCmd = &cobra.Command{
	Use:   "create [OPTIONS] [[HOST]:PORT]",
	Short: "Create new connection to remote Nori node.",
	Long:  `Create new connection to remote Nori node.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		path := filepath.Join(home, consts.ConfigDir, consts.ConnectionsDir)
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var host, portStr string
		var port uint64

		if len(args) == 0 {
			host = consts.DefaultHost
			port = consts.DefaultPort
		} else {
			//A literal IPv6 address in hostport must be enclosed in square brackets, as in "[::1]:80", "[::1%lo0]:80".
			host, portStr, err = net.SplitHostPort(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			port, err = strconv.ParseUint(portStr, 10, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		conn := &connection.Connection{
			Name:     name,
			Host:     host,
			Port:     port,
			CertPath: cert,
		}
		err = conn.Save(path, force)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	flags := utils.NewFlagBuilder(ConnectionCmd, createCmd)
	flags.String(&cert, "cert", "c", "", "Path to certificate file")
	flags.String(&name, "name", "n", "default", "Connection name")
	flags.Bool(&force, "force", "f", false, "Force rewrite connection")
}
