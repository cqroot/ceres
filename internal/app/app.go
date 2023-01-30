package app

import (
	"fmt"
	"path/filepath"

	"github.com/cqroot/sawmill/internal/prompts"
	"github.com/cqroot/sawmill/internal/templater"
)

func Run(tomlPath string, outputDir string) error {
	rootDir := filepath.Dir(tomlPath)

	p := prompts.New(tomlPath)

	co, err := p.Parse()
	if err != nil {
		return err
	}

	ret, err := p.Run(co)
	if err != nil {
		return err
	}

	templateDir := filepath.Join(rootDir, "template")
	if !filepath.IsAbs(outputDir) {
		outputDir = filepath.Join(rootDir, outputDir)
	}

	tmpl := templater.New(
		templateDir, outputDir, ret, co.IncludePathRules, co.ExcludePathRules)

	fmt.Println()
	fmt.Println("Template path :", templateDir)
	fmt.Println("Output path   :", outputDir)
	fmt.Printf("Variables     : %+v\n", ret)
	fmt.Println()

	err = tmpl.Execute()
	if err != nil {
		return err
	}

	fmt.Println("")

	return nil
}
