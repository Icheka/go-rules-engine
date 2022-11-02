package ruleEngine

import (
	"github.com/Icheka/go-rule-engine/src/ast"
	"github.com/Icheka/go-rule-engine/src/evaluator"
)

type results []bool

type EvaluatorOptions struct {
	AllowUndefinedVars bool
}

var defaultOptions = &EvaluatorOptions{
	AllowUndefinedVars: false,
}

type RuleEngine struct {
	EvaluatorOptions
	Rules   []string
	Results results
}

func (re *RuleEngine) EvaluateStruct(jsonText string, fact evaluator.Data) bool {
	return evaluator.EvaluateRule(ast.ParseJSON(jsonText), fact, &evaluator.Options{
		AllowUndefinedVars: re.AllowUndefinedVars,
	})
}

func (re *RuleEngine) AddRule(rule string) *RuleEngine {
	re.Rules = append(re.Rules, rule)
	return re
}

func (re *RuleEngine) AddRules(rules ...string) *RuleEngine {
	re.Rules = append(re.Rules, rules...)
	return re
}

func (re *RuleEngine) EvaluateRules(data evaluator.Data) results {
	for _, rule := range re.Rules {
		re.Results = append(re.Results, re.EvaluateStruct(rule, data))
	}
	return re.Results
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
