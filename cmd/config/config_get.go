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

		reply, err := client.ConfigGet(context.Background(), &protoGenerated.ConfigGetRequest{
			Id: &protoGenerated.ID{
				PluginId: pluginIdSplit[0],
				Version:  pluginIdSplit[1],
			},
		})

		close(closeCh)

		if err != nil {
			fmt.Println(err)
			common.UI.ConfigGetFailure(pluginId)
			if reply != nil {
				fmt.Println(protoGenerated.Error{
					Code:    "",
					Message: err.Error(),
				})
			}
		} else {
			fmt.Println(reply.Map)
			common.UI.ConfigGetSuccess(reply.Map)
		}
	},
}
