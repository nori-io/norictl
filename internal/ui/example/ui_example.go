package main

import (
	"fmt"

	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
	"github.com/nori-io/norictl/internal/ui"
)

func main() {
	uiExample := ui.NewUI()
	uiExample.GetSuccess("nori/session:0.0.0")
	uiExample.GetFailure("nori/session:0.0.1")
	uiExample.InstallSuccess("nori/session:0.0.2")
	uiExample.InstallFailure("nori/session:0.0.3")
	plugins := make([][]string, 2)
	plugins = append(plugins, []string{"nori/session:0.0.4", "author1"})
	plugins = append(plugins, []string{"nori/session:0.0.5", "author2"})
	uiExample.PluginsAll(plugins)
	uiExample.PluginMetaExist(fmt.Sprintf("%v", protoNori.PluginMetaReply{
		ArrayPluginListWithoutStatus: []*protoNori.PluginListWithoutStatus{&protoNori.PluginListWithoutStatus{
			MetaID: &protoNori.ID{
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
				MetaID: &protoNori.ID{
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
	uiExample.PullSuccess("nori/session:0.0.8")
	uiExample.PullFailure("nori/session:0.0.9")
	uiExample.RmSuccess("nori/session:0.0.10")
	uiExample.RmFailure("nori/session:0.0.11")
	uiExample.StartSuccess("nori/session:0.0.12")
	uiExample.StartFailure("nori/session:0.0.13")
	uiExample.StopSuccess("nori/session:0.0.14")
	uiExample.StopFailure("nori/session:0.0.15")

}
