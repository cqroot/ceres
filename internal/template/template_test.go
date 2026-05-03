package template

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderString(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
		env      Env
	}{
		{
			name:     "simple variable",
			content:  "module {{ .package_name }}",
			expected: "module testpkg",
			env:      Env{"package_name": "testpkg"},
		},
		{
			name:     "multiple variables",
			content:  "{{ .project_name }} by {{ .author }}",
			expected: "testproj by Test Author",
			env:      Env{"project_name": "testproj", "author": "Test Author"},
		},
		{
			name:     "year function",
			content:  "Copyright (c) {{ year }} {{ .author }}",
			expected: "Copyright (c) 2026 Test Author",
			env:      Env{"author": "Test Author"},
		},
		{
			name:     "if eq condition - true",
			content:  "{{ if eq .license \"MIT\" }}MIT License{{ end }}",
			expected: "MIT License",
			env:      Env{"license": "MIT"},
		},
		{
			name:     "if eq condition - false",
			content:  "{{ if eq .license \"MIT\" }}MIT{{ else }}Other{{ end }}",
			expected: "Other",
			env:      Env{"license": "Apache"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RenderString(tt.content, tt.env)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRenderStringNoTemplate(t *testing.T) {
	env := Env{"key": "value"}
	content := "no template here"
	result := RenderString(content, env)
	assert.Equal(t, content, result)
}

func TestRenderStringInvalidTemplate(t *testing.T) {
	env := Env{"key": "value"}
	content := "{{ .invalid }"
	result := RenderString(content, env)
	assert.Equal(t, content, result)
}

func TestRenderDir(t *testing.T) {
	srcDir := t.TempDir()
	destDir := t.TempDir()

	err := os.MkdirAll(filepath.Join(srcDir, "subdir"), 0755)
	assert.NoError(t, err)

	srcFiles := map[string]string{
		"file1.txt":    "Hello {{ .name }}",
		"subdir/file2": "Value: {{ .value }}",
	}
	for path, content := range srcFiles {
		fullPath := filepath.Join(srcDir, path)
		err := os.MkdirAll(filepath.Dir(fullPath), 0755)
		assert.NoError(t, err)
		err = os.WriteFile(fullPath, []byte(content), 0644)
		assert.NoError(t, err)
	}

	env := Env{"name": "World", "value": "42"}
	err = RenderDir(srcDir, destDir, env)
	assert.NoError(t, err)

	expectedFiles := map[string]string{
		"file1.txt":    "Hello World",
		"subdir/file2": "Value: 42",
	}
	for path, expected := range expectedFiles {
		fullPath := filepath.Join(destDir, path)
		content, err := os.ReadFile(fullPath)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(content))
	}
}

func TestRenderDirPreservesDirectories(t *testing.T) {
	srcDir := t.TempDir()
	destDir := t.TempDir()

	subdirPath := filepath.Join(srcDir, "a", "b", "c")
	err := os.MkdirAll(subdirPath, 0755)
	assert.NoError(t, err)

	err = os.WriteFile(filepath.Join(subdirPath, "file.txt"), []byte("content"), 0644)
	assert.NoError(t, err)

	err = RenderDir(srcDir, destDir, Env{})
	assert.NoError(t, err)

	info, err := os.Stat(filepath.Join(destDir, "a", "b", "c"))
	assert.NoError(t, err)
	assert.True(t, info.IsDir())
}

func TestRenderFile(t *testing.T) {
	srcDir := t.TempDir()
	destDir := t.TempDir()

	srcPath := filepath.Join(srcDir, "test.txt")
	err := os.WriteFile(srcPath, []byte("Hello {{ .name }}"), 0644)
	assert.NoError(t, err)

	destPath := filepath.Join(destDir, "test.txt")
	err = RenderFile(srcPath, destPath, Env{"name": "World"})
	assert.NoError(t, err)

	content, err := os.ReadFile(destPath)
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", string(content))
}

func TestRenderFileWithYearFunction(t *testing.T) {
	srcDir := t.TempDir()
	destDir := t.TempDir()

	srcPath := filepath.Join(srcDir, "license.txt")
	err := os.WriteFile(srcPath, []byte("Copyright {{ year }}"), 0644)
	assert.NoError(t, err)

	destPath := filepath.Join(destDir, "license.txt")
	err = RenderFile(srcPath, destPath, Env{})
	assert.NoError(t, err)

	content, err := os.ReadFile(destPath)
	assert.NoError(t, err)
	assert.Equal(t, "Copyright 2026", string(content))
}