package ast

type Node interface {
	Evaluate() (interface{}, error)
}
