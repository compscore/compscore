package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Hostname string
	Port     string
)

func init() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	UpdateConfigs()

	viper.OnConfigChange(
		func(e fsnotify.Event) {
			if e.Op == fsnotify.Write {
				UpdateConfigs()
			}
		},
	)

	viper.WatchConfig()

}

func UpdateConfigs() {
	Hostname = viper.GetString("hostname")
	Port = viper.GetString("port")
}
