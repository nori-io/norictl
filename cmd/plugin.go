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
	pluginCmd.PersistentFlags().String("grpc-address", "localhost:29876", "gRPC host and port")
	pluginCmd.PersistentFlags().String("ca", "client.ca", "client ca for gRPC")
	pluginCmd.PersistentFlags().String("ServerHostOverride", "", "ServerHostOverride")
	pluginCmd.PersistentFlags().Bool("dependencies", true, "install dependencies")

	viper.BindPFlag("plugin-path", pluginCmd.PersistentFlags().Lookup("plugin-path"))
	viper.BindPFlag("grpc-address", pluginCmd.PersistentFlags().Lookup("grpc-address"))
	viper.BindPFlag("ca", pluginCmd.PersistentFlags().Lookup("ca"))
	viper.BindPFlag("ServerHostOverride", pluginCmd.PersistentFlags().Lookup("ServerHostOverride"))
	viper.BindPFlag("dependencies", pluginCmd.PersistentFlags().Lookup("dependencies"))
}
