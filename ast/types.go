package ast

import "AST/tokens"

/*
	Define Node structure
*/
type Node struct {
	value interface{} //to accept different struct, making it generic
	left  *Node
	right *Node
}

type BooleanExpression struct {
	value bool
}

type IntegerExpression struct {
	value int
}

type Operator struct {
	value tokens.Token
}
