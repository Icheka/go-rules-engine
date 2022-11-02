package evaluator

import (
	"testing"

	"github.com/Icheka/go-rule-engine/src/ast"
)

func TestEvaluateConditional(t *testing.T) {
	tests := []struct {
		conditional *ast.Conditional
		fact        interface{}
		expected    bool
	}{
		{&ast.Conditional{
			Fact:     "name",
			Operator: "eq",
			Value:    "Icheka",
		},
			"Icheka",
			true,
		},
		{&ast.Conditional{
			Fact:     "name",
			Operator: "eq",
			Value:    "Icheka",
		},
			"Ronie",
			false,
		},
	}

	for i, tt := range tests {
		if ok := EvaluateConditional(tt.conditional, tt.fact); ok != tt.expected {
			t.Errorf("tests[%d] - expected EvaluateConditional to return %t, got=%t", i, tt.expected, ok)
		}
	}
}

func TestEvaluateAllCondition(t *testing.T) {
	tests := []struct {
		payload struct {
			conditions []ast.Conditional
			fact       Data
		}
		expected bool
	}{
		{
			payload: struct {
				conditions []ast.Conditional
				fact       Data
			}{
				conditions: []ast.Conditional{
					{
						Fact:     "planet",
						Operator: "eq",
						Value:    "Neptune",
					},
					{
						Fact:     "colour",
						Operator: "eq",
						Value:    "black",
					},
				},
				fact: Data{
					"planet": "Neptune",
					"colour": "black",
				},
			},
			expected: true,
		},
		{
			payload: struct {
				conditions []ast.Conditional
				fact       Data
			}{
				conditions: []ast.Conditional{
					{
						Fact:     "planet",
						Operator: "eq",
						Value:    "Saturn",
					},
					{
						Fact:     "colour",
						Operator: "eq",
						Value:    "black",
					},
				},
				fact: Data{
					"planet": "Neptune",
					"colour": "black",
				},
			},
			expected: false,
		},
	}

	for i, tt := range tests {
		if ok := EvaluateAllCondition(&tt.payload.conditions, tt.payload.fact); ok != tt.expected {
			t.Errorf("tests[%d] - expected EvaluateAllCondition to be %t, got=%t", i, tt.expected, ok)
		}
	}
}

func TestEvaluateAnyCondition(t *testing.T) {
	tests := []struct {
		payload struct {
			conditions []ast.Conditional
			fact       Data
		}
		expected bool
	}{
		{
			payload: struct {
				conditions []ast.Conditional
				fact       Data
			}{
				conditions: []ast.Conditional{
					{
						Fact:     "planet",
						Operator: "eq",
						Value:    "Neptune",
					},
					{
						Fact:     "colour",
						Operator: "eq",
						Value:    "black",
					},
				},
				fact: Data{
					"planet": "Neptune",
					"colour": "black",
				},
			},
			expected: true,
		},
		{
			payload: struct {
				conditions []ast.Conditional
				fact       Data
			}{
				conditions: []ast.Conditional{
					{
						Fact:     "planet",
						Operator: "eq",
						Value:    "Saturn",
					},
					{
						Fact:     "colour",
						Operator: "eq",
						Value:    "black",
					},
				},
				fact: Data{
					"planet": "Neptune",
					"colour": "black",
				},
			},
			expected: true,
		},
		{
			payload: struct {
				conditions []ast.Conditional
				fact       Data
			}{
				conditions: []ast.Conditional{
					{
						Fact:     "planet",
						Operator: "eq",
						Value:    "Saturn",
					},
					{
						Fact:     "colour",
						Operator: "eq",
						Value:    "white",
					},
				},
				fact: Data{
					"planet": "Neptune",
					"colour": "black",
				},
			},
			expected: false,
		},
	}

	for i, tt := range tests {
		if ok := EvaluateAnyCondition(&tt.payload.conditions, tt.payload.fact); ok != tt.expected {
			t.Errorf("tests[%d] - expected EvaluateAnyCondition to be %t, got=%t", i, tt.expected, ok)
		}
	}
}
