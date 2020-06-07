package checker

import "AST/tokens"

type SyntaxType string

const (
	Int     SyntaxType = "IntegerExpression"
	Bool    SyntaxType = "BooleanExpression"
	Op      SyntaxType = "Operator"
	Invalid SyntaxType = "InvalidExpression"
)

func CheckTokenType(token tokens.Token) (SyntaxType, string) {
	if token == tokens.Divide || token == tokens.Multiply || token == tokens.Plus || token == tokens.Minus {
		return Int, ""
	}
	if token == tokens.EqualTo || token == tokens.NotEqualTo || token == tokens.LogicalAnd || token == tokens.LogicalOr {
		return Bool, ""
	}
	return Invalid, "Error : Invalid Operator"
}
