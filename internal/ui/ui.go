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

func (u *UI) PluginGetSuccess(pluginId string) {
	Green("GET SUCCESSFUL: Plugin %q\n", pluginId)
}

func (u *UI) PluginGetFailure(pluginId string) {
	Red("GET FAILURE: Plugin %q\n", pluginId)
}

func (u *UI) PluginDisableSuccess(pluginId string) {
	Green("DISABLE SUCCESSFUL: Plugin %q\n", pluginId)
}

func (u *UI) PluginDisableFailure(pluginId string) {
	Red("DISABLE FAILURE: Plugin %q\n", pluginId)
}

func (u *UI) PluginEnableSuccess(pluginId string) {
	Green("ENABLE SUCCESSFUL: Plugin %q\n", pluginId)
}

func (u *UI) PluginEnableFailure(pluginId string) {
	Red("ENABLE FAILURE: Plugin %q\n", pluginId)
}

func (u *UI) PluginInstallSuccess(pluginId string) {
	Green("INSTALL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginInstallFailure(pluginId string) {
	Red("INSTALL FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) PluginInstallAllSuccess() {
	Green("INSTALL SUCCESSFUL")
}

func (u *UI) PluginInstallAllFailure() {
	Red("INSTALL FAILURE")
}

func (u *UI) PluginsList(plugins [][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Plugin.ID", "Author", "Interface", "Licenses", "Dependencies"})
	for _, v := range plugins {
		table.Append(v)
	}
	table.Render() // Send output
}

func (u *UI) PluginMeta(pluginInformation string) {
	fmt.Printf("Meta: Plugin  %s \n", pluginInformation)
}

func (u *UI) PluginMetaNotExist(pluginInformation string) {
	fmt.Printf("Plugin %q get failure\n", pluginInformation)
}

func (u *UI) PluginPullSuccess(pluginId string) {
	Green("PULL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginPullFailure(pluginId string) {
	Red("PULL FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) PluginRmSuccess(pluginId string) {
	Green("REMOVE SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginRmFailure(pluginId string) {
	Red("REMOVE FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) PluginStartSuccess(pluginId string) {
	Green("START SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginStartFailure(pluginId string) {
	Red("START FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) PluginStopSuccess(pluginId string) {
	Green("STOP SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginStopFailure(pluginId string) {
	Red("STOP FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) PluginUninstallSuccess(pluginId string) {
	Green("UNINSTALL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginUninstallFailure(pluginId string) {
	Red("UNINSTALL FAILURE: Plugin %s\n", pluginId)
}

func (u *UI) PluginUploadSuccess(path string) {
	Green("UPLOAD SUCCESSFUL: Path:  %s\n", path)
}

func (u *UI) PluginUploadFailure(path string) {
	Red("UPLOAD FAILURE: Path: %s\n", path)
}

func (u *UI) ConfigGetSuccess(keyValueMap map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value"})
	for key, value := range keyValueMap {
		table.Append([]string{key, value})
	}
	table.Render() // Send output
}

func (u *UI) ConfigGetFailure(pluginId string) {
	Red("GET FAILURE: Plugin %q\n", pluginId)
}

func (u *UI) ConfigSetSuccess(pluginId, key, value string) {
	Green("SET SUCCESSFUL: Plugin %q, key: %s, value, %s\n", pluginId, key, value)
}

func (u *UI) ConfigSetFailure(pluginId, key, value string) {
	Red("SET FAILURE: Plugin %q, key: %s, value, %s,\n", pluginId, key, value)
}

func (u *UI) ConfigUploadSuccess(keyValueMap map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value"})
	for key, value := range keyValueMap {
		table.Append([]string{key, value})
	}
	table.Render() // Send output
}

func (u *UI) ConfigUploadFailure(path string) {
	Red("Upload FAILURE: Config's path: %q\n", path)
}
