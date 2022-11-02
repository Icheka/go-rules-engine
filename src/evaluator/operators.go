package evaluator

import (
	"fmt"
)

func EvaluateOperator(fact, value interface{}, operator string) (bool, error) {
	switch operator {
	case "=":
		fallthrough
	case "eq":
		return fact == value, nil
	default:
		return false, fmt.Errorf("unrecognised operator %s", operator)
	}
}
