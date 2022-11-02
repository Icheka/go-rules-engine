package evaluator

import (
	"fmt"

	"github.com/Icheka/go-rule-engine/src/ast"
)

type Data map[string]interface{}

func EvaluateConditional(conditional *ast.Conditional, fact interface{}) bool {
	ok, err := EvaluateOperator(fact, conditional.Value, conditional.Operator)
	if err != nil {
		panic(err)
	}
	return ok
}

func GetFactValue(condition *ast.Conditional, data Data) interface{} {
	value := data[condition.Fact]

	if value == nil {
		panic(fmt.Sprintf("value for fact %s not found", condition.Fact))
	}

	return value
}

func EvaluateAllCondition(conditions *[]ast.Conditional, data Data) bool {
	isFalse := false

	for _, condition := range *conditions {
		value := GetFactValue(&condition, data)
		if !EvaluateConditional(&condition, value) {
			isFalse = true
		}

		if isFalse {
			return false
		}
	}

	return true
}

func EvaluateAnyCondition(conditions *[]ast.Conditional, data Data) bool {
	for _, condition := range *conditions {
		value := GetFactValue(&condition, data)
		if EvaluateConditional(&condition, value) {
			return true
		}
	}

	return false
}

func EvaluateCondition(condition *[]ast.Conditional, kind string, data Data) bool {
	switch kind {
	case "all":
		return EvaluateAllCondition(condition, data)
	case "any":
		return EvaluateAnyCondition(condition, data)
	default:
		panic(fmt.Sprintf("condition type %s is invalid", kind))
	}
}

func EvaluateRule(rule *ast.Rule, data Data) bool {
	any, all := false, false

	if len(rule.Condition.Any) == 0 {
		any = true
	} else {
		any = EvaluateCondition(&rule.Condition.Any, "any", data)
	}
	if len(rule.Condition.All) == 0 {
		all = true
	} else {
		all = EvaluateCondition(&rule.Condition.All, "all", data)
	}

	return any && all
}
