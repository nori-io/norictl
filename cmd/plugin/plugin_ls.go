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
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
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

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)

		reply, err := client.PluginListCommand(context.Background(), &commands.PluginListRequest{})
		close(closeCh)
		if err != nil {
			if reply != nil {
				log.Fatal(protoNori.ErrorReply{
					Status:               false,
					Error:                err.Error(),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				})
			}
			log.Fatal(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "ID", "Name", "Author"})

		list :=  []*protoNori.PluginListWithStatus{{
				MetaID:               nil,
				Author:               nil,
				DependenciesArray:    nil,
				Description:          nil,
				Core:                 nil,
				Interface:            nil,
				License:              nil,
				Links:                nil,
				Repository:           nil,
				Tags:                 nil,
				FlagAll:              true,
				FlagError:            false,
				FlagInstalled:        false,
				FlagRunning:          false,
				FlagInstallable:      false,
				FlagInactive:         false,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},}
		filter := func(list []*protoNori.PluginListWithStatus, f func(p protoNori.PluginListWithStatus) bool) []*protoNori.PluginListWithStatus {
			newList := make([]*protoNori.PluginListWithStatus, 0)

			for _, l := range list {
				if f(*l) {
					newList = append(newList, l)
				}
			}
			return newList
		}

		if listInstalled() {
			list = filter(list, func(p protoNori.PluginListWithStatus) bool {
				return p.FlagInstalled
			})
		}

		if listRunning() {
			list = filter(list, func(p protoNori.PluginListWithStatus) bool {
				return p.FlagRunning
			})
		}

		for i, v := range list {
			table.Append([]string{
				strconv.Itoa(i + 1), v.MetaID.String(), v.Author.String(),
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
