package ast

import "testing"

func TestToDouble(t *testing.T) {
	f, err := toDouble(3.14)
	if err != nil {
		t.Error("Could not convert 3.14 to double")
	}
	if f != 3.14 {
		t.Error("Converting 3.14 to double should be 1.0")
	}
}

func TestIsBool(t *testing.T) {
	if isBool(42) != false {
		t.Error("isBool(42) should be false")
	}
	if isBool(false) != true {
		t.Error("isBool(false) should be true")
	}
}

func TestIsString(t *testing.T) {
	if isString(42) != false {
		t.Error("isString(42) should be false")
	}
	if isString("foo") != true {
		t.Error("isString(foo) should be true")
	}
}
