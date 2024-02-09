package config

var Conf Config

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

	return nil
}

// defaultConf set default value if not sets
func defaultConf() {

}
