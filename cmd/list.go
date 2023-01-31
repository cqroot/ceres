package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cqroot/ceres/internal/repository"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all downloaded templates",
	Long:    "List all downloaded templates",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := repository.Repos()
		cobra.CheckErr(err)

		for _, repo := range repos {
			fmt.Println(repo)
		}
	},
}
