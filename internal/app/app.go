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

package app

import (
	"os"
	"path"

	"github.com/cqroot/ceres/internal/repo"
	"github.com/cqroot/ceres/pkg/logging"
	"github.com/cqroot/ceres/pkg/tmpl"
	"github.com/spf13/cobra"
)

func projPath(projName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(cwd, projName), nil
}

func Apply(repoPath string, logger logging.Logger) error {
	r := repo.New(repoPath, logger)
	err := r.Read()
	if err != nil {
		return err
	}
	vars := r.Vars()

	skelFileInfos, err := r.SkelFileInfos()
	if err != nil {
		return err
	}

	projPath, err := projPath(vars["project_name"].(string))
	if err != nil {
		logger.Error().Msg("Could not determine project path.")
		return err
	}
	logger.Debug().Str("projPath", projPath).Msg("Read project path.")

	for _, skelFileInfo := range skelFileInfos {
		if skelFileInfo.RelPath == "." {
			continue
		}

		if skelFileInfo.IsDir {
			dirPath := path.Join(projPath, skelFileInfo.RelPath)
			err = os.MkdirAll(dirPath, skelFileInfo.Mode)
			if err != nil {
				logger.Err(err).
					Str("dir", dirPath).
					Str("mode", skelFileInfo.Mode.String()).
					Msg("Failed to create directory.")
				return err
			}
			continue
		}

		logger.Debug().Str("file", skelFileInfo.RelPath).Msg("Generate file.")
		cobra.CheckErr(tmpl.Execute(
			path.Join(r.SkelPath(), skelFileInfo.RelPath),
			path.Join(projPath, skelFileInfo.RelPath),
			vars,
		))
	}

	return nil
}
