package config

type Web_s struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	APIPath  string `yaml:"apiPath"`
}

type Teams_s struct {
	Amount     int    `yaml:"amount"`
	NameFormat string `yaml:"nameFormat"`
	Password   string `yaml:"password"`
}

type Scoring_s struct {
	Interval int `yaml:"interval"`
	Retries  int `yaml:"retries"`
}

type Engine_s struct {
	Socket  string `yaml:"socket"`
	Timeout int    `yaml:"timeout"`
}

type Git_s struct {
	Remote string `yaml:"remote"`
	Branch string `yaml:"branch"`
}

type Credentials_s struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Check_s struct {
	Name           string        `yaml:"name"`
	Git            Git_s         `yaml:"git"`
	Credentials    Credentials_s `yaml:"credentials"`
	Target         string        `yaml:"target"`
	Port           int           `yaml:"port"`
	Command        string        `yaml:"command"`
	ExpectedOutput string        `yaml:"expectedOutput"`
	Weight         int           `yaml:"weight"`
}
