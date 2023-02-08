package app

import (
	"fmt"

	"github.com/cqroot/ceres/internal/repoconf"
	"github.com/cqroot/ceres/internal/utils"
)

func TestConfig(tomlPath string) error {
	rc, err := repoconf.ParseToml(tomlPath)
	if err != nil {
		return err
	}

	data, err := repoconf.Data(rc)
	if err != nil {
		return err
	}

	fmt.Println(utils.ColorString("Data:"))
	for k, v := range data {
		fmt.Printf("    %s: %+v\n", utils.ColorString(k), v)
	}
	fmt.Println()

	re := repoconf.NewRuleEvaluator(data)
	if len(rc.IncludePathRules) != 0 {
		fmt.Println(utils.ColorString("Include Path Rules:"))
		checkRule(re, rc.IncludePathRules)
		fmt.Println()
	}
	if len(rc.ExcludePathRules) != 0 {
		fmt.Println(utils.ColorString("Exclude Path Rules"))
		checkRule(re, rc.ExcludePathRules)
		fmt.Println()
	}

	return nil
}

func checkRule(re *repoconf.RuleEvaluator, ruleMap map[string][]string) {
	for path, rule := range ruleMap {
		result := re.EvalRules(rule)
		fmt.Printf("    %s: %+v\n", utils.ColorString(path), result)
	}
}
