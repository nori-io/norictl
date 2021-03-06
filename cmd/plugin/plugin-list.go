// Copyright © 2018 Nori info@nori.io
//
// This program is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package plugin_cmd

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/nori/proto"
	"github.com/nori-io/norictl/client"
	"github.com/nori-io/norictl/client/connection"
	"github.com/nori-io/norictl/client/utils"
)

var (
	listError     func() bool
	listInstalled func() bool
	listRunning   func() bool
)

var listCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Aliases: []string{"list"},
	Short:   "Shows list of plugins on remote Nori node.",
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

		reply, err := cli.PluginListCommand(context.Background(), &commands.PluginListRequest{})
		close(closeCh)
		if err != nil {
			if reply != nil {
				log.Fatal(reply.Error)
			}
			log.Fatal(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "ID", "Name", "Author"})

		list := reply.Data

		filter := func(list []*commands.PluginList, f func(p *commands.PluginList) bool) []*commands.PluginList {
			newList := make([]*commands.PluginList, 0)

			for _, l := range list {
				if f(l) {
					newList = append(newList, l)
				}
			}
			return newList
		}

		if listInstalled() {
			list = filter(list, func(p *commands.PluginList) bool {
				return p.Installed
			})
		}

		if listRunning() {
			list = filter(list, func(p *commands.PluginList) bool {
				return p.Running
			})
		}

		for i, v := range list {
			table.Append([]string{
				strconv.Itoa(i + 1), v.Id, v.Name, v.Author,
			})
		}
		table.Render()
	},
}

func init() {
	flags := utils.NewFlagBuilder(PluginCmd, listCmd)
	flags.Bool(&listError, "error", "e", false, "Show plugins with errors (not implement)") // TODO
	flags.Bool(&listInstalled, "installed", "i", false, "Show only installed plugins")
	flags.Bool(&listRunning, "running", "r", false, "Show only running plugins")
}
