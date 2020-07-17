package config_cmd

import (
	"context"
	"fmt"
	"github.com/nori-io/norictl/internal/errors"
	"strings"

	"github.com/nori-io/nori-common/v2/version"
	"github.com/spf13/cobra"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"

	protoGenerated "github.com/nori-io/norictl/pkg/proto"
)

var setCmd = &cobra.Command{
	Use:   "set [PLUGIN_ID][KEY] [VALUE]",
	Short: "set plugin's config",
	Long:  `Set sets key's value for specify plugin's config.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := connection.CurrentConnection()
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) == 0 {
			errors.ErrorEmptyPluginId()
			return
		}

		if len(args) == 1 {
			fmt.Println("key and value required")
			return
		}
		if len(args) == 2 {
			fmt.Println("value required")
			return
		}

		pluginId := args[0]
		pluginIdSplit := strings.Split(pluginId, ":")
		if len(pluginIdSplit) != 2 {
			errors.ErrorFormatPluginId()
			return
		}
		versionPlugin := pluginIdSplit[1]
		_, err = version.NewVersion(versionPlugin)
		if err != nil {
			errors.ErrorFormatPluginVersion(err)
			return
		}

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)
		fmt.Println("args2 is", args[2])
		reply, err := client.ConfigSetCommand(context.Background(), &protoGenerated.ConfigSetRequest{
			Id: &protoGenerated.ID{
				PluginId: pluginIdSplit[0],
				Version:  pluginIdSplit[1],
			},
			Key:   args[1],
			Value: args[2],
		})

		close(closeCh)

		if err != nil {
			fmt.Println(err)
			common.UI.ConfigSetFailure(pluginId, args[1], args[2])
			if reply != nil {
				fmt.Println(protoGenerated.Error{
					Code:    "",
					Message: reply.String(),
				})
			}
		} else {
			common.UI.ConfigSetSuccess(pluginId, args[1], args[2])
		}
	},
}
