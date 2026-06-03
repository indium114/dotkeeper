package cmd

type State struct {
	Keep  string `json:"keep"`
	Links []Link `json:"links"`
}

type Keep struct {
	Links []Link `yaml:"links"`
}

type Link struct {
	Source string `yaml:"source"`
	Target string `yaml:"target"`
}
