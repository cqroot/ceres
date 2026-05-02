package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type Env map[string]string

func RenderDir(srcDir, destDir string, env Env) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		if relPath == "." {
			return nil
		}

		renderedPath := RenderString(relPath, env)
		destPath := filepath.Join(destDir, renderedPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		return RenderFile(path, destPath, env)
	})
}

func RenderFile(srcPath, destPath string, env Env) error {
	content, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", srcPath, err)
	}

	rendered := RenderString(string(content), env)

	return os.WriteFile(destPath, []byte(rendered), 0644)
}

func RenderString(content string, env Env) string {
	tmpl, err := template.New("").Parse(content)
	if err != nil {
		return content
	}

	tmpl = tmpl.Funcs(template.FuncMap{
		"year": func() string {
			return fmt.Sprintf("%d", time.Now().Year())
		},
	})

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, env); err != nil {
		return content
	}

	return strings.TrimSpace(buf.String())
}
