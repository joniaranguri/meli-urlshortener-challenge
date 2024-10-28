package configs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/olebedev/config"
)

var Conf *config.Config

func init() {

	confDir := os.Getenv("CONF_DIR")

	if confDir == "" {
		confDir, _ = filepath.Abs(filepath.FromSlash("./conf"))
	}

	if _, err := os.Stat(confDir); os.IsNotExist(err) {
		panic(err.Error())
	}

	var err error
	Conf, err = config.ParseYamlFile(path.Join(confDir, "default.yml"))
	if err != nil {
		panic(err.Error())
	}

	confFiles := []string{
		fmt.Sprintf("%s.yml", os.Getenv("SCOPE")),
	}

	for _, file := range confFiles {

		absPath := path.Join(confDir, file)

		if _, err := os.Stat(absPath); os.IsNotExist(err) {

			continue
		}

		newConf, err := config.ParseYamlFile(absPath)
		if err != nil {
			panic(err.Error())
		}

		Conf, err = Conf.Extend(newConf)
		if err != nil {
			panic(err.Error())
		}

	}
}
