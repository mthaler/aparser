package ast

func CreateAST(code []interface{}) Node {
	return createNode(code)
}

func createNode(code []interface{}) Node {
	if len(code) == 0 {
		return nil
	} else if len(code) == 1 {
		c, ok := code[0].([]interface{})
		if ok {
			return createNode(c)
		} else {
			return code[0].(Node)
		}
	} else if len(code) == 2 {
		return createUnaryNode(code)
	} else if len(code) >= 3 {
		if len(code) == 4 {
			_, ok := code[0].(ternaryOperation)
			if ok {
				return createTernaryNode(code)
			} else {
				return createBinaryNode(code)
			}
		} else {
			return createBinaryNode(code)
		}
	}
	panic("Could not build abstract syntax tree!")
}

func createUnaryNode(code []interface{}) Node {
	u, ok := code[0].(unaryOperation)
	if ok {
		o, ok := code[1].(operandNode)
		if ok {
			u.node = o
		}
		subtree, ok := code[1].([]interface{})
		if ok {
			u.node = createNode(subtree)
		}
		return u
	} else {
		panic("first element must be an unary operation")
	}
}

func createBinaryNode(code []interface{}) Node {
	b, ok := code[0].(binaryOperation)
	if ok {
		l, ok := code[1].(operandNode)
		if ok {
			b.left = l
		} else {
			b.left = createNode(code[1].([]interface{}))
		}
		r, ok := code[2].(operandNode)
		if ok {
			b.right = r
		} else {
			b.right = createNode(code[2].([]interface{}))
		}
		return b
	} else {
		panic("first element not a binary operation")
	}
}

func createTernaryNode(code []interface{}) Node {
	t, ok := code[0].(ternaryOperation)
	if ok {
		l, ok := code[1].(Node)
		if ok {
			t.left = l
		} else {
			t.left = createNode(code[1].([]interface{}))
		}
		m, ok := code[2].(Node)
		if ok {
			t.middle = m
		} else {
			t.middle = createNode(code[2].([]interface{}))
		}
		r, ok := code[3].(Node)
		if ok {
			t.right = r
		} else {
			t.right = createNode(code[3].([]interface{}))
		}
		return t
	} else {
		panic("first element not a ternary operation")
	}
}
