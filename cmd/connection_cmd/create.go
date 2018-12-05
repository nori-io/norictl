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
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/secure2work/norictl/client/connection"
	"github.com/secure2work/norictl/client/consts"
	"github.com/secure2work/norictl/client/utils"
)

var (
	cert   string
	name   string
	secure bool
	force  bool
)

var createCmd = &cobra.Command{
	Use:   "create [OPTIONS] [SERVER]",
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

		var hostname string
		if len(args) == 0 {
			hostname = consts.DefaultHostname
		} else {
			hostname = args[0]
		}

		conn := &connection.Connection{
			Name:     name,
			Hostname: hostname,
			Secure:   secure,
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
	flags.Bool(&secure, "secure", "s", false, "Use or not certificate for connection")
	flags.Bool(&force, "force", "f", false, "Force rewrite connection")
}
