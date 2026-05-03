package ceres

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDataDir(t *testing.T) {
	original := overrideDataDir
	defer func() { overrideDataDir = original }()

	overrideDataDir = ""
	dir := GetDataDir()
	assert.NotEmpty(t, dir)
}

func TestGetDataDirWithOverride(t *testing.T) {
	original := overrideDataDir
	defer func() { overrideDataDir = original }()

	overrideDataDir = "/custom/path"
	assert.Equal(t, "/custom/path", GetDataDir())
}

func TestGetRepoCachePath(t *testing.T) {
	original := overrideDataDir
	defer func() { overrideDataDir = original }()

	overrideDataDir = "/data"
	path := GetRepoCachePath("user", "repo")
	assert.Equal(t, filepath.Join("/data", "github.com", "user", "repo"), path)
}

func TestLoadConfig(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "ceres.yaml")

	content := `
promptings:
  - name: test_name
    type: input
    message: "Test message:"
    default: "default_value"

before:
  - echo before

after:
  - echo after
`
	err := os.WriteFile(configPath, []byte(content), 0644)
	assert.NoError(t, err)

	cfg, err := LoadConfig(configPath)
	assert.NoError(t, err)
	assert.Len(t, cfg.Promptings, 1)
	assert.Equal(t, "test_name", cfg.Promptings[0].Name)
	assert.Len(t, cfg.Before, 1)
	assert.Len(t, cfg.After, 1)
}

func TestLoadConfigFileNotFound(t *testing.T) {
	_, err := LoadConfig("/nonexistent/path/ceres.yaml")
	assert.Error(t, err)
}

func TestLoadConfigInvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "ceres.yaml")

	err := os.WriteFile(configPath, []byte("invalid: yaml: content:"), 0644)
	assert.NoError(t, err)

	_, err = LoadConfig(configPath)
	assert.Error(t, err)
}

func TestListReposEmptyDir(t *testing.T) {
	original := overrideDataDir
	defer func() { overrideDataDir = original }()

	tmpDir := t.TempDir()
	overrideDataDir = tmpDir

	githubDir := filepath.Join(tmpDir, "github.com")
	err := os.MkdirAll(githubDir, 0755)
	assert.NoError(t, err)

	repos, err := ListRepos()
	assert.NoError(t, err)
	assert.Len(t, repos, 0)
}

func TestListRepos(t *testing.T) {
	original := overrideDataDir
	defer func() { overrideDataDir = original }()

	tmpDir := t.TempDir()
	overrideDataDir = tmpDir

	userDir := filepath.Join(tmpDir, "github.com", "testuser")
	repoDir := filepath.Join(userDir, "testrepo")
	err := os.MkdirAll(repoDir, 0755)
	assert.NoError(t, err)

	repos, err := ListRepos()
	assert.NoError(t, err)
	assert.Len(t, repos, 1)
	assert.Equal(t, "testuser", repos[0].User)
	assert.Equal(t, "testrepo", repos[0].Name)
}