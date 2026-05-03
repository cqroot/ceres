package cmd

import (
	"fmt"
	"os"
	"strings"

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
			fmt.Fprintf(os.Stderr, "Error: invalid format '%s', expected USER/REPO_NAME\n", arg)
			continue
		}
		user := parts[0]
		repoName := parts[1]

		repoPath := ceres.GetRepoCachePath(user, repoName)
		cloneURL := fmt.Sprintf("https://github.com/%s/%s", user, repoName)

		if _, err := os.Stat(repoPath); err == nil {
			fmt.Printf("Updating: %s/%s\n", user, repoName)
			if err := os.RemoveAll(repoPath); err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to remove existing repo: %v\n", err)
				continue
			}
		} else {
			fmt.Printf("Downloading: %s/%s\n", user, repoName)
		}

		if err := ceres.GitClone(repoPath, cloneURL); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}

		fmt.Printf("Added: %s/%s\n", user, repoName)
		added++
	}

	if added > 0 {
		fmt.Printf("\nAdded %d repos.\n", added)
	}
}