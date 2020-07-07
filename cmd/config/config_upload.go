package config_cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	protoGenerated "github.com/nori-io/norictl/internal/generated/protobuf"
)

func uploadCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "upload [PATH]",
		Short: "upload plugin's config",
		Long:  `Upload shows config file from specify path.`,
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := connection.CurrentConnection()
			if err != nil {
				fmt.Println("%s", err)
				return
			}

			if len(args) == 0 {
				fmt.Println("Path to config's file required")
				return
			}

			path := args[0]

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)

			reply, err := client.ConfigUploadCommand(context.Background(), &protoGenerated.ConfigUploadRequest{
				Path: path,
			})

			close(closeCh)

			if err != nil {
				fmt.Println("%s", err)
				common.UI.ConfigUploadFailure(path)
				if reply != nil {
					fmt.Println("%s", protoGenerated.ErrorReply{
						Status: false,
						Error:  err.Error(),
					})
				}
			} else {
				common.UI.ConfigUploadSuccess(reply.KeyValueMapField.KeyValueMap)
			}
		},
	}
}
