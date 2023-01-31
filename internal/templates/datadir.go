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
