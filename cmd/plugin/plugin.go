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
	"github.com/spf13/cobra"

	"github.com/nori-io/norictl/internal/ui"
)

var PluginCmd = &cobra.Command{
	Use:   "norictl plugin",
	Short: "norictl plugin COMMAND",
}

var UI *ui.UI

func init() {
	UI=ui.NewUI()
}
