package repository

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/adrg/xdg"
)

var once sync.Once

// RootDir returns the root directory of the repository.
func RootDir() (string, error) {
	repoRootDir := filepath.Join(xdg.DataHome, "ceres")

	var err error
	once.Do(func() {
		err = os.MkdirAll(repoRootDir, os.ModePerm)
	})

	return repoRootDir, err
}

// RepoDir returns the path of the given repository.
func RepoDir(repo string) (string, error) {
	rootDir, err := RootDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(rootDir, repo), nil
}

func TemplateDir(repo string) (string, error) {
	repoDir, err := RepoDir(repo)
	if err != nil {
		return "", err
	}

	return filepath.Join(repoDir, "template"), nil
}

// TomlPath returns the toml file path of the given repository.
func TomlPath(repo string) (string, error) {
	repoDir, err := RepoDir(repo)
	if err != nil {
		return "", err
	}

	return filepath.Join(repoDir, "ceres.toml"), nil
}

// Repos returns a list of all repositories.
func Repos() ([]string, error) {
	rootDir, err := RootDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	repos := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			repos = append(repos, entry.Name())
		}
	}

	return repos, err
}
