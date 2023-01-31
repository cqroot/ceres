package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cqroot/ceres/internal/templates"
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
	dataDir, err := templates.DataDir()
	cobra.CheckErr(err)

	repoUrl := args[0]
	if strings.Count(repoUrl, "/") == 1 {
		repoUrl = "https://github.com/" + repoUrl
	}

	templatePath := filepath.Join(dataDir, filepath.Base(repoUrl))
	_, err = os.Stat(templatePath)
	if err != nil {
		if !os.IsNotExist(err) {
			cobra.CheckErr(err)
		}
	} else {
		cobra.CheckErr("template " + templatePath + " already exists")
	}

	gitArgs := []string{"-C", dataDir, "clone", repoUrl}
	fmt.Println("git", gitArgs)

	err = utils.ExecCmd("git", gitArgs...)
	cobra.CheckErr(err)
}
