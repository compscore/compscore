package structs

type Competition struct {
	Server  Server  `mapstructure:"server"`
	Teams   Teams   `mapstructure:"teams"`
	Scoring Scoring `mapstructure:"scoring"`
	Checks  []Check `mapstructure:"checks"`
}

type Server struct {
	APIPath  string `mapstructure:"api-path"`
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
}

type Teams struct {
	Amount     int    `mapstructure:"amount"`
	NameFormat string `mapstructure:"name-format"`
	Password   string `mapstructure:"password"`
}

type Scoring struct {
	Interval int `mapstructure:"interval"`
	Retries  int `mapstructure:"retries"`
}

type Check struct {
	Name           string     `mapstructure:"name"`
	Git            Git        `mapstructure:"git"`
	Credential     Credential `mapstructure:"credential"`
	Target         string     `mapstructure:"target"`
	Port           int        `mapstructure:"port"`
	Command        string     `mapstructure:"command"`
	ExpectedOutput string     `mapstructure:"expected-output"`
	Weight         int        `mapstructure:"weight"`
}

type Git struct {
	Remote string `mapstructure:"remote"`
	Branch string `mapstructure:"branch"`
}

type Credential struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
