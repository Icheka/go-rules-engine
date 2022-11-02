package main

import (
	"fmt"
	"os"

	ruleEngine "github.com/Icheka/go-rule-engine/rule_engine"
)

func a() {
	a, _ := os.ReadFile("s.json")
	b, _ := os.ReadFile("d.json")

	// rule := ast.ParseJSON(string(j))

	fact := map[string]interface{}{
		"myVar": "hello world",
		"name":  "icheka",
	}

	// fmt.Println(evaluator.EvaluateRule(rule, fact))

	fmt.Println(ruleEngine.New(&ruleEngine.EvaluatorOptions{
		AllowUndefinedVars: true,
	}).AddRules(string(a), string(b)).EvaluateRules(fact))
}

func main() {
	a()
}
