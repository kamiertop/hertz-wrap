package config

import (
	"hertz/pkg/consts"
)

var Conf config

type config struct {
	Log    Log    `toml:"log"`
	System System `toml:"system"`
}

type Log struct {
	FileName string `toml:"filename"`
	Level    string `toml:"level"`
}

type System struct {
	Env string `toml:"env"`
}

func InitConfig() error {
	setDefaultConfig()
	return nil
}

// setDefaultConfig set default value if not sets
func setDefaultConfig() {
	if Conf.System.Env == "" {
		Conf.System.Env = consts.DevelopmentMode
	}

}
