// Copyright © 2018 Nori info@nori.io
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

	"github.com/nori-io/norictl/client/connection"
	"github.com/nori-io/norictl/client/consts"
	"github.com/nori-io/norictl/client/utils"
)

var format func() string

var showCmd = &cobra.Command{
	Use:   "show [NAME]",
	Short: "Shows detailed information about specific connection.",
	Long:  `Shows detailed information about specific connection.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(consts.ConnectionsDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		var name string
		if len(args) == 0 {
			name = consts.DefaultName
		} else {
			name = args[0]
		}

		list, err := connection.List(consts.ConnectionsDir)
		if err != nil {
			log.Fatal(err)
		}

		conn, err := list.FilterByName(name)
		if err != nil {
			log.Fatal(err)
		}

		r, err := conn.Render(format())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(r)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	flags := utils.NewFlagBuilder(ConnectionCmd, showCmd)
	flags.String(&format, "format", "f", "table", "Data representation template: json or table")
}
