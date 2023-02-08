package repoconf

import (
	"bytes"
	"fmt"
)

type RuleEvaluator struct {
	data map[string]string
}

func NewRuleEvaluator(data map[string]string) *RuleEvaluator {
	return &RuleEvaluator{
		data: data,
	}
}

func (r RuleEvaluator) Eval(expr string) (bool, error) {
	exprb := []byte(expr)
	neq := false
	pos := bytes.Index(exprb, []byte("=="))
	if pos == -1 {
		pos = bytes.Index(exprb, []byte("!="))
		if pos != -1 {
			neq = true
		}
	}
	if pos != -1 {
		key := expr[0:pos]
		val := expr[pos+2:]
		actual, ok := r.data[string(key)]
		if !ok {
			return false, fmt.Errorf("variable \"%s\" does not exist", key)
		}

		if neq {
			return actual != string(val), nil
		} else {
			return actual == string(val), nil
		}
	}

	return false, fmt.Errorf("syntax error: %s", expr)
}

func (r RuleEvaluator) EvalRules(exprs []string) bool {
	for _, expr := range exprs {
		result, _ := r.Eval(expr)
		if !result {
			return false
		}
	}
	return true
}
