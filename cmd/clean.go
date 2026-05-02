package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cqroot/ceres/pkg/ceres"
	"github.com/spf13/cobra"
)

func newCleanCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "Remove incomplete downloaded repos",
		Run:   runClean,
	}
}

func runClean(cmd *cobra.Command, args []string) {
	dataDir := ceres.GetDataDir()
	githubDir := filepath.Join(dataDir, "github.com")

	entries, err := os.ReadDir(githubDir)
	if err != nil {
		fmt.Printf("No incomplete repos found.\n")
		return
	}

	var cleaned int
	for _, userEntry := range entries {
		if !userEntry.IsDir() {
			continue
		}
		userPath := filepath.Join(githubDir, userEntry.Name())
		userEntries, err := os.ReadDir(userPath)
		if err != nil {
			continue
		}
		for _, repoEntry := range userEntries {
			name := repoEntry.Name()
			if len(name) > 0 && name[0] == '.' {
				repoPath := filepath.Join(userPath, name)
				if err := os.RemoveAll(repoPath); err == nil {
					fmt.Printf("Removed: %s/%s\n", userEntry.Name(), name)
					cleaned++
				}
			}
		}
	}

	if cleaned == 0 {
		fmt.Printf("No incomplete repos found.\n")
	} else {
		fmt.Printf("\nCleaned %d incomplete repos.\n", cleaned)
	}
}