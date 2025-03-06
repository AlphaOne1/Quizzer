package config

type Root struct {
	ListenAddress string `yaml:"listenAddress"`
	ListenPort    int    `yaml:"listenPort"`
	Tasks         []Task `yaml:"tasks"`
}

type Task struct {
	Question string   `yaml:"question"`
	Answers  []Answer `yaml:"answers"`
}

type Answer struct {
	Text      string `yaml:"text"`
	IsCorrect bool   `yaml:"isCorrect"`
}
