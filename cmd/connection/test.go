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
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/secure2work/nori/proto"
	"github.com/secure2work/norictl/client"
	"github.com/secure2work/norictl/client/connection"
	"github.com/secure2work/norictl/client/consts"
	"github.com/secure2work/norictl/client/utils"
)

var verbose bool

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Make connection to remote Nori node to verify connection configuration.",
	Long:  `Make connection to remote Nori node to verify connection configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		path := filepath.Join(home, consts.ConfigDir)

		useFileName := filepath.Join(path, consts.UseConnFilename)

		bs, err := ioutil.ReadFile(useFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		name := string(bytes.TrimSpace(bs))

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

		cli, closeCh := client.NewClient(
			net.JoinHostPort(conn.Host, strconv.Itoa(int(conn.Port))),
			conn.CertPath,
			"",
		)

		msg := fmt.Sprintf("%d", rand.Uint64())

		ping := &commands.PingRequest{Message: msg}

		reply, err := cli.SendPingCommand(context.Background(), ping)
		close(closeCh)
		if err != nil {
			logrus.Fatal(err)
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

	flags := utils.NewFlagBuilder(ConnectionCmd, testCmd)
	// TODO implement verbose
	flags.Bool(&verbose, "verbose", "v", false, "Show connection detailed information (not implemented yet)")
}
