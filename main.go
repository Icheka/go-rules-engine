package main

import (
	"fmt"
	"os"

	"github.com/Icheka/go-rule-engine/src/ast"
	"github.com/Icheka/go-rule-engine/src/evaluator"
)

func main() {
	j, err := os.ReadFile("s.json")
	if err != nil {
		panic("not a valid file")
	}
	rule := ast.ParseJSON(string(j))

	fact := map[string]interface{}{
		"myVar": "hello world",
	}

	fmt.Println(evaluator.EvaluateRule(rule, fact))
}
