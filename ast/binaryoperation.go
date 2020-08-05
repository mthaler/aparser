package ast

import (
	"errors"
	"fmt"
	"math"
)

type binaryOperation struct {
	left     Node
	right    Node
	operator string
}

func (b binaryOperation) Evaluate() (interface{}, error) {
	l, err := b.left.Evaluate()
	if err != nil {
		return nil, err
	}
	r, err := b.right.Evaluate()
	if err != nil {
		return nil, err
	}

	x, errLeft := toDouble(l)
	y, errRight := toDouble(r)
	if errLeft == nil && errRight == nil {
		return evaluateNumeric(x, y, b.operator)
	} else if isBool(l) && isBool(r) {
		return evaluateBoolean(l.(bool), r.(bool), b.operator)
	} else if isString(l) && isString(r) {
		return evaluateText(l.(string), r.(string), b.operator)
	} else {
		return nil, errors.New(fmt.Sprintf("Unknown nodes: %v %v", b.left, b.right))
	}
}

func evaluateNumeric(x, y float64, op string) (interface{}, error) {
	switch op {
	case "==":
		return x == y, nil
	case "!=":
		return x != y, nil
	case ">":
		return x > y, nil
	case ">=":
		return x >= y, nil
	case "<":
		return x < y, nil
	case "<=":
		return x <= y, nil
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		return x / y, nil
	case "^":
		return math.Pow(x, y), nil
	case "atan2":
		return math.Atan2(x, y), nil
	case "max":
		return math.Max(x, y), nil
	case "min":
		return math.Min(x, y), nil
	default:
		return nil, errors.New(fmt.Sprintln("Unknown operator:", op))
	}
}

func evaluateBoolean(x, y bool, op string) (interface{}, error) {
	switch op {
	case "==":
		return x == y, nil
	case "!=":
		return x != y, nil
	case "&&":
		return x && y, nil
	case "||":
		return x || y, nil
	default:
		return nil, errors.New(fmt.Sprintln("Unknown operator:", op))
	}
}

func evaluateText(x, y, op string) (interface{}, error) {
	switch op {
	case "==":
		return x == y, nil
	case "!=":
		return x != y, nil
	case "+":
		return x + y, nil
	default:
		return nil, errors.New(fmt.Sprintln("Unknown operator:", op))
	}
}
