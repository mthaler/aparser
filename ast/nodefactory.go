package ast

import (
	"aparser"
	"strings"
)

// leafs

func CreateUnaryOperation(b *aparser.Buffer) interface{} {
	s := strings.TrimSpace(b.CurrentMatch())
	return unaryOperation{funcName: s}
}

func CreateBinaryOperation(b *aparser.Buffer) interface{} {
	s := strings.TrimSpace(b.CurrentMatch())
	return binaryOperation{operator: s}
}

func CreateDoubleOperand(b *aparser.Buffer) interface{} {
	s := strings.TrimSpace(b.CurrentMatch())
	o, err := parseDoubleOperand(s)
	if err != nil {
		panic(err)
	}
	return o
}

func CreateBooleanOperand(b *aparser.Buffer) interface{} {
	s := strings.TrimSpace(b.CurrentMatch())
	o, err := parseBoolOperand(s)
	if err != nil {
		panic(err)
	}
	return o
}

func CreateStringOperand(b *aparser.Buffer) interface{} {
	s := strings.TrimSpace(b.CurrentMatch())
	o, err := parseStringOperand(s)
	if err != nil {
		panic(err)
	}
	return o
}

func CreateIdentifier(b *aparser.Buffer) interface{} {
	s := strings.TrimSpace(b.CurrentMatch())
	return identifier{name: s}
}

// branches

func CreateBinaryLeftAssoc(b *aparser.Buffer) interface{} {
	if b.CurrentCodeBlockLength() == 1 {
		return aparser.PT
	}
	code := b.CurrentCodeBlock()
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

func CreateBinaryRightAssoc(b *aparser.Buffer) interface{} {
	if b.CurrentCodeBlockLength() == 1 {
		return aparser.PT
	}
	code := b.CurrentCodeBlock()
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

func CreateUnaryPrefix(b *aparser.Buffer) interface{} {
	if b.CurrentCodeBlockLength() == 1 {
		return aparser.PT
	} else if b.CurrentCodeBlockLength() == 2 {
		return b.CurrentCodeBlock()
	}
	panic("current code block has wrong number of elements")
}

func CreateTernaryConditional(b *aparser.Buffer) interface{} {
	if b.CurrentCodeBlockLength() != 3 {
		return aparser.PT
	} else {
		c := b.CurrentCodeBlock()
		r := make([]interface{}, b.CurrentCodeBlockLength()+1)
		r[0] = ternaryOperation{operator: "?:"}
		copy(r[1:], c)
		return r
	}
}

func CreateFunction(b *aparser.Buffer) interface{} {
	if b.CurrentCodeBlockLength() == 1 {
		return aparser.PT
	} else {
		c := b.CurrentCodeBlock()
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
