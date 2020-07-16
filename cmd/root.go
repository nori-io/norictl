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

package cmd

import (
	"fmt"
	"os"

	config_cmd "github.com/nori-io/norictl/cmd/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nori-io/norictl/cmd/connection"
	"github.com/nori-io/norictl/cmd/plugin"
	"github.com/nori-io/norictl/internal/client/consts"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "norictl",
	Short: "A simple command line client for nori",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().String( "config", "", "config file")
	RootCmd.PersistentFlags().BoolP( "verbose", "v", false, "verbose output")

	RootCmd.AddCommand(plugin_cmd.PluginCmd())
	RootCmd.AddCommand(config_cmd.ConfigCmd())
	RootCmd.AddCommand(connection_cmd.ConnectionCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(consts.ConfigPath)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//lvl, err := log.ParseLevel(logLevel())
		if err != nil {
			fmt.Println(err)
		} else {
			//log.SetLevel(lvl)
		}
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}
