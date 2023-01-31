package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cqroot/ceres/internal/templates"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all downloaded templates",
	Long:  "List all downloaded templates",
	Run: func(cmd *cobra.Command, args []string) {
		templates, err := templates.Templates()
		cobra.CheckErr(err)

		for _, template := range templates {
			fmt.Println(template)
		}
	},
}
