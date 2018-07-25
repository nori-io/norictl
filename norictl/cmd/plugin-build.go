package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/secure2work/nori-cli/nori-core/core"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build plugin",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("plugin-path")

		if len(path) == 0 && len(args) > 0 {
			logrus.Fatal("plugin-path is required")
		}

		toolchain, err := core.SetupToolChain()
		if err != nil {
			logrus.Fatal(err)
		}

		toolchain.InstallDependencies = viper.GetBool("dependencies")

		err = toolchain.Do(path)
		if err != nil {
			logrus.Fatal(err)
		} else {
			fmt.Printf("Plugin %q successfully built\n", path)
		}
	},
}

func init() {
	pluginCmd.AddCommand(buildCmd)

	buildCmd.Flags().Bool("dependencies", true, "install dependencies")
	viper.BindPFlag("dependencies", buildCmd.Flags().Lookup("dependencies"))
}
