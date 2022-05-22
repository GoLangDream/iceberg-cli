package project

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

var configFiles = []string{DatabaseConfigFilePath()}

func initConfig() {
	config.AddDriver(yaml.Driver)
	for _, configFile := range configFiles {
		configFile := Path(configFile)
		config.LoadFiles(configFile)
	}
}
