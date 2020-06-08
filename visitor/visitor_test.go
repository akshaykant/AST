package visitor

import (
	"AST/ast"
	"AST/checker"
	"AST/tokens"
	"reflect"
	"testing"
)

type TestCase struct {
	Node          *ast.Node
	ExpectedValue interface{}
	ExpectedType  checker.SyntaxType
	ExpectedError string
	res           Visitable
	visit         Visitor
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

	/*Case 3 : 5-2*/
	ast3Left := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}
	ast3Right := &ast.Node{Value: ast.IntegerExpression{Value: 2}, Left: nil, Right: nil}

	ast3 := &ast.Node{Value: ast.Operator{tokens.Minus}, Left: ast3Left, Right: ast3Right}

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

	/*Case 6 : 5+2 */
	ast6Left := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}
	ast6Right := &ast.Node{Value: ast.IntegerExpression{Value: 2}, Left: nil, Right: nil}

	ast6 := &ast.Node{Value: ast.Operator{tokens.Plus}, Left: ast6Left, Right: ast6Right}

	/*Case 7 : 5+2 == 3+4 || true*/
	ast7Left1 := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}
	ast7Right1 := &ast.Node{Value: ast.IntegerExpression{Value: 2}, Left: nil, Right: nil}

	ast7L := &ast.Node{Value: ast.Operator{tokens.Plus}, Left: ast7Left1, Right: ast7Right1}

	ast7Left2 := &ast.Node{Value: ast.IntegerExpression{Value: 3}, Left: nil, Right: nil}
	ast7Right2 := &ast.Node{Value: ast.IntegerExpression{Value: 4}, Left: nil, Right: nil}

	ast7R := &ast.Node{Value: ast.Operator{tokens.Plus}, Left: ast7Left2, Right: ast7Right2}

	ast7LL := &ast.Node{Value: ast.Operator{tokens.EqualTo}, Left: ast7L, Right: ast7R}

	ast7RR := &ast.Node{Value: ast.BooleanExpression{Value: true}, Left: nil, Right: nil}

	ast7 := &ast.Node{Value: ast.Operator{tokens.LogicalOr}, Left: ast7LL, Right: ast7RR}

	/*Case 8 : 5||2*/
	ast8Left := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}
	ast8Right := &ast.Node{Value: ast.IntegerExpression{Value: 2}, Left: nil, Right: nil}

	ast8 := &ast.Node{Value: ast.Operator{tokens.LogicalOr}, Left: ast8Left, Right: ast8Right}

	/*Case 9 : 5==2*/
	ast9Left := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}
	ast9Right := &ast.Node{Value: ast.IntegerExpression{Value: 2}, Left: nil, Right: nil}

	ast9 := &ast.Node{Value: ast.Operator{tokens.EqualTo}, Left: ast9Left, Right: ast9Right}

	/*Case 10 : 5-2 + 1 * 3/4*/
	ast10Left1 := &ast.Node{Value: ast.IntegerExpression{Value: 3}, Left: nil, Right: nil}
	ast10Right1 := &ast.Node{Value: ast.IntegerExpression{Value: 4}, Left: nil, Right: nil}

	ast10R := &ast.Node{Value: ast.Operator{tokens.Divide}, Left: ast10Left1, Right: ast10Right1}

	ast10Left2 := &ast.Node{Value: ast.IntegerExpression{Value: 1}, Left: nil, Right: nil}

	ast10R2 := &ast.Node{Value: ast.Operator{tokens.Multiply}, Left: ast10Left2, Right: ast10R}

	ast10Left3 := &ast.Node{Value: ast.IntegerExpression{Value: 1}, Left: nil, Right: nil}

	ast10R3 := &ast.Node{Value: ast.Operator{tokens.Plus}, Left: ast10Left3, Right: ast10R2}

	ast10Left4 := &ast.Node{Value: ast.IntegerExpression{Value: 5}, Left: nil, Right: nil}

	ast10 := &ast.Node{Value: ast.Operator{tokens.Minus}, Left: ast10Left4, Right: ast10R3}

	tests := []TestCase{{ast1, ast.IntegerExpression{7}, checker.Int, "", EvalInt{}, &EvaluateIntVisitor{}},
		{ast2, ast.BooleanExpression{true}, checker.Bool, "", EvalBool{}, &EvaluateBoolVisitor{}},
		{ast3, ast.IntegerExpression{3}, checker.Int, "", EvalInt{}, &EvaluateIntVisitor{}},
		{ast4, ast.BooleanExpression{false}, checker.Bool, "", EvalBool{}, &EvaluateBoolVisitor{}},
		{ast5, ast.IntegerExpression{4}, checker.Int, "", EvalInt{}, &EvaluateIntVisitor{}},
		{ast6, nil, checker.Int, "", CheckInt{}, &CheckerIntVisitor{}},
		{ast7, nil, checker.Bool, "", CheckBool{}, &CheckerBoolVisitor{}},
		{ast8, nil, checker.Invalid, "Error:", CheckInt{}, &CheckerIntVisitor{}},
		{ast9, nil, checker.Bool, "", CheckBool{}, &CheckerBoolVisitor{}},
		{ast10, nil, checker.Int, "", CheckInt{}, &CheckerIntVisitor{}},
	}

	runTests(tests, t)
}

func runTests(tests []TestCase, t *testing.T) {
	for i, test := range tests {
		res := test.res.Accept(test.visit, test.Node)

		ch := reflect.TypeOf(res).Name()

		switch ch {

		case "EvalInt":

			actualValue := res.(EvalInt).SyntaxType

			if actualValue == test.ExpectedType {
				t.Logf("test %d pass, actual type %s, expected type %s, value %s ", i, actualValue, test.ExpectedType, test.ExpectedValue)
				continue
			} else {
				t.Fatalf("test %d fail, actual type %s, expected type %s: ", i, actualValue, test.ExpectedType)
			}

		case "EvalBool":
			actualValue := res.(EvalBool).SyntaxType

			if actualValue == test.ExpectedType {
				t.Logf("test %d pass, actual type %s, expected type %s, value %s ", i, actualValue, test.ExpectedType, test.ExpectedValue)
				continue
			} else {
				t.Fatalf("test %d fail, actual type %s, expected type %s: ", i, actualValue, test.ExpectedType)
			}

		case "CheckInt":
			actualValue := res.(CheckInt).SyntaxType

			if actualValue == test.ExpectedType {
				t.Logf("test %d pass, actual type %s, expected type %s, error %s ", i, actualValue, test.ExpectedType, test.ExpectedError)
				continue
			} else {
				t.Fatalf("test %d fail, actual type %s, expected type %s: ", i, actualValue, test.ExpectedError)
			}

		case "CheckBool":
			actualValue := res.(CheckBool).SyntaxType

			if actualValue == test.ExpectedType {
				t.Logf("test %d pass, actual type %s, expected type %s, error %s ", i, actualValue, test.ExpectedType, test.ExpectedError)
				continue
			} else {
				t.Fatalf("test %d fail, actual type %s, expected type %s: ", i, actualValue, test.ExpectedError)
			}
		}
	}
}
