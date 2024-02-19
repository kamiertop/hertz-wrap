package config

import (
	"sync/atomic"

	"hertz/pkg/consts"
)

// conf concurrent safe
var conf = new(atomic.Value)

func Cfg() Config {
	return conf.Load().(Config)
}

type Config struct {
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
	conf.Store(Config{})
	setDefaultConfig()
	return nil
}

// setDefaultConfig set default value if not sets
func setDefaultConfig() {
	if Cfg().System.Env == "" {
		t := Cfg()
		t.System.Env = consts.DevelopmentMode
		conf.Store(t)
	}

}
