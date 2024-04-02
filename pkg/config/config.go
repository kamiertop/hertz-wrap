package config

import (
	"flag"

	"github.com/BurntSushi/toml"

	"hertz/pkg/consts"
)

var configFile = flag.String("config", "config.toml", "config file location")

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
	Env  string `toml:"env"`
	Addr string `toml:"addr"`
}

func InitConfig() error {
	flag.Parse()
	_, err := toml.DecodeFile(*configFile, &Conf)
	if err != nil {
		return err
	}

	setDefaultConfig()
	return nil
}

// setDefaultConfig set default value if not sets
func setDefaultConfig() {
	if Conf.System.Env == "" {
		Conf.System.Env = consts.DevelopmentMode
	}
}
