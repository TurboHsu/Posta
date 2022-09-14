package config

import (
	"fmt"
	"io"
	"os"

	"github.com/TurboHsu/Posta/util"
	"github.com/pelletier/go-toml/v2"
)

var C Config

func Read(f string) {
	_, err := os.Stat(f)
	if err != nil {
		//Config file do not exist
		err := Write(f)
		if err != nil {
			util.LogFatal(fmt.Sprintf("Error when creating config file: %s", err.Error()))
		}
		util.LogFatal("Config file do not exist, created one.")
	}
	file, err := os.Open(f)
	util.HandleErr(err)
	defer file.Close()
	data, err := io.ReadAll(file)
	util.HandleErr(err)
	toml.Unmarshal(data, &C)
}
