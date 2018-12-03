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

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new connection to remote Nori node.",
	Long:  `Create new connection to remote Nori node.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	createCmd.Flags().StringP(CONN_CREATE_CERT, CONN_CREATE_CERT_SHORT, "", "Path to certificate file")
	//createCmd.MarkFlagRequired("cert")
	createCmd.Flags().StringP(CONN_CREATE_NAME, CONN_CREATE_NAME_SHORT, "default", "Connection name")
	createCmd.Flags().BoolP(CONN_CREATE_SECURE, CONN_CREATE_SECURE_SHORT, false, "Use or not certificate for connection")

	viper.BindPFlag(CONN_CREATE_CERT, createCmd.Flags().Lookup(CONN_CREATE_CERT))
	viper.BindPFlag(CONN_CREATE_NAME, createCmd.Flags().Lookup(CONN_CREATE_NAME))
	viper.BindPFlag(CONN_CREATE_SECURE, createCmd.Flags().Lookup(CONN_CREATE_SECURE))

	ConnectionCmd.AddCommand(createCmd)
}
