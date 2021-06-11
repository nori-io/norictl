package errors

import "fmt"

func ErrorEmptyPluginId() {
	fmt.Println("PLUGIN_ID required!")
}

func ErrorFormatPluginId() {
	fmt.Println("Incorrect plugin's id format. Use ':' separator between plugin's name and plugin's version ")
}

func ErrorFormatPluginVersion(err error) {
	fmt.Println("Format of plugin's version is incorrect:", err)
}
