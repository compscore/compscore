package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Name    string
	Web     Web_s
	Teams   Teams_s
	Scoring Scoring_s
	Engine  Engine_s
	Checks  []Check_s
)

func init() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to read config file")
	}

	UpdateConfiguration()

	viper.OnConfigChange(
		func(e fsnotify.Event) {
			if e.Op == fsnotify.Write {
				logrus.Info("Config file changed:", e.Name)
				UpdateConfiguration()
			}
		},
	)

	viper.WatchConfig()
}

func UpdateConfiguration() {
	Name = viper.GetString("name")

	err := viper.UnmarshalKey("web", &Web)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal web")
	}

	err = viper.UnmarshalKey("teams", &Teams)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal teams")
	}

	err = viper.UnmarshalKey("scoring", &Scoring)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal scoring")
	}

	err = viper.UnmarshalKey("engine", &Engine)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal engine")
	}

	err = viper.UnmarshalKey("checks", &Checks)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal checks")
	}

	logrus.Info("Configuration updated")
}
