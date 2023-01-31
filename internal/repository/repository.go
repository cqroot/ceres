package repository

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/adrg/xdg"
)

var once sync.Once

func RootDir() (string, error) {
	repoRootDir := filepath.Join(xdg.DataHome, "ceres")

	var err error
	once.Do(func() {
		err = os.MkdirAll(repoRootDir, os.ModePerm)
	})

	return repoRootDir, err
}

func RepoDir(repoName string) (string, error) {
	rootDir, err := RootDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(rootDir, repoName), nil
}

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
