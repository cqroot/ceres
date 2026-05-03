package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/cqroot/ceres/internal/logger"
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

		renderedPath, _ := RenderString(relPath, env)
		destPath := filepath.Join(destDir, renderedPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, 0o755)
		}

		return RenderFile(path, destPath, env)
	})
}

func RenderFile(srcPath, destPath string, env Env) error {
	content, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", srcPath, err)
	}

	rendered, err := RenderString(string(content), env)
	if err != nil {
		logger.Warnf("failed to render template file %s: %v", srcPath, err)
		return err
	}

	return os.WriteFile(destPath, []byte(rendered), 0o644)
}

func RenderString(content string, env Env) (string, error) {
	tmpl := template.New("").Funcs(template.FuncMap{
		"year": func() string {
			return fmt.Sprintf("%d", time.Now().Year())
		},
	})
	tmpl, err := tmpl.Parse(content)
	if err != nil {
		return content, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, env); err != nil {
		return content, err
	}

	return strings.TrimSpace(buf.String()), nil
}
