package prompting

import (
	"fmt"
	"github.com/cqroot/ceres/pkg/logging"
	"github.com/cqroot/prompt"
)

type Type string

const (
	TypeInput  Type = "input"
	TypeChoose Type = "choose"
)

type Prompting struct {
	Name    string `yaml:"name"`
	Type    Type   `yaml:"type"`
	Message string `yaml:"message"`
	Default string `yaml:"default"`
}

func Prompt(promptings []Prompting, logger logging.Logger) (map[string]any, error) {
	vars := make(map[string]any)
	for _, prompting := range promptings {
		switch prompting.Type {
		case TypeInput:
			val, err := prompt.New().Ask(prompting.Message).Input(prompting.Default)
			if err != nil {
				return nil, err
			}
			vars[prompting.Name] = val
		case TypeChoose:
			val, err := prompt.New().Ask(prompting.Message).Input(prompting.Default)
			if err != nil {
				return nil, err
			}
			vars[prompting.Name] = val
		}
	}
	logger.Debug().Str("vars", fmt.Sprintf("%v", vars)).Msg("Read the variables.")
	return vars, nil
}
