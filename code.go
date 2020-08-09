package aparser

type Code struct {
	currentCodeBlockStart int
	currentCodeBlockEnd   int
	Code                  []interface{}
}

func CreateCode() *Code {
	c := Code{currentCodeBlockStart: 0, currentCodeBlockEnd: 0, Code: []interface{}{}}
	return &c
}

func (c *Code) CurrentCodeBlockStart() int {
	return c.currentCodeBlockStart
}

func (c *Code) SetCurrentCodeBlockStartPosition(value int) {
	c.currentCodeBlockStart = value
}

func (c *Code) CurrentCodeBlockEndPosition() int {
	return c.currentCodeBlockEnd
}

func (c *Code) SetCurrentCodeBlockEndPosition(value int) {
	c.Code = c.Code[:value]
	c.currentCodeBlockEnd = value
}

func (c *Code) CurrentCodeBlockLength() int {
	return c.currentCodeBlockEnd - c.currentCodeBlockStart
}

func (c *Code) CurrentCodeBlock() []interface{} {
	l := c.CurrentCodeBlockLength()
	if l == 0 {
		return nil
	} else {
		result := make([]interface{}, l)
		for i := 0; i < l; i++ {
			result[i] = c.Code[c.currentCodeBlockStart+i]
		}
		return result
	}
}

func (c *Code) SetCurrentCodeBlock(o interface{}) {
	if c.CurrentCodeBlockLength() == 0 {
		c.Code = append(c.Code, o)
		c.currentCodeBlockEnd++
	} else {
		c.Code = deleteRange(c.Code, c.currentCodeBlockStart, c.CurrentCodeBlockLength())
		c.Code = append(c.Code, o)
		c.currentCodeBlockEnd = c.currentCodeBlockStart + 1
	}
}
