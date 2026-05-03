package ceres

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

var overrideDataDir string

func SetDataDir(dir string) {
	overrideDataDir = dir
}

type PromptConfig struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Message string   `yaml:"message"`
	Default string   `yaml:"default"`
	Options []string `yaml:"options,omitempty"`
}

type Config struct {
	Promptings []PromptConfig `yaml:"promptings"`
	Before     []string        `yaml:"before,omitempty"`
	After      []string        `yaml:"after,omitempty"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read ceres.yaml: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse ceres.yaml: %w", err)
	}

	return &cfg, nil
}

func GetDataDir() string {
	if overrideDataDir != "" {
		return overrideDataDir
	}
	return filepath.Join(xdg.DataHome, "ceres")
}

func GetRepoCachePath(user, repo string) string {
	return filepath.Join(GetDataDir(), "github.com", user, repo)
}

type Repo struct {
	User string
	Name string
	Path string
}

func ListRepos() ([]Repo, error) {
	ceresDir := filepath.Join(GetDataDir(), "github.com")

	entries, err := os.ReadDir(ceresDir)
	if err != nil {
		return nil, err
	}

	var repos []Repo
	for _, userEntry := range entries {
		if !userEntry.IsDir() {
			continue
		}
		userPath := filepath.Join(ceresDir, userEntry.Name())
		userEntries, err := os.ReadDir(userPath)
		if err != nil {
			continue
		}
		for _, repoEntry := range userEntries {
			if !repoEntry.IsDir() {
				continue
			}
			repos = append(repos, Repo{
				User: userEntry.Name(),
				Name: repoEntry.Name(),
				Path: filepath.Join(userPath, repoEntry.Name()),
			})
		}
	}

	return repos, nil
}
