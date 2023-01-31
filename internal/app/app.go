package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cqroot/ceres/internal/repository"
	"github.com/cqroot/ceres/internal/script"
	"github.com/cqroot/ceres/internal/templater"
	"github.com/cqroot/ceres/internal/toml"
	"github.com/cqroot/prompt"
)

func Run(repo string) error {
	tomlPath, err := repository.TomlPath(repo)
	if err != nil {
		return err
	}

	proj, err := prompt.New().Ask("Your project name:").Input("project")
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	outputDir := filepath.Join(cwd, proj)

	rootDir := filepath.Dir(tomlPath)

	templateDir := filepath.Join(rootDir, "template")
	if !filepath.IsAbs(outputDir) {
		outputDir = filepath.Join(rootDir, outputDir)
	}

	co, vars, err := getTomlData(tomlPath)
	if err != nil {
		return err
	}

	vars["project_name"] = proj

	tmpl := templater.New(
		templateDir, outputDir, vars, co.IncludePathRules, co.ExcludePathRules)

	fmt.Println()
	fmt.Println("Template path :", templateDir)
	fmt.Println("Output path   :", outputDir)
	fmt.Printf("Variables     : %+v\n", vars)
	fmt.Println()

	err = tmpl.Execute()
	if err != nil {
		return err
	}

	fmt.Println("")

	for _, scriptPath := range co.Scripts.AfterScripts {
		err = script.Run(scriptPath, outputDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func getTomlData(tomlPath string) (*toml.ConfigObject, map[string]string, error) {
	p, err := toml.New(tomlPath)
	if err != nil {
		return nil, nil, err
	}

	co, err := p.Parse()
	if err != nil {
		return nil, nil, err
	}

	ret, err := p.Run(co)
	if err != nil {
		return nil, nil, err
	}

	return co, ret, nil
}
