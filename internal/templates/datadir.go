package templates

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/adrg/xdg"
)

var once sync.Once

func DataDir() (string, error) {
	dataDir := filepath.Join(xdg.DataHome, "ceres")

	var err error
	once.Do(func() {
		err = os.MkdirAll(dataDir, os.ModePerm)
	})

	return dataDir, err
}

func Templates() ([]string, error) {
	dataDir, err := DataDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, err
	}

	choices := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			choices = append(choices, entry.Name())
		}
	}

	return choices, err
}
