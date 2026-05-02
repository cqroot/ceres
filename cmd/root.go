package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cqroot/ceres/internal/prompt"
	"github.com/cqroot/ceres/internal/template"
	"github.com/cqroot/ceres/pkg/ceres"
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "ceres USER/REPO_NAME",
		Short: "Ceres - A project scaffolding tool",
		Long:  `Ceres is a CLI tool for scaffolding projects from templates.`,
		Args:  cobra.ExactArgs(1),
		Run:   runCreate,
	}
	rootCmd.AddCommand(newReposCmd(), newCleanCmd(), newRemoveCmd())
	return &rootCmd
}

func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runCreate(cmd *cobra.Command, args []string) {
	parts := strings.Split(args[0], "/")
	if len(parts) != 2 {
		fmt.Fprintf(os.Stderr, "Error: invalid format, expected USER/REPO_NAME\n")
		os.Exit(1)
	}
	user := parts[0]
	repoName := parts[1]

	repoPath := ceres.GetRepoCachePath(user, repoName)
	templatePath := filepath.Join(repoPath, "template")
	ceresYamlPath := filepath.Join(repoPath, "ceres.yaml")

	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		fmt.Printf("Repo not found locally, cloning from GitHub...\n")
		cloneURL := fmt.Sprintf("https://github.com/%s/%s", user, repoName)

		dataDir := ceres.GetDataDir()
		downloadPath := filepath.Join(dataDir, "github.com", user, fmt.Sprintf(".%s-tmp", repoName))
		os.MkdirAll(downloadPath, 0755)

		gitCmd := exec.Command("git", "clone", cloneURL, downloadPath)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		if err := gitCmd.Run(); err != nil {
			os.RemoveAll(downloadPath)
			fmt.Fprintf(os.Stderr, "Error: failed to clone repository: %v\n", err)
			os.Exit(1)
		}

		os.Rename(downloadPath, repoPath)
	}

	if _, err := os.Stat(ceresYamlPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: ceres.yaml not found in repo\n")
		os.Exit(1)
	}

	cfg, err := ceres.LoadConfig(ceresYamlPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to load ceres.yaml: %v\n", err)
		os.Exit(1)
	}

	env := make(template.Env)

	fmt.Println()
	for _, p := range cfg.Promptings {
		if p.Name == "" {
			p.Name = "value"
		}
		answer, err := prompt.Ask(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		env[p.Name] = answer
	}

	fmt.Println("\n----------")
	for k, v := range env {
		fmt.Printf("%s: %s\n", k, v)
	}

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: template directory not found in repo\n")
		os.Exit(1)
	}

	outputDir := env["project_name"]
	if outputDir == "" {
		outputDir = repoName
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	if err := template.RenderDir(templatePath, outputDir, env); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to render template: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nProject created successfully: %s/\n", outputDir)
}
