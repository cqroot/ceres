package templater

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/cqroot/ceres/internal/repoconf"
)

type Templater struct {
	tmplDir          string
	outputDir        string
	data             map[string]string
	verbose          bool
	includePathRules map[string][]string
	excludePathRules map[string][]string
	ruleEvaluator    repoconf.RuleEvaluator
}

func New(
	tmplDir string, outputDir string, data map[string]string,
	includePathRules map[string][]string, excludePathRules map[string][]string,
) *Templater {
	return &Templater{
		tmplDir:          tmplDir,
		outputDir:        outputDir,
		data:             data,
		verbose:          true,
		includePathRules: includePathRules,
		excludePathRules: excludePathRules,
		ruleEvaluator:    *repoconf.NewRuleEvaluator(data),
	}
}

func (t *Templater) SetVerbose(verbose bool) {
	t.verbose = verbose
}

func (t Templater) validateOutputDir() error {
	f, err := os.Open(t.outputDir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(t.outputDir, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err != nil {
		if err == io.EOF {
			return nil
		} else {
			return err
		}
	}
	return fmt.Errorf("output directory %s is not empty", t.outputDir)
}

func (t Templater) copyFile(input, output string) error {
	err := makeParentDirs(output)
	if err != nil {
		return err
	}

	sourceFileStat, err := os.Stat(input)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", output)
	}

	source, err := os.Open(input)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(output)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func (t Templater) Execute() error {
	err := t.validateOutputDir()
	if err != nil {
		return err
	}

	if t.verbose {
		defer fmt.Println()
	}

	err = filepath.Walk(t.tmplDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relpath, err := filepath.Rel(t.tmplDir, path)
		if err != nil {
			return err
		}
		if relpath == "." {
			return nil
		}

		excludeRule, ok := t.excludePathRules[relpath]
		if ok && t.ruleEvaluator.EvalRules(excludeRule) {
			if info.IsDir() {
				return filepath.SkipDir
			} else {
				return nil
			}
		}
		includeRule, ok := t.includePathRules[relpath]
		if ok && !t.ruleEvaluator.EvalRules(includeRule) {
			if info.IsDir() {
				return filepath.SkipDir
			} else {
				return nil
			}
		}

		if info.IsDir() {
			return nil
		}

		inputFile := filepath.Join(t.tmplDir, relpath)
		if filepath.Ext(relpath) != ".tmpl" {
			outputFile := filepath.Join(t.outputDir, relpath)
			if t.verbose {
				fmt.Print(outputFile)
			}
			err := t.copyFile(
				inputFile,
				outputFile,
			)
			if err != nil {
				return err
			}
		} else {
			outputFile := filepath.Join(t.outputDir, relpath[:len(relpath)-len(filepath.Ext(relpath))])
			if t.verbose {
				fmt.Print(outputFile)
			}
			err := ExecuteTemplate(
				inputFile,
				outputFile,
				t.data,
			)
			if err != nil {
				return err
			}
		}

		if t.verbose {
			fmt.Println("  … ", text.FgGreen.Sprint("OK!"))
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
