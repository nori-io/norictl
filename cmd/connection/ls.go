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
	host     string
	insecure bool
	secureLs bool
	quiet    bool
	formatLs string
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List connections",
	Long:  `List connections.`,
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

		list, err := connection.List(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if insecure {
			list = list.FilterBySecure(false)
		}

		if secureLs {
			list = list.FilterBySecure(true)
		}

		if len(host) > 0 {
			list = list.FilterByHost(host)
		}

		if quiet {
			for _, l := range list {
				fmt.Println(l.Name)
			}
		} else {
			str, err := list.Render(formatLs)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(str)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	flags := utils.NewFlagBuilder(ConnectionCmd, lsCmd)
	flags.String(&host, "host", "x", "", "Show connections only for given hostname")
	flags.Bool(&insecure, "insecure", "i", false, "Show only insecure connections")
	flags.Bool(&secureLs, "secure", "s", false, "Show only secure connections")
	flags.Bool(&quiet, "quiet", "q", false, "Only display connection names")
	flags.String(&formatLs, "format", "f", "table", "Data representation template: json or table")
}
