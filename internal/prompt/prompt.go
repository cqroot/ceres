package prompt

import (
	"fmt"
	"strings"

	"github.com/cqroot/ceres/pkg/ceres"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
	"github.com/cqroot/prompt/input"
)

func Ask(promptCfg ceres.PromptConfig) (string, error) {
	p := prompt.New().Ask(promptCfg.Message)

	switch promptCfg.Type {
	case "input", "string":
		return p.Input(promptCfg.Default)
	case "multiline":
		return p.Input(promptCfg.Default)
	case "int":
		return p.Input(promptCfg.Default, input.WithValidateFunc(func(s string) error {
			var n int
			_, err := fmt.Sscanf(s, "%d", &n)
			return err
		}))
	case "bool":
		return p.Choose([]string{"yes", "no"})
	case "choose":
		if len(promptCfg.Options) == 0 {
			return p.Input(promptCfg.Default)
		}
		opts := []choose.Option{}
		if promptCfg.Default != "" {
			for i, opt := range promptCfg.Options {
				if opt == promptCfg.Default {
					opts = append(opts, choose.WithDefaultIndex(i))
					break
				}
			}
		}
		return p.Choose(promptCfg.Options, opts...)
	case "multichoose":
		if len(promptCfg.Options) == 0 {
			return p.Input(promptCfg.Default)
		}
		result, err := p.MultiChoose(promptCfg.Options)
		if err != nil {
			return "", err
		}
		return strings.Join(result, ", "), nil
	default:
		return p.Input(promptCfg.Default)
	}
}
