package ceres

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GitClone(repoPath string, cloneURL string) error {
	downloadPath := filepath.Join(filepath.Dir(repoPath), fmt.Sprintf(".%s-tmp", filepath.Base(repoPath)))

	if err := os.MkdirAll(filepath.Dir(downloadPath), 0o755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	gitCmd := exec.Command("git", "clone", cloneURL, downloadPath)
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		if cleanupErr := os.RemoveAll(downloadPath); cleanupErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to cleanup %s: %v\n", downloadPath, cleanupErr)
		}
		return fmt.Errorf("failed to clone: %w", err)
	}

	if err := os.Rename(downloadPath, repoPath); err != nil {
		if cleanupErr := os.RemoveAll(downloadPath); cleanupErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to cleanup %s: %v\n", downloadPath, cleanupErr)
		}
		return fmt.Errorf("failed to rename: %w", err)
	}

	return nil
}
