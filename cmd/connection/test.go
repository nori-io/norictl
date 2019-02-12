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
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/nori/proto"
	"github.com/nori-io/norictl/client"
	"github.com/nori-io/norictl/client/connection"
	"github.com/nori-io/norictl/client/utils"
)

var verbose func() bool

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Make connection to remote Nori node to verify connection configuration.",
	Long:  `Make connection to remote Nori node to verify connection configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := connection.CurrentConnection()
		if err != nil {
			log.Fatal(err)
		}

		cli, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)

		msg := fmt.Sprintf("%d", rand.Uint64())

		ping := &commands.PingRequest{Message: msg}

		reply, err := cli.SendPingCommand(context.Background(), ping)
		close(closeCh)
		if err != nil {
			log.Fatal(err)
		}
		if reply.Message == msg {
			fmt.Println("OK")
		} else {
			fmt.Println("Error: Pong message does not match")
		}
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())

	utils.NewFlagBuilder(ConnectionCmd, testCmd)
}
