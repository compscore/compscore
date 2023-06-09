package config

import (
	"github.com/compscore/compscore/structs"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Competition structs.Competition
)

func init() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Error reading config file")
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
	err := viper.Unmarshal(&Competition)
	if err != nil {
		logrus.WithError(err).Error("Error unmarshalling config file")
	}
}
