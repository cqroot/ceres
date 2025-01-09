package variables

import (
	"github.com/cqroot/prompt"
)

type VarType int

const (
	VarTypeInput VarType = iota
	VarTypeChoose
)

type VarPrompting struct {
	Name    string
	Type    VarType
	Message string
	Default string
}

func Prompt(promptings []VarPrompting) (map[string]any, error) {
	vars := make(map[string]any)
	for _, prompting := range promptings {
		switch prompting.Type {
		case VarTypeInput:
			val, err := prompt.New().Ask(prompting.Message).Input(prompting.Default)
			if err != nil {
				return nil, err
			}
			vars[prompting.Name] = val
		case VarTypeChoose:
			val, err := prompt.New().Ask(prompting.Message).Input(prompting.Default)
			if err != nil {
				return nil, err
			}
			vars[prompting.Name] = val
		}
	}
	return vars, nil
}
