package ast

import "strconv"

type operandNode struct {
	Node
}

func createOperandNode(value Node) Node {
	return operandNode{value}
}

type boolOperand struct {
	Value bool
}

func newBoolOperand(value bool) operandNode {
	return operandNode{boolOperand{Value: value}}
}

func parseBoolOperand(s string) (operandNode, error) {
	b, err := strconv.ParseBool(s)
	return newBoolOperand(b), err
}

func (b boolOperand) Evaluate() (interface{}, error) {
	return b.Value, nil
}

type doubleOperand struct {
	Value float64
}

func newDoubleOperand(value float64) operandNode {
	return operandNode{doubleOperand{Value: value}}
}

func parseDoubleOperand(s string) (operandNode, error) {
	f, err := strconv.ParseFloat(s, 64)
	return newDoubleOperand(f), err
}

func (d doubleOperand) Evaluate() (interface{}, error) {
	return d.Value, nil
}

type stringOperand struct {
	Value string
}

func newStringOperand(s string) operandNode {
	return operandNode{stringOperand{Value: s}}
}

func parseStringOperand(s string) (operandNode, error) {
	s = s[1 : len(s)-1]
	return newStringOperand(s), nil
}

func (s stringOperand) Evaluate() (interface{}, error) {
	return s.Value, nil
}
