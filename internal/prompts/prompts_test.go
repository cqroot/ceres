package prompts_test

import (
	"testing"

	"github.com/cqroot/sawmill/internal/prompts"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	p := prompts.New("./testdata/test.toml")

	co, err := p.Parse()
	require.Nil(t, err)

	require.Equal(t, ".", co.CommonItem.Output)
	require.Equal(t,
		[]string{"input_1", "toggle_1", "choose_1", "choose_2"},
		co.CommonItem.Variables,
	)

	require.Equal(t, prompts.VariableItem{
		Message: "Add input_1?",
		Type:    "input",
		Meta:    []string{"input_1"},
	}, co.Variable["input_1"])
	require.Equal(t, prompts.VariableItem{
		Message: "Add toggle_1?",
		Type:    "toggle",
		Meta:    []string{"Yes", "No"},
	}, co.Variable["toggle_1"])
	require.Equal(t, prompts.VariableItem{
		Message: "Add choose_1?",
		Type:    "choose",
		Meta:    []string{"item 1", "item 2", "item 3"},
	}, co.Variable["choose_1"])
	require.Equal(t, prompts.VariableItem{
		Message: "Add choose_2?",
		Type:    "choose",
		Meta:    []string{"item 1", "item 2", "item 3"},
	}, co.Variable["choose_2"])

	require.Equal(t, 1, len(co.IncludePathRules))
	require.Equal(t, prompts.Rule{
		Key: "choose_1", Value: "item 1",
	}, co.IncludePathRules["dir/subdir"])
	require.Equal(t, prompts.Rule{
		Key: "choose_2", Value: "item 2",
	}, co.ExcludePathRules["dir/subdir"])
}
