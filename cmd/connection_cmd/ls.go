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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/secure2work/norictl/client/utils"
)

var (
	host     bool
	insecure bool
	secureLs bool
	quiet    bool
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List connections",
	Long:  `List connections.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ls called")
	},
}

func init() {
	flags := utils.NewFlagBuilder(ConnectionCmd, lsCmd)
	flags.Bool(&host, "host", "", false, "Show connections only for given hostname")
	flags.Bool(&insecure, "insecure", "i", false, "Show only insecure connections")
	flags.Bool(&secureLs, "secure", "s", false, "Show only secure connections")
	flags.Bool(&quiet, "quiet", "q", false, "Only display connection names")
}
