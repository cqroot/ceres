package templates

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

func DataDir() string {
	dataDir := filepath.Join(xdg.DataHome, "sawmill")
	return dataDir
}
