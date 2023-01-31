package cmd

import (
	"github.com/cqroot/sawmill/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sawmill",
	Short: "sawmill",
	Long:  "sawmill",
	Run:   runRootCmd,
}

func runRootCmd(cmd *cobra.Command, args []string) {
	cobra.CheckErr(app.Run())
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
