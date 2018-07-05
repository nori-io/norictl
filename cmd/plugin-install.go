package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "installing plugins",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("plugin-path")
		if len(path) == 0 && len(args) > 0 {
			path = args[0]
		}

		toolchain, err := SetupToolChain()
		if err != nil {
			log.Fatal(err)
		}

		if !viper.GetBool("dep-install") {
			toolchain.InstallDependencies = false
		}

		err = toolchain.Do(path)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Plugin %q successfully installed\n", path)
	},
}

func init() {
	pluginCmd.AddCommand(installCmd)

	installCmd.Flags().Bool("dep-install", true, "install dependencies")
	viper.BindPFlag("dep-install", installCmd.Flags().Lookup("dep-install"))
}
