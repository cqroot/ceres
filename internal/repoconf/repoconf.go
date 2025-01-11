package repoconf

import (
	"os"

	"github.com/cqroot/ceres/internal/prompting"
	"gopkg.in/yaml.v3"
)

type RepoConf struct {
	Promptings []prompting.Prompting `yaml:"promptings"`
}

func Read(path string) (*RepoConf, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := RepoConf{}
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
