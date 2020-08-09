package expr

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
}
