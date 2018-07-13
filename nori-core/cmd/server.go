// Copyright © 2018 Secure2Work info@secure2work.com
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/secure2work/nori-cli/nori-core/core"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Run: func(cmd *cobra.Command, args []string) {
		enable := viper.GetBool("grpc-enable")
		addr := viper.GetString("grpc-address")

		server := core.NewServer([]string{"../plugins"}, addr, enable)
		err := server.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Bool("grpc-enable", true, "use gRPC")
	serverCmd.Flags().String("grpc-address", "localhost:12345", "gRPC host and port")

	viper.BindPFlag("grpc-enable", serverCmd.Flags().Lookup("grpc-enable"))
	viper.BindPFlag("grpc-address", serverCmd.Flags().Lookup("grpc-address"))
}
