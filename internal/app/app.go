package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cqroot/ceres/internal/repoconf"
	"github.com/cqroot/ceres/internal/repository"
	"github.com/cqroot/ceres/internal/script"
	"github.com/cqroot/ceres/internal/templater"
	"github.com/cqroot/ceres/internal/utils"
)

func repoconfAndData(tomlPath string) (*repoconf.RepoConf, map[string]string, error) {
	rc, err := repoconf.ParseToml(tomlPath)
	if err != nil {
		return nil, nil, err
	}

	data, err := repoconf.Data(rc)
	if err != nil {
		return nil, nil, err
	}

	return rc, data, nil
}

func Run(repo string) error {
	tomlPath, err := repository.TomlPath(repo)
	if err != nil {
		return err
	}

	rc, data, err := repoconfAndData(tomlPath)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	outputDir := filepath.Join(cwd, data["project_name"])
	rootDir := filepath.Dir(tomlPath)

	templateDir := filepath.Join(rootDir, "template")
	if !filepath.IsAbs(outputDir) {
		outputDir = filepath.Join(rootDir, outputDir)
	}

	tmpl := templater.New(
		templateDir, outputDir, data, rc.IncludePathRules, rc.ExcludePathRules)

	fmt.Println()
	fmt.Println(utils.ColorString("Repository :"), templateDir)
	fmt.Println(utils.ColorString("Project    :"), outputDir)
	fmt.Printf("%s %+v\n", utils.ColorString("Variables  :"), data)
	fmt.Println()

	err = tmpl.Execute()
	if err != nil {
		return err
	}

	fmt.Println("")

	for _, scriptPath := range rc.Scripts.AfterScripts {
		err = script.Run(scriptPath, outputDir)
		if err != nil {
			return err
		}
	}

	return nil
}
