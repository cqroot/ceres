package templates

import (
	"fmt"
	"os"

	"github.com/cqroot/prompt"
)

func ChooseTemplate() (string, error) {
	dataDir, err := DataDir()
	if err != nil {
		return dataDir, err
	}

	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return "", err
	}

	choices := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			choices = append(choices, entry.Name())
		}
	}

	if len(choices) == 0 {
		return "", fmt.Errorf("there are no templates locally")
	}

	choice, err := prompt.New().
		Ask("Choose the template you want to use:").
		Choose(choices)
	if err != nil {
		return "", err
	}

	return choice, nil
}
