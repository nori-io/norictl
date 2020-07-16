package config_cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	protoGenerated "github.com/nori-io/norictl/pkg/proto"
)

var uploadCmd = &cobra.Command{
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
		close(closeCh)

		f, err := os.Open(path)
		if err != nil {
			fmt.Println("%s", err)
			return
		}

		defer f.Close()

		_, err = ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("%s", err)
			return
		}
		path = filepath.Base(path)

		reply, err := client.ConfigUploadCommand(context.Background(), &protoGenerated.ConfigUploadRequest{
			Config: []byte{},
		})

		if err != nil {
			fmt.Println("%s", err)
			common.UI.ConfigUploadFailure(path)
			if reply != nil {
				fmt.Println("%s", protoGenerated.Error{
					Code:    "",
					Message: reply.String(),
				})
			}
		} else {
			common.UI.ConfigUploadSuccess(reply.Map)
		}
	},
}
