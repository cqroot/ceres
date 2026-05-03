package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cqroot/ceres/internal/logger"
	"github.com/cqroot/ceres/internal/prompt"
	"github.com/cqroot/ceres/internal/template"
	"github.com/cqroot/ceres/pkg/ceres"
	"github.com/cqroot/ceres/pkg/version"
	"github.com/spf13/cobra"
)

var dataDir string

func newRootCmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "ceres USER/REPO_NAME",
		Short: "Ceres - A project scaffolding tool",
		Long:  `Ceres is a CLI tool for scaffolding projects from templates.`,
		Args:  cobra.ExactArgs(1),
		Run:   runCreate,
	}
	rootCmd.PersistentFlags().StringVarP(&dataDir, "data-dir", "d", "", "Override data directory")
	rootCmd.AddCommand(newAddCmd(), newCleanCmd(), newRemoveCmd(), newReposCmd())
	cobra.OnInitialize(func() {
		if dataDir != "" {
			ceres.SetDataDir(dataDir)
		}
	})
	rootCmd.Version = version.Get().String()
	return &rootCmd
}

func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func runCreate(cmd *cobra.Command, args []string) {
	parts := strings.Split(args[0], "/")
	if len(parts) != 2 {
		logger.Errorf("invalid format, expected USER/REPO_NAME")
		os.Exit(1)
	}
	user := parts[0]
	repoName := parts[1]

	repoPath := ceres.GetRepoCachePath(user, repoName)
	templatePath := filepath.Join(repoPath, "template")
	ceresYamlPath := filepath.Join(repoPath, "ceres.yaml")

	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		logger.Info("Repo not found locally, cloning from GitHub...")
		cloneURL := fmt.Sprintf("https://github.com/%s/%s", user, repoName)

		if err := ceres.GitClone(repoPath, cloneURL); err != nil {
			logger.Errorf("%v", err)
			os.Exit(1)
		}
	}

	if _, err := os.Stat(ceresYamlPath); os.IsNotExist(err) {
		logger.Error("ceres.yaml not found in repo")
		os.Exit(1)
	}

	cfg, err := ceres.LoadConfig(ceresYamlPath)
	if err != nil {
		logger.Errorf("failed to load ceres.yaml: %v", err)
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
			logger.Errorf("%v", err)
			os.Exit(1)
		}
		env[p.Name] = answer
	}

	env["repo"] = args[0]

	fmt.Println("\n----------")
	for k, v := range env {
		fmt.Printf("%s: %s\n", k, v)
	}

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		logger.Error("template directory not found in repo")
		os.Exit(1)
	}

	outputDir := env["project_name"]
	if outputDir == "" {
		outputDir = repoName
	}

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		logger.Errorf("failed to create output directory: %v", err)
		os.Exit(1)
	}

	if err := template.RenderDir(templatePath, outputDir, env); err != nil {
		logger.Errorf("failed to render template: %v", err)
		os.Exit(1)
	}

	if err := ceres.SaveEnv(outputDir, env); err != nil {
		logger.Errorf("failed to save env file: %v", err)
		os.Exit(1)
	}

	for _, cmdStr := range cfg.Before {
		renderedCmd, _ := template.RenderString(cmdStr, env)
		if err := runCommand(renderedCmd, outputDir); err != nil {
			logger.Errorf("before command failed: %v", err)
			os.Exit(1)
		}
	}

	for _, cmdStr := range cfg.After {
		renderedCmd, _ := template.RenderString(cmdStr, env)
		if err := runCommand(renderedCmd, outputDir); err != nil {
			logger.Errorf("after command failed: %v", err)
			os.Exit(1)
		}
	}

	logger.Infof("Project created successfully: %s/", outputDir)
}

func runCommand(cmdStr string, workDir string) error {
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Dir = workDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
