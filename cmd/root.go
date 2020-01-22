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
	"bytes"
	"fmt"
	"os"

	 "github.com/nori-io/logger"
	logger2 "github.com/nori-io/nori-common/logger"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nori-io/norictl/cmd/certs"
	"github.com/nori-io/norictl/cmd/connection"
	"github.com/nori-io/norictl/cmd/plugin"
	"github.com/nori-io/norictl/internal/client/consts"
	"github.com/nori-io/norictl/internal/client/utils"
)

var cfgFile string
var logLevel func() string
var LoggerNoriCtl logger2.Logger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "norictl",
	Short: "A simple command line client for nori",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	buf := bytes.Buffer{}
	LoggerNoriCtl=logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf))
	//LoggerNoriCtl:=logger.New()
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	flags := utils.NewFlagBuilder(nil, rootCmd)
	flags.StringP(&logLevel, "verbose", "", "error", "set verbose level (debug info warn error fatal panic)")

	rootCmd.AddCommand(plugin_cmd.PluginCmd)
	rootCmd.AddCommand(certs_cmd.CertsCmd)
	rootCmd.AddCommand(connection_cmd.ConnectionCmd)
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
		lvl, err := log.ParseLevel(logLevel())
		if err != nil {
			log.Error(err)
		} else {
			log.SetLevel(lvl)
		}
		log.Info("Using config file: ", viper.ConfigFileUsed())
	}
}
