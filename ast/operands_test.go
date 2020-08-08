package ast

import "testing"

func TestParseBoolOperand(t *testing.T) {
	o, err := parseBoolOperand("false")
	if err != nil {
		t.Error(err)
	}
	checkEvaluate(t, o, false)
	o, err = parseBoolOperand("true")
	if err != nil {
		t.Error(err)
	}
	checkEvaluate(t, o, true)
}

func TestParseDoubleOperand(t *testing.T) {
	o, err := parseDoubleOperand("3.14")
	if err != nil {
		t.Error(err)
	}
	checkEvaluate(t, o, 3.14)
}

func TestParseStringOperand(t *testing.T) {
	o, err := parseStringOperand("\"foo\"")
	if err != nil {
		t.Error(err)
	}
	checkEvaluate(t, o, "foo")
}
