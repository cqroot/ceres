package prompts

import (
	"os"

	"github.com/cqroot/prompt"
	"github.com/pelletier/go-toml/v2"
)

type Prompts struct {
	tomlPath string
}

func New(tomlPath string) *Prompts {
	return &Prompts{
		tomlPath: tomlPath,
	}
}

func (p *Prompts) Parse() (*ConfigObject, error) {
	bs, err := os.ReadFile(p.tomlPath)
	if err != nil {
		return nil, err
	}

	var co ConfigObject

	err = toml.Unmarshal(bs, &co)
	if err != nil {
		return nil, err
	}
	return &co, nil
}

func (p *Prompts) Run(co *ConfigObject) (map[string]string, error) {
	ppt := prompt.New()
	ret := make(map[string]string)
	var val string
	var err error

	for _, varName := range co.CommonItem.Variables {
		variable := co.Variable[varName]
		switch variable.Type {
		case "input":
			val, err = ppt.Ask(variable.Message).Input(variable.Meta[0])
			if err != nil {
				return nil, err
			}

		case "toggle":
			val, err = ppt.Ask(variable.Message).Toggle(variable.Meta)
			if err != nil {
				return nil, err
			}

		case "choose":
			val, err = ppt.Ask(variable.Message).Choose(variable.Meta)
			if err != nil {
				return nil, err
			}
		}

		ret[varName] = val
	}

	return ret, nil
}
