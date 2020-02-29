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
	"net"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/client/consts"
	"github.com/nori-io/norictl/internal/client/utils"
)

var (
	cert  func() string
	name  func() string
	force func() bool
)

var createCmd = &cobra.Command{
	Use:   "create [OPTIONS] [[HOST]:PORT]",
	Short: "Create new connection to remote Nori node.",
	Long:  `Create new connection to remote Nori node.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(consts.ConnectionsDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
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
				log.Fatal(err)
			}
			port, err = strconv.ParseUint(portStr, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
		}

		conn := &connection.Connection{
			Name:     name(),
			Host:     host,
			Port:     port,
			CertPath: cert(),
		}
		err = conn.Save(consts.ConnectionsDir, force())
		if err != nil {
			log.Fatal(err)
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
