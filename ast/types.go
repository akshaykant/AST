package ast

import "AST/tokens"

/*
	Define Node structure
*/
type Node struct {
	Value interface{} //to accept different struct, making it generic
	Left  *Node
	Right *Node
}

type BooleanExpression struct {
	Value bool
}

type IntegerExpression struct {
	Value int
}

type Operator struct {
	Value tokens.Token
}
