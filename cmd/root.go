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
