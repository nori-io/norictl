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
	Green("GET SUCCESSFUL: Plugin %q \n", pluginId)
}

func (u *UI) PluginGetFailure(pluginId string, err error) {
	Red("GET FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) PluginDisableSuccess(pluginId string) {
	Green("DISABLE SUCCESSFUL: Plugin %q\n", pluginId)
}

func (u *UI) PluginDisableFailure(pluginId string, err error) {
	Red("DISABLE FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) PluginEnableSuccess(pluginId string) {
	Green("ENABLE SUCCESSFUL: Plugin %q\n", pluginId)
}

func (u *UI) PluginEnableFailure(pluginId string, err error) {
	Red("ENABLE FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) PluginInstallSuccess(pluginId string) {
	Green("INSTALL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginInstallFailure(pluginId string, err error) {
	Red("INSTALL FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) PluginInstallAllSuccess() {
	Green("INSTALL SUCCESSFUL")
}

func (u *UI) PluginInstallAllFailure(err error) {
	Red("INSTALL FAILURE error %s \n", err)
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

func (u *UI) PluginPullFailure(pluginId string, err error) {
	Red("PULL FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) PluginRmSuccess(pluginId string) {
	Green("REMOVE SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginRmFailure(pluginId string, err error) {
	Red("REMOVE FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) PluginStartAllSuccess() {
	Green("START SUCCESSFUL")
}

func (u *UI) PluginStartAllFailure(err error) {
	Red("START FAILURE error %s \n", err)
}

func (u *UI) PluginStartSuccess(pluginId string) {
	Green("START SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginStartFailure(pluginId string, err error) {
	Red("START FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) PluginStopSuccess(pluginId string) {
	Green("STOP SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginStopFailure(pluginId string, err error) {
	Red("STOP FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) PluginStopAllSuccess() {
	Green("STOP SUCCESSFUL")
}

func (u *UI) PluginStopAllFailure(err error) {
	Red("STOP FAILURE error %s \n", err)
}

func (u *UI) PluginUninstallSuccess(pluginId string) {
	Green("UNINSTALL SUCCESSFUL: Plugin %s\n", pluginId)
}

func (u *UI) PluginUninstallFailure(pluginId string, err error) {
	Red("UNINSTALL FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) PluginUninstallAllSuccess() {
	Green("UNINSTALL SUCCESSFUL")
}

func (u *UI) PluginUninstallAllFailure(err error) {
	Red("UNINSTALL FAILURE error %s \n", err)
}

func (u *UI) PluginUploadSuccess(path string) {
	Green("UPLOAD SUCCESSFUL: Path:  %s\n", path)
}

func (u *UI) PluginUploadFailure(path string, err error) {
	Red("UPLOAD FAILURE: Path: %s, error %s\n", path, err)
}

func (u *UI) ConfigGetSuccess(keyValueMap map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value"})
	for key, value := range keyValueMap {
		table.Append([]string{key, value})
	}
	table.Render() // Send output
}

func (u *UI) ConfigGetFailure(pluginId string, err error) {
	Red("GET FAILURE: Plugin %q, error %s \n", pluginId, err)
}

func (u *UI) ConfigSetSuccess(pluginId, key, value string) {
	Green("SET SUCCESSFUL: Plugin %q, key: %s, value, %s\n", pluginId, key, value)
}

func (u *UI) ConfigSetFailure(pluginId, key, value string, err error) {
	Red("SET FAILURE: Plugin %q, key: %s, value, %s, error: %s\n", pluginId, key, value, err)
}

func (u *UI) ConfigUploadSuccess(keyValueMap map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value"})
	for key, value := range keyValueMap {
		table.Append([]string{key, value})
	}
	table.Render() // Send output
}

func (u *UI) ConfigUploadFailure(path string, err error) {
	Red("Upload FAILURE: Config's Path: %s, error %s\n", path, err)
}
