package toml

import (
	"os"
	"path/filepath"

	"github.com/cqroot/prompt"
	"github.com/pelletier/go-toml/v2"
)

type Toml struct {
	tomlPath string
	rootDir  string
}

func New(tomlPath string) (*Toml, error) {
	rootDir, err := filepath.Abs(filepath.Dir(tomlPath))
	if err != nil {
		return nil, err
	}

	return &Toml{
		tomlPath: tomlPath,
		rootDir:  rootDir,
	}, nil
}

func (p *Toml) Parse() (*ConfigObject, error) {
	bs, err := os.ReadFile(p.tomlPath)
	if err != nil {
		return nil, err
	}

	var co ConfigObject

	err = toml.Unmarshal(bs, &co)
	if err != nil {
		return nil, err
	}

	// relpath to abspath
	for i := 0; i < len(co.Scripts.AfterScripts); i++ {
		path := &co.Scripts.AfterScripts[i]
		if !filepath.IsAbs(*path) {
			*path = filepath.Join(p.rootDir, *path)
		}
	}

	return &co, nil
}

func (p *Toml) Run(co *ConfigObject) (map[string]string, error) {
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
