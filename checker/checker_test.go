package checker

import (
	"AST/ast"
	"AST/tokens"
	"testing"
)

type TestCase struct {
	Node          *ast.Node
	ExpectedType  SyntaxType
	ExpectedError string
}

func TestChecker(t *testing.T) {

	/*Case 1 : 5+2 */
	ast1Left := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}
	ast1Right := &ast.Node{Value: ast.IntegerExpression{Value: 2}, Left: nil, Right: nil}

	ast1 := &ast.Node{Value: ast.Operator{tokens.Plus}, Left: ast1Left, Right: ast1Right}

	/*Case 2 : 5+2 == 3+4 || true*/
	ast2Left1 := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}
	ast2Right1 := &ast.Node{Value: ast.IntegerExpression{Value: 2}, Left: nil, Right: nil}

	ast2L := &ast.Node{Value: ast.Operator{tokens.Plus}, Left: ast2Left1, Right: ast2Right1}

	ast2Left2 := &ast.Node{Value: ast.IntegerExpression{Value: 3}, Left: nil, Right: nil}
	ast2Right2 := &ast.Node{Value: ast.IntegerExpression{Value: 4}, Left: nil, Right: nil}

	ast2R := &ast.Node{Value: ast.Operator{tokens.Plus}, Left: ast2Left2, Right: ast2Right2}

	ast2LL := &ast.Node{Value: ast.Operator{tokens.EqualTo}, Left: ast2L, Right: ast2R}

	ast2RR := &ast.Node{Value: ast.BooleanExpression{Value: true}, Left: nil, Right: nil}

	ast2 := &ast.Node{Value: ast.Operator{tokens.LogicalOr}, Left: ast2LL, Right: ast2RR}

	/*Case 3 : 5||2*/
	ast3Left := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}
	ast3Right := &ast.Node{Value: ast.IntegerExpression{Value: 2}, Left: nil, Right: nil}

	ast3 := &ast.Node{Value: ast.Operator{tokens.LogicalOr}, Left: ast3Left, Right: ast3Right}

	/*Case 4 : 5==2*/
	ast4Left := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}
	ast4Right := &ast.Node{Value: ast.IntegerExpression{Value: 2}, Left: nil, Right: nil}

	ast4 := &ast.Node{Value: ast.Operator{tokens.EqualTo}, Left: ast4Left, Right: ast4Right}

	/*Case 5 : 5-2 + 1 * 3/4*/
	ast5Left1 := &ast.Node{Value: ast.IntegerExpression{Value: 3}, Left: nil, Right: nil}
	ast5Right1 := &ast.Node{Value: ast.IntegerExpression{Value: 4}, Left: nil, Right: nil}

	ast5R := &ast.Node{Value: ast.Operator{tokens.Divide}, Left: ast5Left1, Right: ast5Right1}

	ast5Left2 := &ast.Node{Value: ast.IntegerExpression{Value: 1}, Left: nil, Right: nil}

	ast5R2 := &ast.Node{Value: ast.Operator{tokens.Multiply}, Left: ast5Left2, Right: ast5R}

	ast5Left3 := &ast.Node{Value: ast.IntegerExpression{Value: 1}, Left: nil, Right: nil}

	ast5R3 := &ast.Node{Value: ast.Operator{tokens.Plus}, Left: ast5Left3, Right: ast5R2}

	ast5Left4 := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}

	ast5 := &ast.Node{Value: ast.Operator{tokens.Minus}, Left: ast5Left4, Right: ast5R3}

	tests := []TestCase{{ast1, Int, ""},
		{ast2, Bool, ""},
		{ast3, Invalid, "Error:"},
		{ast4, Bool, ""},
		{ast5, Int, ""}}

	runTests(tests, t)
}

func runTests(tests []TestCase, t *testing.T) {
	for i, test := range tests {
		actualType, _ := TypeCheck(test.Node, "")

		if actualType == test.ExpectedType {
			t.Logf("test %d pass, actual type %s, expected type %s: ", i, actualType, test.ExpectedType)
			continue
		} else {
			t.Fatalf("test %d fail, actual type %s, expected type %s: ", i, actualType, test.ExpectedType)
		}
	}
}
