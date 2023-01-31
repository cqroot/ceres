package repository

import (
	"fmt"

	"github.com/cqroot/prompt"
)

func ChooseRepo() (string, error) {
	repos, err := Repos()
	if err != nil {
		return "", err
	}

	if len(repos) == 0 {
		return "", fmt.Errorf("there are no templates locally")
	}

	repo, err := prompt.New().
		Ask("Choose the template repository you want to use:").
		Choose(repos)
	if err != nil {
		return "", err
	}

	return repo, nil
}
