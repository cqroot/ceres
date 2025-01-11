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

package tmpl

import (
	"os"
	"path/filepath"
	"text/template"
)

func Execute(tmplPath string, outputPath string, variables map[string]any) error {
	t, err := template.New(filepath.Base(tmplPath)).Funcs(FuncMap).ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	info, err := os.Stat(tmplPath)
	if err != nil {
		return err
	}

	outputFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return t.Execute(outputFile, variables)
}
