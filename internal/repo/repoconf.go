package repo

import (
	"os"

	"github.com/cqroot/ceres/internal/prompting"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Promptings []prompting.Prompting `yaml:"promptings"`
}

func readConfig(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := Config{}
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
