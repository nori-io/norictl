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

// Package plugin_cmd implements commands for work with plugins
//by command prompt*/
package plugin_cmd

import (
	"github.com/nori-io/nori-common/v2/logger"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/client/utils"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var (
	listAll         func() bool
	listError       func() bool
	listInactive    func() bool
	listInstallable func() bool
	listInstalled   func() bool
	listRunning     func() bool
)

func lsCmd(log logger.Logger) *cobra.Command {

	return &cobra.Command{
		Use:     "norictl plugin ls [OPTIONS]",
		Aliases: []string{"list"},
		Short:   "Shows list of plugins on remote Nori node.",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsLs(log)
			conn, err := connection.CurrentConnection()
			if err != nil {
				log.Fatal("%s", err)
			}

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)

			reply, err := client.PluginListCommand(context.Background(), &protoNori.PluginListRequest{
				FlagAll:              listAll(),
				FlagError:            listError(),
				FlagInstalled:        listInstalled(),
				FlagRunning:          listRunning(),
				FlagInstallable:      listInstallable(),
				FlagInactive:         listInactive(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			})
			close(closeCh)
			if err != nil {
				if reply != nil {
					log.Fatal("%s", protoNori.ErrorReply{
						Status:               false,
						Error:                err.Error(),
						XXX_NoUnkeyedLiteral: struct{}{},
						XXX_unrecognized:     nil,
						XXX_sizecache:        0,
					})
				}
				log.Fatal("%s", err)
			}

			list := []*protoNori.PluginListWithStatus{{
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
			}}

			filter := func(list []*protoNori.PluginListWithStatus, f func(p protoNori.PluginListWithStatus) bool) []*protoNori.PluginListWithStatus {
				newList := make([]*protoNori.PluginListWithStatus, 0)
				plugins := make([][]string, len(list))
				for _, l := range list {
					if f(*l) {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})
					}
				}
				common.UI.PluginsAll(plugins)
				return newList
			}

			if listAll() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})
					}
					common.UI.PluginsAll(plugins)
					return p.FlagAll
				})
			}

			if listError() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})

					}
					common.UI.PluginsError(plugins)
					return p.FlagError
				})
			}

			if listInactive() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})
					}
					common.UI.PluginsInactive(plugins)
					return p.FlagInactive
				})
			}

			if listInstallable() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})

					}
					common.UI.PluginsInstallable(plugins)
					return p.FlagInstallable
				})
			}

			if listInstalled() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})

					}
					common.UI.PluginsInstalled(plugins)
					return p.FlagInstalled
				})
			}

			if listRunning() {
				list = filter(list, func(p protoNori.PluginListWithStatus) bool {
					newList := make([]*protoNori.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.MetaID.String(), l.Author.String()})
					}
					common.UI.PluginsRunning(plugins)
					return p.FlagRunning
				})
			}

		},
	}
}

func init() {
}

func setFlagsLs(log logger.Logger) {
	flags := utils.NewFlagBuilder(PluginCmd(log), lsCmd(log))
	flags.Bool(&listAll, "all", "--all", false, "Show all plugins")                                          // TODO
	flags.Bool(&listError, "error", "-e", false, "Show plugins with errors (not implement)")                 // TODO
	flags.Bool(&listInactive, "inactive", "--inactive", false, "Show plugins that are not running")          // TODO
	flags.Bool(&listInstallable, "installable", "--installable", false, "Show plugins that need to install") // TODO
	flags.Bool(&listInstalled, "installed", "-i", false, "Show only installed plugins")
	flags.Bool(&listRunning, "running", "-r", false, "Show only running plugins")
}
