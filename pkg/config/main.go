package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/compscore/compscore/pkg/structs"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	ConfigFile string = "config.yml"

	Name       string
	Engine     structs.Engine_s
	Web        structs.Web_s
	Teams      structs.Teams_s
	Scoring    structs.Scoring_s
	Checks     []structs.Check_s
	AdminUsers []structs.AdminUser_s
)

func Init() {
	viper.SetConfigFile(ConfigFile)

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

	err := viper.UnmarshalKey("engine", &Engine)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal engine config")
	}

	Web, err = web()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal web config")
	}

	err = viper.UnmarshalKey("teams", &Teams)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal teams config")
	}

	err = viper.UnmarshalKey("scoring", &Scoring)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal scoring config")
	}

	err = viper.UnmarshalKey("checks", &Checks)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal checks config")
	}

	err = viper.UnmarshalKey("users", &AdminUsers)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal users config")
	}
}

func web() (structs.Web_s, error) {
	hostname := os.Getenv("DOMAIN")
	portStr := os.Getenv("PORT")
	jwtKey := os.Getenv("JWT_SECRET")
	timeoutStr := os.Getenv("TIMEOUT")
	releaseStr := os.Getenv("RELEASE")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return structs.Web_s{}, err
	}

	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		return structs.Web_s{}, err
	}

	release := strings.ToLower(releaseStr) == "true" || strings.ToLower(releaseStr) == "yes" || releaseStr == "1"

	return structs.Web_s{
		Hostname: hostname,
		Port:     port,
		JWTKey:   jwtKey,
		Timeout:  timeout,
		Release:  release,
	}, nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
