package consts

import (
	"fmt"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	ConfigDir      string
	ConnectionsDir string
	ConfigPath     string
	UseFilePath    string
)

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
	}

	ConfigDir = filepath.Join(home, configDir)
	ConnectionsDir = filepath.Join(ConfigDir, connectionsDir)
	ConfigPath = filepath.Join(ConfigDir, configName)
	UseFilePath = filepath.Join(ConfigDir, UseConnFilename)
}
