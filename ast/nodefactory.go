package ast

import (
	"strings"
)

// leafs

func CreateUnaryOperation(text string, c *Code) interface{} {
	s := strings.TrimSpace(text)
	return unaryOperation{funcName: s}
}

func CreateBinaryOperation(text string, c *Code) interface{} {
	s := strings.TrimSpace(text)
	return binaryOperation{operator: s}
}

func CreateDoubleOperand(text string, c *Code) interface{} {
	s := strings.TrimSpace(text)
	o, err := parseDoubleOperand(s)
	if err != nil {
		panic(err)
	}
	return o
}

func CreateBooleanOperand(text string, c *Code) interface{} {
	s := strings.TrimSpace(text)
	o, err := parseBoolOperand(s)
	if err != nil {
		panic(err)
	}
	return o
}

func CreateStringOperand(text string, c *Code) interface{} {
	s := strings.TrimSpace(text)
	o, err := parseStringOperand(s)
	if err != nil {
		panic(err)
	}
	return o
}

func CreateIdentifier(text string, c *Code) interface{} {
	s := strings.TrimSpace(text)
	return identifier{name: s}
}

// branches

func CreateBinaryLeftAssoc(text string, c *Code) interface{} {
	if c.CurrentCodeBlockLength() == 1 {
		return PT
	}
	code := c.CurrentCodeBlock()
	return createBinaryLeftAssocNode(code)
}

func createBinaryLeftAssocNode(code []interface{}) interface{} {
	if len(code) < 3 {
		return code
	}
	// copy the first three elements in a new array
	binop := make([]interface{}, 3)
	copy(binop, code[:3])
	// convert to prefix by swapping the first two elements
	binop[0], binop[1] = binop[1], binop[0]
	// replace the first three elements with the new array
	result := make([]interface{}, len(code)-2)
	result[0] = binop
	copy(result[1:], code[3:])
	return createBinaryLeftAssocNode(result)
}

func CreateBinaryRightAssoc(text string, c *Code) interface{} {
	if c.CurrentCodeBlockLength() == 1 {
		return PT
	}
	code := c.CurrentCodeBlock()
	return createBinaryRightAssocNode(code)
}

func createBinaryRightAssocNode(code []interface{}) interface{} {
	if len(code) < 3 {
		return code
	}
	// copy the last three elements in a new array
	binop := make([]interface{}, 3)
	copy(binop, code[len(code)-3:])
	// convert to prefix by swapping the first two elements
	binop[0], binop[1] = binop[1], binop[0]
	// replace the first three elements with the new array
	result := make([]interface{}, len(code)-2)
	result[len(code)-3] = binop
	copy(result[0:len(code)-3], code[3:])
	return createBinaryRightAssocNode(result)
}

func CreateUnaryPrefix(text string, c *Code) interface{} {
	if c.CurrentCodeBlockLength() == 1 {
		return PT
	} else if c.CurrentCodeBlockLength() == 2 {
		return c.CurrentCodeBlock()
	}
	panic("current code block has wrong number of elements")
}

func CreateTernaryConditional(text string, c *Code) interface{} {
	if c.CurrentCodeBlockLength() != 3 {
		return PT
	} else {
		code := c.CurrentCodeBlock()
		r := make([]interface{}, c.CurrentCodeBlockLength() + 1)
		r[0] = ternaryOperation{operator: "?:"}
		copy(r[1:], code)
		return r
	}
}

func CreateFunction(text string, c *Code) interface{} {
	if c.CurrentCodeBlockLength() == 1 {
		return PT
	} else {
		c := c.CurrentCodeBlock()
		f, ok := c[0].(identifier)
		if ok {
			switch len(c) - 1 {
			case 1:
				c[0] = unaryOperation{funcName: f.name}
			case 2:
				c[0] = binaryOperation{operator: f.name}
			case 3:
				c[0] = ternaryOperation{operator: f.name}
			default:
				panic("functions with more then three parameters are not supperted")
			}
			return c
		} else {
			panic("the first element must be an identifier")
		}
	}
}
