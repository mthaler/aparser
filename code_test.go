package aparser

import "testing"

func TestBuffer_SetCurrentCodeBlock(t *testing.T) {
	c := Code{}
	if len(c.Code) != 0 {
		t.Error("Length of code should be 0")
	}
	c.SetCurrentCodeBlock(42)
	if len(c.Code) != 1 {
		t.Error("Length of code should be 1")
	}
	if c.Code[0] != 42 {
		t.Error("First code element should be 42")
	}
	if c.currentCodeBlockEnd != 1 {
		t.Error("Current code block end position should be 1")
	}
}

func TestBuffer_SetCurrentCodeBlockEndPosition(t *testing.T) {
	c := Code{}
	if len(c.Code) != 0 {
		t.Error("Length of code should be 0")
	}
	c.SetCurrentCodeBlock(42)
	if len(c.Code) != 1 {
		t.Error("Length of code should be 1")
	}
	if c.Code[0] != 42 {
		t.Error("First code element should be 42")
	}
	if c.currentCodeBlockEnd != 1 {
		t.Error("Current code block end position should be 1")
	}
	c.SetCurrentCodeBlockEndPosition(0)
	if len(c.Code) != 0 {
		t.Error("Length of code should be 0")
	}
	if c.currentCodeBlockEnd != 0 {
		t.Error("Current code block end position should be 1")
	}
}
