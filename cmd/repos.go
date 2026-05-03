package cmd

import (
	"github.com/cqroot/ceres/internal/logger"
	"github.com/cqroot/ceres/pkg/ceres"
	"github.com/spf13/cobra"
)

func newReposCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "repos",
		Short: "List downloaded repos",
		Run:   runRepos,
	}
}

func runRepos(cmd *cobra.Command, args []string) {
	dataDir := ceres.GetDataDir()
	logger.Info("Data Directory: ", dataDir)

	repos, err := ceres.ListRepos()
	if err != nil {
		logger.Info("No repos found.")
		return
	}

	logger.Info("Repos:")
	for _, repo := range repos {
		logger.Infof("  %s/%s", repo.User, repo.Name)
	}
}
