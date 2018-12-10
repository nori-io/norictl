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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/secure2work/norictl/client/connection"
	"github.com/secure2work/norictl/client/consts"
	"github.com/secure2work/norictl/client/utils"
)

var (
	host     func() string
	insecure func() bool
	secureLs func() bool
	quiet    func() bool
	formatLs func() string
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List connections",
	Long:  `List connections.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(consts.ConnectionsDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		list, err := connection.List(consts.ConnectionsDir)
		if err != nil {
			log.Fatal(err)
		}

		if insecure() && secureLs() {
			log.Error("")
			insecure = func() bool { return false }
			secureLs = func() bool { return false }
		}

		if insecure() {
			list = list.FilterBySecure(false)
		}

		if secureLs() {
			list = list.FilterBySecure(true)
		}

		if len(host()) > 0 {
			list = list.FilterByHost(host())
		}

		if quiet() {
			for _, l := range list {
				fmt.Println(l.Name)
			}
		} else {
			str, err := list.Render(formatLs())
			if err != nil {
				log.Fatal(err)
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
