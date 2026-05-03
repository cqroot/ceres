package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cqroot/ceres/internal/logger"
	"github.com/cqroot/ceres/pkg/ceres"
	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove USER/REPO_NAME [USER/REPO_NAME...]",
		Short: "Remove downloaded repos",
		Args:  cobra.MinimumNArgs(1),
		Run:   runRemove,
	}
}

func runRemove(cmd *cobra.Command, args []string) {
	var removed int
	for _, arg := range args {
		parts := strings.Split(arg, "/")
		if len(parts) != 2 {
			logger.Errorf("invalid format '%s', expected USER/REPO_NAME", arg)
			continue
		}
		user := parts[0]
		repoName := parts[1]

		repoPath := ceres.GetRepoCachePath(user, repoName)
		if _, err := os.Stat(repoPath); os.IsNotExist(err) {
			fmt.Printf("Repo not found: %s/%s\n", user, repoName)
			continue
		}

		if err := os.RemoveAll(repoPath); err != nil {
			logger.Errorf("failed to remove %s/%s: %v", user, repoName, err)
			continue
		}

		parentDir := filepath.Dir(repoPath)
		if entries, err := os.ReadDir(parentDir); err == nil && len(entries) == 0 {
			if err := os.Remove(parentDir); err != nil {
				logger.Warnf("failed to remove empty parent dir %s: %v", parentDir, err)
			}
		}

		fmt.Printf("Removed: %s/%s\n", user, repoName)
		removed++
	}

	if removed > 0 {
		fmt.Printf("\nRemoved %d repos.\n", removed)
	}
}
