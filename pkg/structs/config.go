package structs

import "time"

type Web_s struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	JWTKey   string `yaml:"jwtKey"`
	Timeout  int    `yaml:"timeout"`
	Release  bool   `yaml:"release"`
}

type Redis_s struct {
	Url           string        `yaml:"url"`
	Password      string        `yaml:"password"`
	FastRefresh   time.Duration `yaml:"fastRefresh"`
	MediumRefresh time.Duration `yaml:"mediumRefresh"`
	SlowRefresh   time.Duration `yaml:"slowRefresh"`
}

type Teams_s struct {
	Amount     int    `yaml:"amount"`
	NameFormat string `yaml:"nameFormat"`
	Password   string `yaml:"password"`
}

type Scoring_s struct {
	Interval int `yaml:"interval"`
}

type Engine_s struct {
	Socket  string `yaml:"socket"`
	Timeout int    `yaml:"timeout"`
}

type Release_s struct {
	Org  string `yaml:"org"`
	Repo string `yaml:"repo"`
	Tag  string `yaml:"tag"`
}

type Credentials_s struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Check_s struct {
	Name           string                 `yaml:"name"`
	Release        Release_s              `yaml:"release"`
	Credentials    Credentials_s          `yaml:"credentials"`
	Target         string                 `yaml:"target"`
	Command        string                 `yaml:"command"`
	ExpectedOutput string                 `yaml:"expectedOutput"`
	Weight         int                    `yaml:"weight"`
	Options        map[string]interface{} `yaml:"options"`
}

type AdminUser_s struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
