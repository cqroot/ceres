package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cqroot/ceres/internal/app"
	"github.com/cqroot/ceres/internal/repository"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Read config file but do not distribute files",
	Long:  "Read config file but do not distribute files",
	Args:  cobra.MatchAll(cobra.RangeArgs(0, 1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		repo := ""

		if len(args) == 0 {
			var err error
			repo, err = repository.ChooseRepo()
			cobra.CheckErr(err)
		} else {
			repo = args[0]
		}

		tomlPath, err := repository.TomlPath(repo)
		cobra.CheckErr(err)

		err = app.TestConfig(tomlPath)
		cobra.CheckErr(err)
	},
}
