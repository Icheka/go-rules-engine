package evaluator

import "testing"

func TestEvaluateOperator(t *testing.T) {
	tests := []struct {
		fact     interface{}
		value    interface{}
		operator string
		expected bool
	}{
		{"hi", "hi", "eq", true},
		{"hi", "hi", "=", true},
		{"hi", "his", "=", false},
		{"hi", 4, "=", false},
		{4, 4, "=", true},
	}

	for i, tt := range tests {
		ok, err := EvaluateOperator(tt.fact, tt.value, tt.operator)
		if err != nil {
			t.Errorf("tests[%d] - unexpected error (%s)", i, err)
		}
		if ok != tt.expected {
			t.Errorf("tests[%d] - expected EvaluateOperator to be %t, got=%t", i, tt.expected, ok)
		}
	}
}
