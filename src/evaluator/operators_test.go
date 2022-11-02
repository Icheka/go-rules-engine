package evaluator

import "testing"

func TestEvaluateOperator(t *testing.T) {
	tests := []struct {
		identifier interface{}
		value      interface{}
		operator   string
		expected   bool
	}{
		{"hi", "hi", "eq", true},
		{"hi", "hi", "=", true},
		{"hi", "his", "=", false},
		{"hi", 4, "=", false},
		{4, 4, "=", true},

		{4, 4, "!=", false},
		{4, 5, "neq", true},

		{4, 5, "<", true},
		{6, 5, "lt", false},

		{4, 5, ">", false},
		{6, 5, "gt", true},

		{4, 5, ">=", false},
		{6, 5, "gte", true},
		{5, 5, "gte", true},

		{4, 5, "<=", true},
		{6, 5, "lte", false},
		{5, 5, "lte", true},
	}

	for i, tt := range tests {
		ok, err := EvaluateOperator(tt.identifier, tt.value, tt.operator)
		if err != nil {
			t.Errorf("tests[%d] - unexpected error (%s)", i, err)
		}
		if ok != tt.expected {
			t.Errorf("tests[%d] - expected EvaluateOperator to be %t, got=%t", i, tt.expected, ok)
		}
	}
}
