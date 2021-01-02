package aparser

import (
	"testing"
)

func TestEval(t *testing.T) {
	r, err := Eval("3 + 4")
	if err != nil {
		t.Error(err)
	}
	if r != 7.0 {
		t.Error("Result should be 7.0")
	}
	r, err = Eval("3 + 4 + 5")
	if err != nil {
		t.Error(err)
	}
	if r != 12.0 {
		t.Error("Result should be 12.0")
	}
	r, err = Eval("2 * 3 + 4")
	if err != nil {
		t.Error(err)
	}
	if r != 10.0 {
		t.Error("Result should be 10.0")
	}
	r, err = Eval("2 + 3 * 4")
	if err != nil {
		t.Error(err)
	}
	if r != 14.0 {
		t.Error("Result should be 14.0")
	}
}
