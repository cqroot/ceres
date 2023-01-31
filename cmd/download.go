package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cqroot/ceres/internal/repository"
	"github.com/cqroot/ceres/internal/utils"
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
	repoUrl := args[0]
	if strings.Count(repoUrl, "/") == 1 {
		repoUrl = "https://github.com/" + repoUrl
	}

	repoDir, err := repository.RepoDir(filepath.Base(repoUrl))
	if err != nil {
		cobra.CheckErr(err)
	}

	_, err = os.Stat(repoDir)
	if err != nil {
		if !os.IsNotExist(err) {
			cobra.CheckErr(err)
		}
	} else {
		cobra.CheckErr("template " + repoDir + " already exists")
	}

	rootDir, err := repository.RootDir()
	cobra.CheckErr(err)

	gitArgs := []string{"-C", rootDir, "clone", repoUrl}
	fmt.Println("git", gitArgs)

	err = utils.ExecCmd("git", gitArgs...)
	cobra.CheckErr(err)
}
