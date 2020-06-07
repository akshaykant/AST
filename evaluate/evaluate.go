package evaluate

import (
	"AST/ast"
	"AST/checker"
	"AST/tokens"
	"reflect"
)

func Evaluate(root *ast.Node) interface{} {

	syntaxType := reflect.TypeOf(root.Value)

	if syntaxType.Name() == string(checker.Bool) {
		return root.Value
	}

	if syntaxType.Name() == string(checker.Int) {
		return root.Value
	}

	if syntaxType.Name() == string(checker.Op) {

		operatorToken := root.Value.(ast.Operator).Value

		leftValue := Evaluate(root.Left)

		rightValue := Evaluate(root.Right)

		switch operatorToken {

		case tokens.Plus:
			return ast.IntegerExpression{leftValue.(ast.IntegerExpression).Value + rightValue.(ast.IntegerExpression).Value}

		case tokens.Minus:
			return ast.IntegerExpression{leftValue.(ast.IntegerExpression).Value - rightValue.(ast.IntegerExpression).Value}

		case tokens.Multiply:
			return ast.IntegerExpression{leftValue.(ast.IntegerExpression).Value * rightValue.(ast.IntegerExpression).Value}

		case tokens.Divide:
			return ast.IntegerExpression{leftValue.(ast.IntegerExpression).Value / rightValue.(ast.IntegerExpression).Value}

		case tokens.EqualTo:
			syntaxTypeLeft := reflect.TypeOf(leftValue)
			syntaxTypeRight := reflect.TypeOf(rightValue)

			if syntaxTypeLeft.Name() == string(checker.Bool) && syntaxTypeRight.Name() == string(checker.Bool) {
				return ast.BooleanExpression{leftValue.(ast.BooleanExpression).Value == rightValue.(ast.BooleanExpression).Value}
			}

			if syntaxTypeLeft.Name() == string(checker.Int) && syntaxTypeRight.Name() == string(checker.Int) {
				return ast.BooleanExpression{leftValue.(ast.IntegerExpression).Value == rightValue.(ast.IntegerExpression).Value}
			}

		case tokens.NotEqualTo:
			syntaxTypeLeft := reflect.TypeOf(leftValue)
			syntaxTypeRight := reflect.TypeOf(rightValue)

			if syntaxTypeLeft.Name() == string(checker.Bool) && syntaxTypeRight.Name() == string(checker.Bool) {
				return ast.BooleanExpression{leftValue.(ast.BooleanExpression).Value != rightValue.(ast.BooleanExpression).Value}
			}

			if syntaxTypeLeft.Name() == string(checker.Int) && syntaxTypeRight.Name() == string(checker.Int) {
				return ast.BooleanExpression{leftValue.(ast.IntegerExpression).Value != rightValue.(ast.IntegerExpression).Value}
			}

		case tokens.LogicalAnd:
			return ast.BooleanExpression{leftValue.(ast.BooleanExpression).Value && rightValue.(ast.BooleanExpression).Value}

		case tokens.LogicalOr:
			return ast.BooleanExpression{leftValue.(ast.BooleanExpression).Value || rightValue.(ast.BooleanExpression).Value}

		}
	}
	return nil
}
