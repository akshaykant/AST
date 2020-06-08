package visitor

import (
	"AST/ast"
	"AST/checker"
	"AST/evaluate"
	"fmt"
)

//Visitable interface to accept different visitors
type Visitable interface {
	Accept(Visitor, *ast.Node) Visitable
}

//Visitor interface to handle response from the Visitable
type Visitor interface {
	Visit(Visitable)
}

//Different type of visitors for different tasks
type EvaluateBoolVisitor struct{}

type EvaluateIntVisitor struct{}

type CheckerBoolVisitor struct{}

type CheckerIntVisitor struct{}

//create struct for result of Bool and Int Evaluation and Type Check
type EvalBool struct {
	Error      string
	Result     bool
	SyntaxType checker.SyntaxType
}

type EvalInt struct {
	Error      string
	Result     int
	SyntaxType checker.SyntaxType
}

type CheckBool struct {
	Error      string
	SyntaxType checker.SyntaxType
}

type CheckInt struct {
	Error      string
	SyntaxType checker.SyntaxType
}

func (res EvalBool) Accept(visitor Visitor, node *ast.Node) Visitable {
	//Type Check, if the expression is correct and can be evaluated
	syntaxType, err := checker.TypeCheck(node, "")

	//Set the value incase of error
	if err != "" {
		res.Error = err
		res.SyntaxType = syntaxType
	} else {
		//Evaluate the expression and set the values.
		result := evaluate.Evaluate(node)
		res.Result = result.(ast.BooleanExpression).Value
		res.SyntaxType = syntaxType
	}

	visitor.Visit(res)
	return res

}

func (res EvalInt) Accept(visitor Visitor, node *ast.Node) Visitable {
	//Type Check, if the expression is correct and can be evaluated
	syntaxType, err := checker.TypeCheck(node, "")

	//Set the value incase of error
	if err != "" {
		res.Error = err
		res.SyntaxType = syntaxType
	} else {
		//Evaluate the expression and set the values.
		result := evaluate.Evaluate(node)
		res.Result = result.(ast.IntegerExpression).Value
		res.SyntaxType = syntaxType
	}

	visitor.Visit(res)
	return res
}

func (res CheckBool) Accept(visitor Visitor, node *ast.Node) Visitable {
	//Type Check, if the expression is correct and can be evaluated
	syntaxType, err := checker.TypeCheck(node, "")

	res.SyntaxType = syntaxType
	//Set the value incase of error
	if err != "" {
		res.Error = err
	}

	visitor.Visit(res)
	return res

}

func (res CheckInt) Accept(visitor Visitor, node *ast.Node) Visitable {
	//Type Check, if the expression is correct and can be evaluated
	syntaxType, err := checker.TypeCheck(node, "")

	res.SyntaxType = syntaxType
	//Set the value incase of error
	if err != "" {
		res.Error = err
	}

	visitor.Visit(res)
	return res

}

func (*EvaluateBoolVisitor) Visit(visitable Visitable) {

	res := visitable

	if res.(EvalBool).Error != "" {
		fmt.Print(res.(EvalBool).SyntaxType)
		fmt.Print(res.(EvalBool).Result)
	} else {
		fmt.Print(res.(EvalBool).SyntaxType)
		fmt.Print(res.(EvalBool).Error)
	}

}

func (*EvaluateIntVisitor) Visit(visitable Visitable) {

	res := visitable

	if res.(EvalInt).Error != "" {
		fmt.Print(res.(EvalInt).SyntaxType)
		fmt.Print(res.(EvalInt).Result)
	} else {
		fmt.Print(res.(EvalInt).SyntaxType)
		fmt.Print(res.(EvalInt).Error)
	}

}

func (*CheckerBoolVisitor) Visit(visitable Visitable) {

	res := visitable

	fmt.Print(res.(CheckBool).SyntaxType)
	if res.(CheckBool).Error != "" {
		fmt.Print(res.(CheckBool).SyntaxType)
	} else {
		fmt.Print(res.(CheckBool).SyntaxType)
		fmt.Print(res.(CheckBool).Error)
	}

}

func (*CheckerIntVisitor) Visit(visitable Visitable) {

	res := visitable

	fmt.Print(res.(CheckInt).SyntaxType)
	if res.(CheckInt).Error != "" {
		fmt.Print(res.(CheckInt).SyntaxType)
	} else {
		fmt.Print(res.(CheckInt).SyntaxType)
		fmt.Print(res.(CheckInt).Error)
	}

}
