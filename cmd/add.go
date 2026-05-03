package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cqroot/ceres/internal/logger"
	"github.com/cqroot/ceres/pkg/ceres"
	"github.com/spf13/cobra"
)

func newAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add USER/REPO_NAME [USER/REPO_NAME...]",
		Short: "Download or update repos from GitHub",
		Args:  cobra.MinimumNArgs(1),
		Run:   runAdd,
	}
}

func runAdd(cmd *cobra.Command, args []string) {
	var added int
	for _, arg := range args {
		parts := strings.Split(arg, "/")
		if len(parts) != 2 {
			logger.Errorf("invalid format '%s', expected USER/REPO_NAME", arg)
			continue
		}
		user := parts[0]
		repoName := parts[1]

		repoPath := ceres.GetRepoCachePath(user, repoName)
		cloneURL := fmt.Sprintf("https://github.com/%s/%s", user, repoName)

		if _, err := os.Stat(repoPath); err == nil {
			logger.Infof("Updating: %s/%s", user, repoName)
			if err := os.RemoveAll(repoPath); err != nil {
				logger.Errorf("failed to remove existing repo: %v", err)
				continue
			}
		} else {
			logger.Infof("Downloading: %s/%s", user, repoName)
		}

		if err := ceres.GitClone(repoPath, cloneURL); err != nil {
			logger.Errorf("%v", err)
			continue
		}

		logger.Infof("Added: %s/%s", user, repoName)
		added++
	}

	if added > 0 {
		logger.Infof("Added %d repos.", added)
	}
}
