package config

import (
	"bytes"
	"os"
	"text/template"

	"github.com/compscore/compscore/pkg/checks"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	ConfigFile        string = "config.yml"
	RunningConfigFile string = "running-config.yml"

	RunningConfig *structs.RunningConfig_s
)

func Init() {
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

func RegenerateConfiguration() {
	_, runningConfig, err := GenerateIntialConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to generate running config")
	}

	RunningConfig = runningConfig
}

func UpdateConfiguration() {
	exists, err := FileExists(RunningConfigFile)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to check if config file exists")
	}

	if !exists {
		_, runningConfig, err := GenerateIntialConfig()
		if err != nil {
			logrus.WithError(err).Fatal("Failed to generate running config")
		}

		RunningConfig = runningConfig
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
	for _, team := range RunningConfig.Teams {
		for _, check := range team.Checks {
			_, err = checks.GetCheckFunction(check.Release.Org, check.Release.Repo, check.Release.Tag)
			if err != nil {
				logrus.WithError(err).Fatalf("Failed to load check; %s: %s", check.Name, check.Release.Org+"/"+check.Release.Repo+"@"+check.Release.Tag)
			}
		}
	}
}

func GenerateIntialConfig() (*structs.Config_s, *structs.RunningConfig_s, error) {
	var (
		runningConfig *structs.RunningConfig_s = &structs.RunningConfig_s{}
		config        *structs.Config_s        = &structs.Config_s{}

		name      string
		web_s     structs.Web_s
		teams_s   structs.Teams_s
		scoring_s structs.Scoring_s
		engine_s  structs.Engine_s
		checks_s  []structs.Check_s

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

	err = viper.UnmarshalKey("web", &web_s)
	if err != nil {
		return config, runningConfig, err
	}
	runningConfig.Web = web_s
	config.Web = web_s

	err = viper.UnmarshalKey("teams", &teams_s)
	if err != nil {
		return config, runningConfig, err
	}
	runningConfig.Teams = teams
	config.Teams = teams_s

	err = viper.UnmarshalKey("scoring", &scoring_s)
	if err != nil {
		return config, runningConfig, err
	}

	runningConfig.Scoring = scoring_s
	config.Scoring = scoring_s

	err = viper.UnmarshalKey("engine", &engine_s)
	if err != nil {
		return config, runningConfig, err
	}
	runningConfig.Engine = engine_s
	config.Engine = engine_s

	err = viper.UnmarshalKey("checks", &checks_s)
	if err != nil {
		return config, runningConfig, err
	}
	config.Checks = checks_s

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
			Number:   int8(i + 1),
			Password: teams_s.Password,
		}

		_checks := []structs.Check_s{}

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

			_, tag, err := checks.GetReleaseAssetWithTag(check.Release.Org, check.Release.Repo, check.Release.Tag)
			if err != nil {
				return config, runningConfig, err
			}

			_checks = append(_checks, structs.Check_s{
				Name: check.Name,
				Release: structs.Release_s{
					Org:  check.Release.Org,
					Repo: check.Release.Repo,
					Tag:  tag,
				},
				Credentials: structs.Credentials_s{
					Username: check.Credentials.Username,
					Password: check.Credentials.Password,
				},
				Command:        check.Command,
				Target:         target.String(),
				ExpectedOutput: check.ExpectedOutput,
				Weight:         check.Weight,
			})

			team.Checks = _checks
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
