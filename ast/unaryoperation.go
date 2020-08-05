package ast

import (
	"errors"
	"fmt"
	"math"
)

type unaryOperation struct {
	node     Node
	funcName string
}

func (u unaryOperation) Evaluate() (interface{}, error) {
	v, err := u.node.Evaluate()
	if err != nil {
		return nil, err
	}
	x, err := toDouble(v)
	if err != nil {
		return nil, err
	}
	switch u.funcName {
	case "-":
		return -x, nil
	case "abs":
		return math.Abs(x), nil
	case "acos":
		return math.Acos(x), nil
	case "asin":
		return math.Asin(x), nil
	case "atan":
		return math.Atan(x), nil
	case "cos":
		return math.Cos(x), nil
	case "cosh":
		return math.Cosh(x), nil
	case "exp":
		return math.Exp(x), nil
	case "log":
		return math.Log(x), nil
	case "log10":
		return math.Log10(x), nil
	case "sign":
		if x > 0 {
			return 1, nil
		} else if x < 0 {
			return -1, nil
		} else {
			return 0, nil
		}
	case "sin":
		return math.Sin(x), nil
	case "sinh":
		return math.Sinh(x), nil
	case "sqrt":
		return math.Sqrt(x), nil
	case "tan":
		return math.Tan(x), nil
	case "tanh":
		return math.Tanh(x), nil
	case "round":
		return math.Round(x), nil
	default:
		return nil, errors.New(fmt.Sprintln("Unknown unary function:", u.funcName))
	}
}