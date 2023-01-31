package cmd

import (
	"github.com/cqroot/ceres/internal/app"
	"github.com/cqroot/ceres/internal/repository"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ceres",
	Short: "ceres",
	Long:  "ceres",
	Args:  cobra.MatchAll(cobra.RangeArgs(0, 1), cobra.OnlyValidArgs),
	Run:   runRootCmd,
}

func runRootCmd(cmd *cobra.Command, args []string) {
	repo := ""

	if len(args) == 0 {
		var err error
		repo, err = repository.ChooseRepo()
		cobra.CheckErr(err)
	} else {
		repo = args[0]
	}

	cobra.CheckErr(app.Run(repo))
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
