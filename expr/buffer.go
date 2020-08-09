package expr

import (
	"aparser/ast"
)

type buffer struct {
	text            []rune
	currentPosition int
	matchPosition   int
	ast.Code
}

func CreateBuffer(text string) *buffer {
	b := buffer{text: []rune(text), Code: ast.Code{}}
	return &b
}

func (b *buffer) setCurrentPosition(value int) {
	if value >= 0 && value <= len(b.text) {
		b.currentPosition = value
	} else {
		panic("illegal current position")
	}
}

func (b *buffer) incrementCurrentPosition() {
	b.setCurrentPosition(b.currentPosition + 1)
}

func (b *buffer) incrementCurrentPositionBy(value int) {
	b.setCurrentPosition(b.currentPosition + value)
}

func (b *buffer) rest() []rune {
	return b.text[b.currentPosition:]
}

func (b *buffer) currentMatch() string {
	return string(b.text[b.matchPosition:b.currentPosition])
}

func (b *buffer) setMatchPosition(value int) {
	if value >= 0 && value <= len(b.text) {
		b.matchPosition = value
	} else {
		panic("illegal match position")
	}
}

func (b *buffer) currentChar() rune {
	return b.text[b.currentPosition]
}

func (b *buffer) hasMoreChars() bool {
	return b.currentPosition < len(b.text)
}
