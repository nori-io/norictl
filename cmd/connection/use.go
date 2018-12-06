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

var useCmd = &cobra.Command{
	Use:   "use [NAME]",
	Short: "Define a connection to use.",
	Long:  `Define a connection to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		path := filepath.Join(home, consts.ConfigDir)
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fpath := filepath.Join(path, consts.UseConnFilename)

		path = filepath.Join(path, consts.ConnectionsDir)

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

		err = conn.Use(fpath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	ConnectionCmd.AddCommand(useCmd)
}
