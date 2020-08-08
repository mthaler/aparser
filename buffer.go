/*
Package aparser implements a simple library for creating arithmetic expression parsers
*/
package aparser

type Buffer struct {
	Text            []rune
	currentPosition int
	matchPosition   int
	Code
}

func CreateBuffer(text string) *Buffer {
	b := Buffer{Text: []rune(text), Code: Code{}}
	return &b
}

func (b *Buffer) CurrentPosition() int {
	return b.currentPosition
}

func (b *Buffer) SetCurrentPosition(value int) {
	if value >= 0 && value <= len(b.Text) {
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
	return b.Text[b.currentPosition:]
}

func (b *Buffer) MatchPosition() int {
	return b.matchPosition
}

func (b *Buffer) CurrentMatch() string {
	return string(b.Text[b.matchPosition:b.currentPosition])
}

func (b *Buffer) SetMatchPosition(value int) {
	if value >= 0 && value <= len(b.Text) {
		b.matchPosition = value
	} else {
		panic("illegal match position")
	}
}

func (b *Buffer) CurrentChar() rune {
	return b.Text[b.currentPosition]
}

func (b *Buffer) HasMoreChars() bool {
	return b.currentPosition < len(b.Text)
}
