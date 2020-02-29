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

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/ui"
)

func PluginCmd(log logger.Logger) *cobra.Command {

	PluginCmd := &cobra.Command{
		Use:   "norictl plugin",
		Short: "norictl plugin COMMAND",
	}

	PluginCmd.AddCommand(getCmd(log))
	PluginCmd.AddCommand(installCmd(log))
	PluginCmd.AddCommand(interfaceCmd(log))
	PluginCmd.AddCommand(lsCmd(log))
	PluginCmd.AddCommand(metaCmd(log))
	PluginCmd.AddCommand(pullCmd(log))
	PluginCmd.AddCommand(rmCmd(log))
	PluginCmd.AddCommand(startCmd(log))
	PluginCmd.AddCommand(stopCmd(log))
	PluginCmd.AddCommand(uninstallCmd(log))
	PluginCmd.AddCommand(uploadCmd(log))
	return PluginCmd
}

func init() {
	common.UI = ui.NewUI()
}
