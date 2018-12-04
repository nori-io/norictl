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
	"github.com/spf13/viper"

	. "github.com/secure2work/norictl/client/consts"
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
	lsCmd.Flags().BoolP(CONN_LS_HOST, CONN_LS_HOST_SHORT, false, "Show connections only for given hostname")
	lsCmd.Flags().BoolP(CONN_LS_INSECURE, CONN_LS_INSECURE_SHORT, false, "Show only insecure connections")
	lsCmd.Flags().BoolP(CONN_LS_SECURE, CONN_LS_SECURE_SHORT, false, "Show only secure connections")
	lsCmd.Flags().BoolP(CONN_LS_QUIET, CONN_LS_QUIET_SHORT, false, "Only display connection names")

	viper.BindPFlag(CONN_LS_HOST_VIPER, lsCmd.Flags().Lookup(CONN_LS_HOST))
	viper.BindPFlag(CONN_LS_INSECURE_VIPER, lsCmd.Flags().Lookup(CONN_LS_INSECURE))
	viper.BindPFlag(CONN_LS_SECURE_VIPER, lsCmd.Flags().Lookup(CONN_LS_SECURE))
	viper.BindPFlag(CONN_LS_QUIET_VIPER, lsCmd.Flags().Lookup(CONN_LS_QUIET))

	ConnectionCmd.AddCommand(lsCmd)
}
