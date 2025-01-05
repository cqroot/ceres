package tmpl

import (
	"os"
	"path/filepath"
	"text/template"
)

func Execute(tmplPath string, outputPath string, variables map[string]any) error {
	t, err := template.New(filepath.Base(tmplPath)).Funcs(FuncMap).ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	info, err := os.Stat(tmplPath)
	if err != nil {
		return err
	}

	outputFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return t.Execute(outputFile, variables)
}
