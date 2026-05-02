package cmd

import (
	"fmt"

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
	fmt.Printf("Data Directory: %s\n\n", dataDir)

	repos, err := ceres.ListRepos()
	if err != nil {
		fmt.Printf("No repos found.\n")
		return
	}

	fmt.Printf("Repos:\n")
	for _, repo := range repos {
		fmt.Printf("  %s/%s\n", repo.User, repo.Name)
	}
}
