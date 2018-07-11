package assertion

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const (
	EqualOp       = "equal"
	NotEqualOp    = "not_equal"
	ContainsOp    = "contain"
	NotContainsOp = "not_contain"
)

type Operator func(val interface{}, expected interface{}) (bool, error)

var opMap = map[string]Operator{
	EqualOp:       equal,
	NotEqualOp:    notEqual,
	ContainsOp:    contain,
	NotContainsOp: notContain,
}

func AvailableOps() []string {
	var ops []string
	for op := range opMap {
		ops = append(ops, op)
	}
	return ops
}

func NewOperator(op string) (Operator, error) {
	opFunc, ok := opMap[op]
	if !ok {
		return nil, errors.New("operator not supported")
	}

	return opFunc, nil
}

func equal(val interface{}, expected interface{}) (bool, error) {
	if !reflect.DeepEqual(val, expected) {
		return false, errors.New(fmt.Sprintf("expected equal %s, given %s", expected, val))
	}
	return true, nil
}

func notEqual(val interface{}, expected interface{}) (bool, error) {
	if reflect.DeepEqual(val, expected) {
		return false, errors.New(fmt.Sprintf("expected not equal %s, given %s", expected, val))
	}
	return true, nil
}

func contain(val interface{}, expected interface{}) (bool, error) {
	valStr := val.(string)
	expectedStr := expected.(string)
	if !strings.Contains(valStr, expectedStr) {
		return false, errors.New(fmt.Sprintf("expected contain %s, given %s", expected, val))
	}
	return true, nil
}

func notContain(val interface{}, expected interface{}) (bool, error) {
	valStr := val.(string)
	expectedStr := expected.(string)
	if strings.Contains(valStr, expectedStr) {
		return false, errors.New(fmt.Sprintf("expected not contain %s, given %s", expected, val))
	}
	return true, nil
}
