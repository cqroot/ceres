package ceres

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GitClone(repoPath string, cloneURL string) error {
	downloadPath := filepath.Join(filepath.Dir(repoPath), fmt.Sprintf(".%s-tmp", filepath.Base(repoPath)))

	os.MkdirAll(filepath.Dir(downloadPath), 0755)

	gitCmd := exec.Command("git", "clone", cloneURL, downloadPath)
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		os.RemoveAll(downloadPath)
		return fmt.Errorf("failed to clone: %w", err)
	}

	if err := os.Rename(downloadPath, repoPath); err != nil {
		os.RemoveAll(downloadPath)
		return fmt.Errorf("failed to rename: %w", err)
	}

	return nil
}