package templates

import (
	"fmt"

	"github.com/cqroot/prompt"
)

func ChooseTemplate() (string, error) {
	choices, err := Templates()
	if err != nil {
		return "", err
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
