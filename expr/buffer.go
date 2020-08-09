package expr

import (
	"aparser/ast"
)

type Buffer struct {
	text            []rune
	currentPosition int
	matchPosition   int
	ast.Code
}

func CreateBuffer(text string) *Buffer {
	b := Buffer{text: []rune(text), Code: ast.Code{}}
	return &b
}

func (b *Buffer) CurrentPosition() int {
	return b.currentPosition
}

func (b *Buffer) SetCurrentPosition(value int) {
	if value >= 0 && value <= len(b.text) {
		b.currentPosition = value
	} else {
		panic("illegal current position")
	}
}

func (b *Buffer) IncrementCurrentPosition() {
	b.SetCurrentPosition(b.currentPosition + 1)
}

func (b *Buffer) IncrementCurrentPositionBy(value int) {
	b.SetCurrentPosition(b.currentPosition + value)
}

func (b *Buffer) Rest() []rune {
	return b.text[b.currentPosition:]
}

func (b *Buffer) MatchPosition() int {
	return b.matchPosition
}

func (b *Buffer) CurrentMatch() string {
	return string(b.text[b.matchPosition:b.currentPosition])
}

func (b *Buffer) SetMatchPosition(value int) {
	if value >= 0 && value <= len(b.text) {
		b.matchPosition = value
	} else {
		panic("illegal match position")
	}
}

func (b *Buffer) CurrentChar() rune {
	return b.text[b.currentPosition]
}

func (b *Buffer) HasMoreChars() bool {
	return b.currentPosition < len(b.text)
}
