package config_cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/nori-io/norictl/internal/errors"

	"github.com/spf13/cobra"

	"github.com/nori-io/nori-grpc/pkg/api/proto"
	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
)

var getCmd = &cobra.Command{
	Use:   "get [PLUGIN_ID]",
	Short: "get plugin's config",
	Long:  `Get shows specify plugin's config`,
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

		pluginId := args[0]
		pluginIdSplit := strings.Split(pluginId, ":")
		if len(pluginIdSplit) != 2 {
			errors.ErrorFormatPluginId()
			return
		}

		/*@todo		versionPlugin := pluginIdSplit[1]


		_, err = version.NewVersion(versionPlugin)
		if err != nil {
			errors.ErrorFormatPluginVersion(err)
			return
		}
		*/
		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)

		/*reply, err := proto.NoriClient.ConfigGet(context.Background(), &proto.ConfigGetRequest{
			Id: &proto.ID{
				PluginId: pluginIdSplit[0],
				Version:  pluginIdSplit[1],
			},
		}, nil)*/

		reply, err := client.ConfigGet(context.Background(), &proto.ConfigGetRequest{Id: &proto.ID{
			PluginId: pluginIdSplit[0],
			Version:  pluginIdSplit[1],
		}})

		close(closeCh)

		if err != nil {
			fmt.Println(err)
			common.UI.ConfigGetFailure(pluginId, err)
			return
		}

		if reply.Map == nil {
			fmt.Println("Config not found")
			common.UI.ConfigGetFailure(pluginId, err)

			return
		}
		fmt.Println(reply.Map)
		common.UI.ConfigGetSuccess(reply.Map)

	},
}
