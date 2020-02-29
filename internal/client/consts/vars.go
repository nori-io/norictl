package consts

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
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
		log.Fatal(err)
	}

	ConfigDir = filepath.Join(home, configDir)
	ConnectionsDir = filepath.Join(ConfigDir, connectionsDir)
	ConfigPath = filepath.Join(ConfigDir, configName)
	UseFilePath = filepath.Join(ConfigDir, UseConnFilename)
}
