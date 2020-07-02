package main

import (
	"fmt"

	"github.com/nori-io/norictl/internal/generated/protobuf/common"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
	"github.com/nori-io/norictl/internal/ui"
)

func main() {
	uiExample := ui.NewUI()
	uiExample.PluginGetSuccess("nori/session:0.0.0")
	uiExample.PluginGetFailure("nori/session:0.0.1")
	uiExample.PluginInstallSuccess("nori/session:0.0.2")
	uiExample.PluginInstallFailure("nori/session:0.0.3")
	plugins := make([][]string, 2)
	plugins = append(plugins, []string{"nori/session:0.0.4", "author1"})
	plugins = append(plugins, []string{"nori/session:0.0.5", "author2"})
	uiExample.PluginsAll(plugins)
	uiExample.PluginsInstalled(plugins)

	uiExample.PluginMetaExist(fmt.Sprintf("%v", protoNori.PluginMetaReply{
		ArrayPluginListWithoutStatus: []*protoNori.PluginListWithoutStatus{&protoNori.PluginListWithoutStatus{
			MetaID: &common.ID{
				Id:                   "nori/session",
				Version:              "0.0.6",
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Author: &protoNori.Author{
				Name:                 "author",
				URI:                  "",
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			DependenciesArray:    nil,
			Description:          nil,
			Core:                 nil,
			Interface:            nil,
			License:              nil,
			Links:                nil,
			Repository:           nil,
			Tags:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},
			&protoNori.PluginListWithoutStatus{
				MetaID: &common.ID{
					Id:                   "nori/session",
					Version:              "0.0.7",
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				Author: &protoNori.Author{
					Name:                 "author",
					URI:                  "",
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				DependenciesArray:    nil,
				Description:          nil,
				Core:                 nil,
				Interface:            nil,
				License:              nil,
				Links:                nil,
				Repository:           nil,
				Tags:                 nil,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}},
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0},
	))
	uiExample.PluginPullSuccess("nori/session:0.0.8")
	uiExample.PluginPullFailure("nori/session:0.0.9")
	uiExample.PluginRmSuccess("nori/session:0.0.10")
	uiExample.PluginRmFailure("nori/session:0.0.11")
	uiExample.PluginStartSuccess("nori/session:0.0.12")
	uiExample.PluginStartFailure("nori/session:0.0.13")
	uiExample.PluginStopSuccess("nori/session:0.0.14")
	uiExample.PluginStopFailure("nori/session:0.0.15")
	uiExample.PluginUninstallSuccess("nori/session:0.0.14")
	uiExample.PluginUninstallFailure("nori/session:0.0.15")
	uiExample.PluginUploadSuccess("/plugins")
	uiExample.PluginUploadFailure("/plugins")
	mapKeyValue := map[string]string{"port": "8080", "address": "http://localhost"}
	uiExample.ConfigGetSuccess(mapKeyValue)
	uiExample.ConfigGetFailure("nori/session:0.0.16")
	uiExample.ConfigSetSuccess("nori/session:0.0.17", "port", "8080")
	uiExample.ConfigSetFailure("nori/session:0.0.18", "port", "0000")
	uiExample.ConfigUploadSuccess(mapKeyValue)
	uiExample.ConfigUploadFailure("/path")

}
