/*
Copyright (C) 2025 Keith Chu <cqroot@outlook.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
