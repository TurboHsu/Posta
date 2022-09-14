package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Write(f string) error {
	var c Config
	content, err := toml.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(f, []byte(content), 0644)
	return err
}
