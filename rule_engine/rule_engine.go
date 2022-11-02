package ruleEngine

import (
	"github.com/Icheka/go-rule-engine/src/ast"
	"github.com/Icheka/go-rule-engine/src/evaluator"
)

type EvaluatorOptions struct {
	AllowUndefinedVars bool
}

var defaultOptions = &EvaluatorOptions{
	AllowUndefinedVars: false,
}

type RuleEngine struct {
	EvaluatorOptions
}

func (re *RuleEngine) EvaluateStruct(jsonText string, fact evaluator.Data) bool {
	return evaluator.EvaluateRule(ast.ParseJSON(jsonText), fact, &evaluator.Options{
		AllowUndefinedVars: re.AllowUndefinedVars,
	})
}

func New(options *EvaluatorOptions) *RuleEngine {
	opts := options
	if opts == nil {
		opts = defaultOptions
	}

	return &RuleEngine{
		EvaluatorOptions: *opts,
	}
}
