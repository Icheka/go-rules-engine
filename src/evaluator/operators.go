package evaluator

import (
	"errors"
	"fmt"
)

func EvaluateOperator(fact, value interface{}, operator string) (bool, error) {
	switch operator {
	case "=":
		fallthrough
	case "eq":
		return fact == value, nil
	default:
		return false, errors.New(fmt.Sprintf("unrecognised operator %s", operator))
	}
}
