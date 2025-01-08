package fstree

import (
	"os"
	"path/filepath"
)

type FileInfo struct {
	Name    string
	RelPath string
	IsDir   bool
	Mode    os.FileMode
}

func FileInfos(path string) ([]FileInfo, error) {
	files := make([]FileInfo, 0)

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(path, p)
		if err != nil {
			return err
		}

		files = append(files, FileInfo{
			Name:    info.Name(),
			RelPath: filepath.ToSlash(relPath),
			IsDir:   info.IsDir(),
			Mode:    info.Mode(),
		})
		return nil
	})

	return files, err
}
