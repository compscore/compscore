package config

import (
	"bytes"
	"os"
	"text/template"

	"github.com/compscore/compscore/pkg/structs"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	ConfigFile        string = "config.yml"
	RunningConfigFile string = "running-config.yml"
	CheckFileName     string = "check.so"

	RunningConfig *structs.RunningConfig_s
	Config        *structs.Config_s
)

func init() {
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
	exists, err := FileExists(RunningConfigFile)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to check if config file exists")
	}

	if !exists {
		config, runningConfig, err := GenerateIntialConfig()
		if err != nil {
			logrus.WithError(err).Fatal("Failed to generate running config")
		}

		RunningConfig = runningConfig
		Config = config
	} else {
		UpdateRunningConfig()
	}
}

func UpdateRunningConfig() {
	viper.SetConfigFile(RunningConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to read running config")
	}

	err = viper.Unmarshal(&RunningConfig)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal running config")
	}
}

func GenerateIntialConfig() (*structs.Config_s, *structs.RunningConfig_s, error) {
	var (
		runningConfig *structs.RunningConfig_s = &structs.RunningConfig_s{}
		config        *structs.Config_s        = &structs.Config_s{}

		name    string
		web     structs.Web_s
		teams_s structs.Teams_s
		engine  structs.Engine_s
		checks  []structs.Check_s

		teams []structs.Team_s
	)

	viper.SetConfigFile(ConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		return config, runningConfig, err
	}

	name = viper.GetString("name")
	runningConfig.Name = name
	config.Name = name

	err = viper.UnmarshalKey("web", &web)
	if err != nil {
		return config, runningConfig, err
	}
	runningConfig.Web = web
	config.Web = web

	err = viper.UnmarshalKey("teams", &teams_s)
	if err != nil {
		return config, runningConfig, err
	}
	runningConfig.Teams = teams
	config.Teams = teams_s

	err = viper.UnmarshalKey("engine", &engine)
	if err != nil {
		return config, runningConfig, err
	}
	runningConfig.Engine = engine
	config.Engine = engine

	err = viper.UnmarshalKey("checks", &checks)
	if err != nil {
		return config, runningConfig, err
	}
	config.Checks = checks

	name_template := template.Must(template.New("name").Parse(teams_s.NameFormat))

	for i := 0; i < teams_s.Amount; i++ {
		team_name := bytes.NewBuffer([]byte{})

		err := name_template.Execute(
			team_name,
			struct{ Team int }{Team: i + 1},
		)
		if err != nil {
			return config, runningConfig, err
		}

		team := structs.Team_s{
			Name:     team_name.String(),
			Number:   i + 1,
			Password: teams_s.Password,
		}

		checks := []structs.Check_s{}

		for _, check := range config.Checks {
			target := bytes.NewBuffer([]byte{})

			target_template := template.Must(template.New("target").Parse(check.Target))

			err := target_template.Execute(
				target,
				struct{ Team int }{Team: i + 1},
			)
			if err != nil {
				return config, runningConfig, err
			}

			checks = append(checks, structs.Check_s{
				Name: check.Name,
				Git: structs.Git_s{
					Remote: check.Git.Remote,
					Branch: check.Git.Branch,
				},
				Credentials: structs.Credentials_s{
					Username: check.Credentials.Username,
					Password: check.Credentials.Password,
				},
				Port:           check.Port,
				Command:        check.Command,
				Target:         target.String(),
				ExpectedOutput: check.ExpectedOutput,
				Weight:         check.Weight,
			})

			team.Checks = checks
			teams = append(teams, team)
		}
	}
	runningConfig.Teams = teams

	file, err := os.Create(RunningConfigFile)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to create running config file")
		return config, runningConfig, err
	}

	out, err := yaml.Marshal(runningConfig)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to marshal running config")
		return config, runningConfig, err
	}

	_, err = file.Write(out)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to write running config to file")
		return config, runningConfig, err
	}

	err = file.Close()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to close running config file")
		return config, runningConfig, err
	}

	viper.SetConfigFile(RunningConfigFile)
	err = viper.ReadInConfig()
	if err != nil {
		return config, runningConfig, err
	}

	return config, runningConfig, nil
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
