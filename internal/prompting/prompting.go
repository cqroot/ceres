/*
Copyright (C) 2025 Keith Chu <cqroot@outlook.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package prompting

import (
	"github.com/cqroot/prompt"
)

type Type string

const (
	TypeInput  Type = "input"
	TypeChoose Type = "choose"
)

type Prompting struct {
	Name    string `yaml:"name"`
	Type    Type   `yaml:"type"`
	Message string `yaml:"message"`
	Default string `yaml:"default"`
}

func Prompt(promptings []Prompting) (map[string]any, error) {
	vars := make(map[string]any)
	for _, prompting := range promptings {
		switch prompting.Type {
		case TypeInput:
			val, err := prompt.New().Ask(prompting.Message).Input(prompting.Default)
			if err != nil {
				return nil, err
			}
			vars[prompting.Name] = val
		case TypeChoose:
			val, err := prompt.New().Ask(prompting.Message).Input(prompting.Default)
			if err != nil {
				return nil, err
			}
			vars[prompting.Name] = val
		}
	}
	return vars, nil
}
