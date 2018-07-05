package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "plugin",
}

func init() {
	rootCmd.AddCommand(pluginCmd)

	pluginCmd.PersistentFlags().String("plugin-path", "", "path to plugin")
	viper.BindPFlag("plugin-path", pluginCmd.PersistentFlags().Lookup("plugin-path"))
}
