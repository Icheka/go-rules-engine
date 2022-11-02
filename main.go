package main

import (
	"fmt"
	"os"

	ruleEngine "github.com/Icheka/go-rule-engine/rule_engine"
)

func a() {
	j, err := os.ReadFile("s.json")
	if err != nil {
		panic("not a valid file")
	}
	// rule := ast.ParseJSON(string(j))

	fact := map[string]interface{}{
		"myVars": "hello worlds",
	}

	// fmt.Println(evaluator.EvaluateRule(rule, fact))

	fmt.Println(ruleEngine.New(&ruleEngine.EvaluatorOptions{
		AllowUndefinedVars: true,
	}).EvaluateStruct(string(j), fact))
}

func main() {
	a()
}
