package ruleEngine

import (
	"github.com/Icheka/go-rules-engine/src/ast"
	"github.com/Icheka/go-rules-engine/src/evaluator"
)

type results []ast.Event

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

func (re *RuleEngine) EvaluateStruct(jsonText *ast.Rule, identifier evaluator.Data) bool {
	return evaluator.EvaluateRule(jsonText, identifier, &evaluator.Options{
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
	for _, j := range re.Rules {
		rule := ast.ParseJSON(j)

		if re.EvaluateStruct(rule, data) {
			re.Results = append(re.Results, rule.Event)
		}
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
