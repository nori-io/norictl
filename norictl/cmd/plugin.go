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
	pluginCmd.PersistentFlags().String("grpc-address", "localhost:12345", "gRPC host and port")

	viper.BindPFlag("plugin-path", pluginCmd.PersistentFlags().Lookup("plugin-path"))
	viper.BindPFlag("grpc-address", pluginCmd.PersistentFlags().Lookup("grpc-address"))
}
