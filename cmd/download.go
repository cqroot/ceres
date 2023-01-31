package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cqroot/sawmill/internal/templates"
	"github.com/cqroot/sawmill/internal/utils"
)

func init() {
	rootCmd.AddCommand(downloadCmd)
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download the template from the git repository",
	Long:  "Download the template from the git repository",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:   runDownloadCmd,
}

func runDownloadCmd(cmd *cobra.Command, args []string) {
	dataDir, err := templates.DataDir()
	cobra.CheckErr(err)

	gitArgs := []string{"-C", dataDir, "clone", args[0]}
	fmt.Println("git", gitArgs)

	err = utils.ExecCmd("git", gitArgs...)
	cobra.CheckErr(err)
}
