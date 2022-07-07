package expr

import (
	"testing"
)

func TestCheckExpr(t *testing.T) {
	exprList := [][]string{
		{"(", "a", "+", "b", ")", "*", "c"},
	}
	for _, expr := range exprList {
		got, err := CheckExpr(expr)
		if err != nil {
			t.Errorf("CheckExpr() expr: %v, result: %v, error = %v", expr, got, err)
			return
		}
		t.Logf("CheckExpr() expr: %v, result: %v", expr, got)
	}
}

func TestPrefixToPostFix(t *testing.T) {
	prefixList := [][]string{
		{"(", "a", "+", "b", ")", "*", "c"},
	}

	for _, prefix := range prefixList {
		got, err := PrefixToPostFix(prefix)
		if err != nil {
			t.Errorf("PrefixToPostFix() prefix: %v error: %v", prefix, err)
			return
		}
		t.Logf("PrefixToPostFix() prefix: %v, got = %v", prefix, got)
	}
}
