package templater

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func makeParentDirs(dir string) error {
	parentDir := filepath.Dir(dir)
	err := os.MkdirAll(parentDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("make parent dirs for %s: %w", dir, err)
	}

	return nil
}

func ExecuteTemplate(input string, output string, data any) error {
	err := makeParentDirs(output)
	if err != nil {
		return err
	}

	tname := filepath.Base(input)
	tmpl, err := template.New(tname).Funcs(funcMap).ParseFiles(input)
	if err != nil {
		return fmt.Errorf(
			"execute template {name: %s, input:%s}: %w", tname, input, err)
	}

	fOutput, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer fOutput.Close()

	err = tmpl.Execute(fOutput, data)
	return err
}
