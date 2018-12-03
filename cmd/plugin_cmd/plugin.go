package plugin_cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var PluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "plugin",
}

func init() {
	PluginCmd.PersistentFlags().String("plugin-path", "", "path to plugin")
	PluginCmd.PersistentFlags().String("plugin-id", "", "plugin id")
	PluginCmd.PersistentFlags().String("grpc-address", "localhost:29876", "gRPC host and port")
	PluginCmd.PersistentFlags().String("ca", "client.ca", "client ca for gRPC")
	PluginCmd.PersistentFlags().String("ServerHostOverride", "", "ServerHostOverride")
	PluginCmd.PersistentFlags().Bool("dependencies", true, "install dependencies")

	viper.BindPFlag("plugin-path", PluginCmd.PersistentFlags().Lookup("plugin-path"))
	viper.BindPFlag("grpc-address", PluginCmd.PersistentFlags().Lookup("grpc-address"))
	viper.BindPFlag("ca", PluginCmd.PersistentFlags().Lookup("ca"))
	viper.BindPFlag("ServerHostOverride", PluginCmd.PersistentFlags().Lookup("ServerHostOverride"))
	viper.BindPFlag("dependencies", PluginCmd.PersistentFlags().Lookup("dependencies"))
	viper.BindPFlag("plugin-id", PluginCmd.PersistentFlags().Lookup("plugin-id"))
}
