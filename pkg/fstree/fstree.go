package fstree

import (
	"os"
	"path/filepath"
)

func ShouldSkipPath(relPath string, info os.FileInfo) bool {
	if relPath == "." {
		return true
	}

	if info.IsDir() {
		return true
	}

	return false
}

type FileInfo struct {
	Name    string
	RelPath string
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

		if ShouldSkipPath(relPath, info) {
			return nil
		}

		files = append(files, FileInfo{
			Name:    info.Name(),
			RelPath: filepath.ToSlash(relPath),
		})
		return nil
	})

	return files, err
}
