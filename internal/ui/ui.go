package ui

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type UI struct {
	//color
}

func NewUI() *UI {
	//todo load config
	return &UI{}
}

func (u *UI) GetSuccess(pluginId string) {
	fmt.Printf("GET SUCCESSFUL: Plugin %q\n", pluginId)
}

func (u *UI) GetFailure(pluginId string) {
	fmt.Printf("GET FAILURE: Plugin %q\n", pluginId)
}

func (u *UI) InstallSuccess(pluginId string) {
	fmt.Printf("INSTALL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) InstallFailure(pluginId string) {
	fmt.Printf("INSTALL FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) PluginsAll(plugins [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Plugin.MetaId", "Author"})

	for _, v := range plugins {
		table.Append(v)
	}
	table.Render() // Send output
}

func (u *UI) PluginsInstalled(plugins [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Plugin.MetaId", "Author"})

	for _, v := range plugins {
		table.Append(v)
	}
	table.Render() // Send output
}

func (u *UI) PluginMetaExist(pluginInformation string) {
	fmt.Printf("EXIST: Plugin  %s \n", pluginInformation)
}

func (u *UI) PluginMetaNotExist(pluginInformation string) {
	fmt.Printf("Plugin %q get failure\n", pluginInformation)
}

func (u *UI) PullSuccess(pluginId string) {
	fmt.Printf("PULL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PullFailure(pluginId string) {
	fmt.Printf("PULL FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) RmSuccess(pluginId string) {
	fmt.Printf("REMOVE SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) RmFailure(pluginId string) {
	fmt.Printf("REMOVE FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) StartSuccess(pluginId string) {
	fmt.Printf("START SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) StartFailure(pluginId string) {
	fmt.Printf("START FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) StopSuccess(pluginId string) {
	fmt.Printf("STOP SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) StopFailure(pluginId string) {
	fmt.Printf("STOP FAILURE: Plugin %s\n", pluginId)
}
