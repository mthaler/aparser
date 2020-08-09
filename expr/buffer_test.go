package expr

import "testing"

func TestCreateBuffer(t *testing.T) {
	b := CreateBuffer("foo")
	if string(b.text) != "foo" {
		t.Error("text should be ff")
	}
	if b.currentPosition != 0 {
		t.Error("Current position should be 0")
	}
}

func TestBuffer_Rest(t *testing.T) {
	b := CreateBuffer("foo")
	b.SetCurrentPosition(1)
	if string(b.Rest()) != "oo" {
		t.Error("Rest should be oo")
	}
}

func TestBuffer_IncrementCurrentPosition(t *testing.T) {
	b := CreateBuffer("foo")
	if b.currentPosition != 0 {
		t.Error("Current position should be 0")
	}
	b.IncrementCurrentPosition()
	if b.currentPosition != 1 {
		t.Error("Current position should be 1")
	}
}

func TestBuffer_IncrementCurrentPositionBy(t *testing.T) {
	b := CreateBuffer("foo")
	if b.currentPosition != 0 {
		t.Error("Current position should be 0")
	}
	b.IncrementCurrentPositionBy(2)
	if b.currentPosition != 2 {
		t.Error("Current position should be 2")
	}
}

func TestBuffer_CurrentChar(t *testing.T) {
	b := CreateBuffer("foo")
	b.SetCurrentPosition(1)
	if b.CurrentChar() != 'o' {
		t.Error("Current char should be o")
	}
}

func TestBuffer_HasMoreChars(t *testing.T) {
	b := CreateBuffer("foo")
	if !b.HasMoreChars() {
		t.Error("HasMoreChars should be true")
	}
	b.IncrementCurrentPosition()
	if !b.HasMoreChars() {
		t.Error("HasMoreChars should be true")
	}
	b.IncrementCurrentPosition()
	if !b.HasMoreChars() {
		t.Error("HasMoreChars should be true")
	}
	b.IncrementCurrentPosition()
	if b.HasMoreChars() {
		t.Error("HasMoreChars should be false")
	}
}
