package config

import (
	"flag"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var configFile = flag.String("config", "config.toml", "config file location")

var (
	Conf config
	V    *viper.Viper
)

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
	V = viper.New()
	V.SetConfigFile(*configFile)
	if err := V.ReadInConfig(); err != nil {
		return fmt.Errorf("read config file failed, err: %w", err)
	}

	if err := V.Unmarshal(&Conf, func(config *mapstructure.DecoderConfig) {
		config.TagName = "toml"
	}); err != nil {
		return fmt.Errorf("init config failed, unmarshal error: %w", err)
	}
	fmt.Printf("init config success, env: %#v\n", Conf)
	return nil
}

func RewriteConfig() error {
	return V.WriteConfig()
}
