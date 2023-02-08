package repoconf

import (
	"os"
	"path/filepath"

	"github.com/cqroot/prompt"
	"github.com/pelletier/go-toml/v2"
)

func NewFromToml(tomlPath string) (*RepoConf, error) {
	bs, err := os.ReadFile(tomlPath)
	if err != nil {
		return nil, err
	}

	rootDir, err := filepath.Abs(filepath.Dir(tomlPath))
	if err != nil {
		return nil, err
	}

	var rc RepoConf

	err = toml.Unmarshal(bs, &rc)
	if err != nil {
		return nil, err
	}

	// relpath to abspath
	for i := 0; i < len(rc.Scripts.AfterScripts); i++ {
		path := &rc.Scripts.AfterScripts[i]
		if !filepath.IsAbs(*path) {
			*path = filepath.Join(rootDir, *path)
		}
	}

	return &rc, nil
}

func Data(rc *RepoConf) (map[string]string, error) {
	ppt := prompt.New()
	data := make(map[string]string)

	var val string
	var err error

	proj, err := prompt.New().Ask("Your project name:").Input("project")
	if err != nil {
		return nil, err
	}
	data["project_name"] = proj

	for _, varName := range rc.Common.Variables {
		variable := rc.Variable[varName]
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

		data[varName] = val
	}

	return data, nil
}
