package ui

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type UI struct {
	color o
}

func NewUI() *UI {
	//todo load config
	return &UI{}
}

func (u *UI) GetSuccess(pluginId string) {
	fmt.Printf("Plugin %q successfully get\n", pluginId)
}

func (u *UI) GetFailure(pluginId string) {
	fmt.Printf("Plugin %q get failure\n", pluginId)
}

func (u *UI) InstallSuccess (pluginId string){
	fmt.Printf("Plugin %s installed, %3d :\n", pluginId)
}

func (u *UI) InstallFailure(pluginId string){
	fmt.Printf("Plugin %s install failure, %3d :\n", pluginId)
}

func (u *UI) AllPlugins(plugins [][]string){
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Plugin.MetaId", "Author"})

	for _, v := range plugins {
		table.Append(v)
	}
	table.Render() // Send output
}