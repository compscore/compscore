package structs

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

type Team_s struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Number   int8   `yaml:"number"`
}

type Config_s struct {
	Name    string    `yaml:"name"`
	Web     Web_s     `yaml:"web"`
	Teams   Teams_s   `yaml:"teams"`
	Scoring Scoring_s `yaml:"scoring"`
	Engine  Engine_s  `yaml:"engine"`
	Checks  []Check_s `yaml:"checks"`
}

type RunningConfig_s struct {
	Name    string    `yaml:"name"`
	Web     Web_s     `yaml:"web"`
	Scoring Scoring_s `yaml:"scoring"`
	Engine  Engine_s  `yaml:"engine"`
	Teams   []Team_s  `yaml:"teams"`
}
