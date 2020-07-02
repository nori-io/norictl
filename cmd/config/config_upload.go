package config_cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/nori-io/nori-common/v2/logger"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	commonProtoGenerated "github.com/nori-io/norictl/internal/generated/protobuf/common"
	"github.com/nori-io/norictl/internal/generated/protobuf/config"
)

func uploadCmd(log logger.Logger) *cobra.Command {

	return &cobra.Command{
		Use:   "norictl config upload [PATH]",
		Short: "upload plugin's config",
		Long:  `Upload shows config file from specify path.`,
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := connection.CurrentConnection()
			if err != nil {
				log.Fatal("%s", err)
			}

			if len(args) == 0 {
				log.Fatal("PLUGIN_ID required!")
			}

			path := args[0]

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)

			reply, err := client.ConfigUploadCommand(context.Background(), &config.ConfigUploadRequest{
				Path:                 path,
			})

			close(closeCh)

			if err != nil {
				log.Fatal("%s", err)
				common.UI.ConfigUploadFailure(path)
				if reply != nil {
					log.Fatal("%s", commonProtoGenerated.ErrorReply{
						Status:               false,
						Error:                err.Error(),
					})
				}
			} else {
				common.UI.ConfigUploadSuccess(reply.KeyValueMapField.KeyValueMap)
			}
		},
	}
}

func init() {

}
