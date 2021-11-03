package Config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	configOnce sync.Once
	config     *Config
	err        error
)

type Config struct {
	HttpAddr string
	Mysql    map[string]string
}

func newConfig() *Config {
	return &Config{
		HttpAddr: "127.0.0.1",
		Mysql:    map[string]string{},
	}
}

func GetConfig() (*Config, error) {
	configOnce.Do(func() {
		config, err = prepareConfig()
	})
	return config, err
}

func Initialize(cfgfile string) {
	viper.SetConfigFile(cfgfile)
}

func prepareConfig() (*Config, error) {
	var (
		err error
	)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := newConfig()

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
