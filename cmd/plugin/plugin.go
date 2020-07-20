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
	"github.com/spf13/cobra"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/ui"
)

func PluginCmd() *cobra.Command {

	PluginCmd := &cobra.Command{
		Use:   "plugin",
		Short: "plugin COMMAND",
	}

	PluginCmd.AddCommand(getCmd)
	PluginCmd.AddCommand(disableCmd)
	PluginCmd.AddCommand(enableCmd)
	PluginCmd.AddCommand(installCmd)
	PluginCmd.AddCommand(interfaceCmd)
	PluginCmd.AddCommand(lsCmd)
	PluginCmd.AddCommand(metaCmd)
	PluginCmd.AddCommand(pullCmd)
	PluginCmd.AddCommand(rmCmd)
	PluginCmd.AddCommand(startCmd)
	PluginCmd.AddCommand(stopCmd)
	PluginCmd.AddCommand(uninstallCmd)
	PluginCmd.AddCommand(uploadCmd)

	return PluginCmd
}

func init() {
	common.UI = ui.NewUI()
}
