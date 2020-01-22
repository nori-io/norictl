package ui

import (
	"fmt"
	"os"

	. "github.com/fatih/color"
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
	Green("GET SUCCESSFUL: Plugin %q\n", pluginId)
}

func (u *UI) GetFailure(pluginId string) {
	Red("GET FAILURE: Plugin %q\n", pluginId)
}

func (u *UI) InstallSuccess(pluginId string) {
	Green("INSTALL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) InstallFailure(pluginId string) {
	Red("INSTALL FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) PluginsAll(plugins [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Plugin.MetaId", "Author"})

	for _, v := range plugins {
		table.Append(v)
	}
	table.Render() // Send output
}

func (u *UI) PluginsError(plugins [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Plugin.MetaId", "Author"})
	for _, v := range plugins {
		table.Append(v)
	}
	table.Render() // Send output
}

func (u *UI) PluginsInactive(plugins [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Plugin.MetaId", "Author"})
	for _, v := range plugins {
		table.Append(v)
	}
	table.Render() // Send output
}

func (u *UI) PluginsInstallable(plugins [][]string) {
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

func (u *UI) PluginsRunning(plugins [][]string) {
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
	Green("PULL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PullFailure(pluginId string) {
	Red("PULL FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) RmSuccess(pluginId string) {
	Green("REMOVE SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) RmFailure(pluginId string) {
	Red("REMOVE FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) StartSuccess(pluginId string) {
	Green("START SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) StartFailure(pluginId string) {
	Red("START FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) StopSuccess(pluginId string) {
	Green("STOP SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) StopFailure(pluginId string) {
	Red("STOP FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) UninstallSuccess(pluginId string) {
	Green("UNINSTALL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) UninstallFailure(pluginId string) {
	Red("UNINSTALL FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) UploadSuccess(path string) {
	Green("UPLOAD SUCCESSFUL: Path:  %s\n", path)
}

func (u *UI) UploadFailure(path string) {
	Red("UPLOAD FAILURE: Path: %s\n", path)
}
