package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cqroot/ceres/internal/templates"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Sawmill",
	Long:  "Print the version number of Sawmill",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sawmill v0.0.0")
		fmt.Println()

		dataDir, err := templates.DataDir()
		cobra.CheckErr(err)

		fmt.Println("Your templates are stored in", dataDir)
	},
}
