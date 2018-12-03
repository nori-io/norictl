package plugin_cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/secure2work/nori/core/grpc"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build plugin",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("plugin-path")

		if len(path) == 0 && len(args) > 0 {
			logrus.Fatal("plugin-path is required")
		}

		toolchain, err := grpc.SetupToolChain()
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
	PluginCmd.AddCommand(buildCmd)
}
