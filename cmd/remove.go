package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
			fmt.Fprintf(os.Stderr, "Error: invalid format '%s', expected USER/REPO_NAME\n", arg)
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
			fmt.Fprintf(os.Stderr, "Error: failed to remove %s/%s: %v\n", user, repoName, err)
			continue
		}

		parentDir := filepath.Dir(repoPath)
		if entries, err := os.ReadDir(parentDir); err == nil && len(entries) == 0 {
			if err := os.Remove(parentDir); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: failed to remove empty parent dir %s: %v\n", parentDir, err)
			}
		}

		fmt.Printf("Removed: %s/%s\n", user, repoName)
		removed++
	}

	if removed > 0 {
		fmt.Printf("\nRemoved %d repos.\n", removed)
	}
}
