package config_cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/nori-io/nori-common/v2/version"
	"github.com/spf13/cobra"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"

	protoGenerated "github.com/nori-io/norictl/internal/generated/protobuf"
)

func setCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "set [PLUGIN_ID][KEY] [VALUE]",
		Short: "set plugin's config",
		Long:  `Set sets key's value for specify plugin's config.`,
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := connection.CurrentConnection()
			if err != nil {
				fmt.Println("%s", err)
			}

			if len(args) == 0 {
				fmt.Println("PLUGIN_ID required!")
				return
			}

			pluginId := args[0]
			pluginIdSplit := strings.Split(pluginId, ":")
			versionPlugin := pluginIdSplit[1]
			_, err = version.NewVersion(versionPlugin)
			if err != nil {
				fmt.Println("Format of plugin's version is incorrect:", err)
			}

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)

			reply, err := client.ConfigSetCommand(context.Background(), &protoGenerated.ConfigSetRequest{
				Id: &protoGenerated.ID{
					Id:                   pluginIdSplit[0],
					Version:              pluginIdSplit[1],
				},
				Key:                  args[1],
				Value:                args[2],
			})

			close(closeCh)

			if err != nil {
				fmt.Println("%s", err)
				common.UI.ConfigSetFailure(pluginId, args[1], args[2])
				if reply != nil {
					fmt.Println("%s", protoGenerated.ErrorReply{
						Status:               false,
						Error:                err.Error(),
					})
				}
			} else {
				common.UI.ConfigSetSuccess(pluginId, args[1], args[2])
			}
		},
	}
}

func init() {

}
