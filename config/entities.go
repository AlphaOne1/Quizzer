package config

type Root struct {
	ListenAddress string `json:"listenAddress"`
	ListenPort    string `json:"listenPort"`
	Tasks         []Task `json:"tasks"`
}

type Task struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"isCorrect"`
}
