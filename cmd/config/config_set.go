package config_cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/nori-io/nori-common/v2/version"
	"github.com/spf13/cobra"

	"github.com/nori-io/nori-common/v2/logger"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/generated/protobuf/config"

	commonProtoGenerated "github.com/nori-io/norictl/internal/generated/protobuf/common"
)

func setCmd(log logger.FieldLogger) *cobra.Command {

	return &cobra.Command{
		Use:   "set [PLUGIN_ID][KEY] [VALUE]",
		Short: "set plugin's config",
		Long:  `Set sets key's value for specify plugin's config.`,
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := connection.CurrentConnection()
			if err != nil {
				log.Fatal("%s", err)
			}

			if len(args) == 0 {
				log.Fatal("PLUGIN_ID required!")
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

			reply, err := client.ConfigSetCommand(context.Background(), &config.ConfigSetRequest{
				Id: &commonProtoGenerated.ID{
					Id:                   pluginIdSplit[0],
					Version:              pluginIdSplit[1],
				},
				Key:                  args[1],
				Value:                args[2],
			})

			close(closeCh)

			if err != nil {
				log.Fatal("%s", err)
				common.UI.ConfigSetFailure(pluginId, args[1], args[2])
				if reply != nil {
					log.Fatal("%s", commonProtoGenerated.ErrorReply{
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
