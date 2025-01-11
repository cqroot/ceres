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

package cmd

import (
	"github.com/cqroot/ceres/internal/app"
	"github.com/cqroot/ceres/pkg/logging"
	"github.com/spf13/cobra"
)

func RunRootCmd(cmd *cobra.Command, args []string) {
	repoPath := args[0]
	logger := logging.New()
	cobra.CheckErr(app.Apply(repoPath, logger))
}

func NewRootCmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "ceres",
		Short: "Ceres - Manage your project templates",
		Long:  "Ceres - Manage your project templates",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run:   RunRootCmd,
	}

	return &rootCmd
}

func Execute() {
	err := NewRootCmd().Execute()
	cobra.CheckErr(err)
}
