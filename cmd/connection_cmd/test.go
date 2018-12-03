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

	. "github.com/secure2work/norictl/cmd/consts"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Make connection to remote Nori node to verify connection configuration.",
	Long:  `Make connection to remote Nori node to verify connection configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
	},
}

func init() {
	testCmd.Flags().BoolP(CONN_TEST_VERBOSE, CONN_TEST_VERBOSE_SHORT, false, "Show connection detailed information")

	viper.BindPFlag(CONN_TEST_VERBOSE, testCmd.Flags().Lookup(CONN_TEST_VERBOSE))

	ConnectionCmd.AddCommand(testCmd)
}
