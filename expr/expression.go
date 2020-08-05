package expr

import (
	"aparser"
)

type Expression interface {
	Id() string
	SetId(id string)
	CreateNode() func(buffer *aparser.Buffer) interface{}
	SetCreateNode(m func(buffer *aparser.Buffer) interface{})
	parse(buffer *aparser.Buffer) bool
}

func Parse(e Expression, buffer *aparser.Buffer) bool {
	currentPosition := buffer.CurrentPosition()
	matchPosition := buffer.MatchPosition()
	buffer.SetMatchPosition(buffer.CurrentPosition())
	codeBlockStart := buffer.CurrentCodeBlockStart()
	codeBlockEnd := buffer.CurrentCodeBlockEndPosition()
	buffer.SetCurrentCodeBlockStartPosition(codeBlockEnd)

	match := e.parse(buffer)
	if match {
		result := onMatch(e, buffer)
		if result != nil && result != aparser.PT {
			buffer.SetCurrentCodeBlock(result)
		}
	} else {
		buffer.SetCurrentPosition(currentPosition)
		buffer.SetMatchPosition(matchPosition)
		buffer.SetCurrentCodeBlockEndPosition(codeBlockEnd)
	}
	buffer.SetMatchPosition(matchPosition)
	buffer. SetCurrentCodeBlockStartPosition(codeBlockStart)
	return match
}

func onMatch(e Expression, buffer *aparser.Buffer) interface{} {
	m := e.CreateNode()
	if m != nil {
		return m(buffer)
	} else {
		return aparser.PT
	}
}
