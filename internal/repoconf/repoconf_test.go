package repoconf_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/ceres/internal/repoconf"
)

func TestNewFromToml(t *testing.T) {
	rc, err := repoconf.NewFromToml("./testdata/test.toml")
	require.Nil(t, err)

	// common
	require.Equal(t, ".", rc.Common.Output)
	require.Equal(t,
		[]string{"input_1", "toggle_1", "choose_1", "choose_2"},
		rc.Common.Variables,
	)

	// variables
	require.Equal(t, repoconf.VariableItem{
		Message: "Add input_1?",
		Type:    "input",
		Meta:    []string{"input_1"},
	}, rc.Variable["input_1"])
	require.Equal(t, repoconf.VariableItem{
		Message: "Add toggle_1?",
		Type:    "toggle",
		Meta:    []string{"Yes", "No"},
	}, rc.Variable["toggle_1"])
	require.Equal(t, repoconf.VariableItem{
		Message: "Add choose_1?",
		Type:    "choose",
		Meta:    []string{"item 1", "item 2", "item 3"},
	}, rc.Variable["choose_1"])
	require.Equal(t, repoconf.VariableItem{
		Message: "Add choose_2?",
		Type:    "choose",
		Meta:    []string{"item 1", "item 2", "item 3"},
	}, rc.Variable["choose_2"])

	// include_path_ruls
	require.Equal(t, 1, len(rc.IncludePathRules))
	require.Equal(t, []string{"choose_1==item 1"}, rc.IncludePathRules["dir/subdir"])
	require.Equal(t, []string{"choose_2==item 2", "choose_1!=item 1"}, rc.ExcludePathRules["dir/subdir"])

	// scripts
	require.Equal(t, 1, len(rc.Scripts.AfterScripts))
}
