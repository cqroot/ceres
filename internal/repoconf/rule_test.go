package repoconf_test

import (
	"testing"

	"github.com/cqroot/ceres/internal/repoconf"
	"github.com/stretchr/testify/require"
)

func TestRuleEval(t *testing.T) {
	rule := repoconf.NewRuleEvaluator(map[string]string{
		"val_1": "val-1",
	})

	testEval := func(t *testing.T, expr string, expected bool) {
		res, _ := rule.Eval(expr)
		require.Equal(t, expected, res)
	}

	testEval(t, "val_1==val-1", true)
	testEval(t, "val_1!=val-1", false)
	testEval(t, "val_1==val-2", false)
	// key does st
	testEval(t, "val_2==val-2", false)
	testEval(t, "val_2!=val-2", false)
	// syntax er
	testEval(t, "==val", false)
	testEval(t, "val==", false)
	testEval(t, "!=val", false)
	testEval(t, "val!=", false)
	testEval(t, "", false)
	testEval(t, "val", false)
}
