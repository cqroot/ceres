package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cqroot/prompt"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/cqroot/ceres/internal/repoconf"
	"github.com/cqroot/ceres/internal/repository"
	"github.com/cqroot/ceres/internal/script"
	"github.com/cqroot/ceres/internal/templater"
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

	rc, data, err := repoconfAndData(tomlPath)
	if err != nil {
		return err
	}

	data["project_name"] = proj

	tmpl := templater.New(
		templateDir, outputDir, data, rc.IncludePathRules, rc.ExcludePathRules)

	fmt.Println()
	fmt.Println(text.FgCyan.Sprint("Repository :"), templateDir)
	fmt.Println(text.FgCyan.Sprint("Project    :"), outputDir)
	fmt.Printf("%s %+v\n", text.FgCyan.Sprint("Variables  :"), data)
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
