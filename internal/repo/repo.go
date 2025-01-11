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

package repo

import (
	"path"

	"github.com/cqroot/ceres/internal/prompting"
	"github.com/cqroot/ceres/pkg/fstree"
	"github.com/cqroot/ceres/pkg/logging"
)

type Repo struct {
	repoPath string
	confPath string
	skelPath string
	vars     map[string]any
	logger   logging.Logger
}

func New(repoPath string, logger logging.Logger) *Repo {
	r := Repo{
		repoPath: repoPath,
		confPath: path.Join(repoPath, "ceres.yaml"),
		skelPath: path.Join(repoPath, "skeleton"),
		vars:     nil,
		logger:   logger,
	}

	return &r
}

func (r *Repo) SkelPath() string {
	return r.skelPath
}

func (r *Repo) Read() error {
	conf, err := readConfig(r.confPath)
	if err != nil {
		r.logger.Err(err).Msg("Failed to read config.")
		return err
	}
	r.logger.Debug().Any("repoConf", conf).Msg("Read conf.")

	r.vars, err = prompting.Prompt(conf.Promptings)
	if err != nil {
		r.logger.Err(err).Msg("Failed to prompt.")
		return err
	}
	r.logger.Debug().Any("vars", r.vars).Msg("Read vars.")

	return nil
}

func (r *Repo) Vars() map[string]any {
	return r.vars
}

func (r *Repo) SkelFileInfos() ([]fstree.FileInfo, error) {
	skelFileInfos, err := fstree.FileInfos(r.skelPath)
	if err != nil {
		r.logger.Err(err).Msg("Failed to list files.")
		return nil, err
	}
	return skelFileInfos, nil
}
