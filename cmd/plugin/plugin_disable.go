package plugin_cmd

import (
	"fmt"
	"github.com/nori-io/nori-grpc/pkg/api/proto"
	"github.com/nori-io/norictl/internal/errors"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"strings"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
)

var disableCmd = &cobra.Command{
	Use:   "disable [PLUGIN_ID]",
	Short: "disable plugin",
	Long:  `Disable plugin. Disabled plugins are non usable plugins. Nori engine does not init, start, stop or resolve disabled plugins.`,
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
		defer close(closeCh)

		reply, err := client.PluginDisable(context.Background(), &proto.PluginRequest{
			Id: &proto.ID{
				PluginId: pluginIdSplit[0],
				Version:  pluginIdSplit[1],
			},
		})

		if (err != nil) || (reply.Error.GetCode() != "") {
			if err != nil {
				fmt.Println(err)
			}
			if reply.Error.GetCode() != "" {
				fmt.Println(proto.Error{
					Code:    reply.Error.GetCode(),
					Message: reply.Error.GetMessage(),
				})
			}
			common.UI.PluginDisableFailure(pluginId)
			return
		} else {
			common.UI.PluginDisableSuccess(pluginId)
		}
	},
}
