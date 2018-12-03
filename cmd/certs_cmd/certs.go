package certs_cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CertsCmd = &cobra.Command{
	Use:   "certs",
	Short: "certs",
}

func init() {
	CertsCmd.PersistentFlags().String("pem", "server.pem", "path to pem file")
	CertsCmd.PersistentFlags().String("key", "server.key", "path to key file")
	CertsCmd.PersistentFlags().String("passkey", "", "secret passkey")

	viper.BindPFlag("pem", CertsCmd.PersistentFlags().Lookup("pem"))
	viper.BindPFlag("key", CertsCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("passkey", CertsCmd.PersistentFlags().Lookup("passkey"))
}
