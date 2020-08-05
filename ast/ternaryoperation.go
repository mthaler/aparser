package ast

import (
	"errors"
	"fmt"
)

type ternaryOperation struct {
	left     Node
	middle   Node
	right    Node
	operator string
}

func (t ternaryOperation) Evaluate() (interface{}, error) {
	l, err := t.left.Evaluate()
	if err != nil {
		return nil, err
	}
	m, err := t.middle.Evaluate()
	if err != nil {
		return nil, err
	}
	r, err := t.right.Evaluate()
	if err != nil {
		return nil, err
	}

	if isBool(l) {
		bl := l.(bool)
		switch t.operator {
		case "?:":
			if bl {
				return m, nil
			} else {
				return r, nil
			}
		case "if":
			if bl {
				return m, nil
			} else {
				return r, nil
			}
		default:
			return nil, errors.New(fmt.Sprintln("Unknown operator:", t.operator))
		}
	} else {
		return nil, errors.New("Cannot evaluate expr")
	}
}