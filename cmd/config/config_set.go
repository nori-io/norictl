package config_cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/nori-io/nori-grpc/pkg/api/proto"
	"github.com/nori-io/norictl/internal/errors"

	"github.com/spf13/cobra"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
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
		/* @todo versionPlugin := pluginIdSplit[1]
		_, err = version.NewVersion(versionPlugin)
		if err != nil {
			errors.ErrorFormatPluginVersion(err)
			return
		}*/

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)
		fmt.Println("args2 is", args[2])
		reply, err := client.ConfigSet(context.Background(), &proto.ConfigSetRequest{
			Id: &proto.ID{
				PluginId: pluginIdSplit[0],
				Version:  pluginIdSplit[1],
			},
			Key:   args[1],
			Value: args[2],
		})

		close(closeCh)

		if err != nil {
			fmt.Println(err)
			common.UI.ConfigSetFailure(pluginId, args[1], args[2], err)
			if reply != nil {
				fmt.Println(proto.Error{
					Code:    "",
					Message: reply.String(),
				})
			}
		} else {
			common.UI.ConfigSetSuccess(pluginId, args[1], args[2])
		}
	},
}
