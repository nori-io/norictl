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
)

var rmCmd = &cobra.Command{
	Use:   "rm [NAME]",
	Short: "Remove connection",
	Long:  `Remove connection configuration.`,
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

		var name string
		if len(args) == 0 {
			fmt.Println("Error: Name required.\n")
			cmd.Usage()
			os.Exit(1)
		} else {
			name = args[0]
		}

		list, err := connection.List(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		conn, err := list.FilterByName(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = conn.Remove()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	ConnectionCmd.AddCommand(rmCmd)
}
