package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var certsCmd = &cobra.Command{
	Use:   "certs",
	Short: "certs",
}

func init() {
	rootCmd.AddCommand(certsCmd)

	certsCmd.PersistentFlags().String("pem", "server.pem", "path to pem file")
	certsCmd.PersistentFlags().String("key", "server.key", "path to key file")
	certsCmd.PersistentFlags().String("passkey", "", "secret passkey")

	viper.BindPFlag("pem", certsCmd.PersistentFlags().Lookup("pem"))
	viper.BindPFlag("key", certsCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("passkey", certsCmd.PersistentFlags().Lookup("passkey"))
}
