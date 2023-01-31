package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cqroot/ceres/internal/repository"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Ceres",
	Long:  "Print the version number of Ceres",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ceres v0.0.0")
		fmt.Println()

		rootDir, err := repository.RootDir()
		cobra.CheckErr(err)

		fmt.Println("Your template repositories are stored in:", rootDir)
	},
}
