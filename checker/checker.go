package checker

import (
	"AST/ast"
	"AST/tokens"
	"reflect"
)

/*
Function responsible for type checking the AST.
Expecting the root AST Node along with the Syntax type expected by the parent node.
Returns the Syntax type of the child nodes, so that the node can check with the expected syntax type of the parent.
Also returns the error in case of type mismatch
*/
func TypeCheck(root *ast.Node, expectedType SyntaxType) (SyntaxType, string) {

	if root == nil {
		return Invalid, "Error: Empty Expression"
	}

	//Change to Switch Statements
	syntaxType := reflect.TypeOf(root.Value)

	switch syntaxType.Name() {
	case string(Op):
		//If operator is one of Bool or Int operator, pass it on to the recursive call
		//Need to have left and right node
		if root.Left != nil && root.Right != nil {
			currentOperatorToken := root.Value.(ast.Operator).Value
			//Get the expected syntax type for the operator token to pass it on to the recursive call
			currentNodeSyntaxType, err := CheckTokenType(currentOperatorToken)

			if err != "" {
				return Invalid, err
			}
			nodeLeftSyntaxType, err := TypeCheck(root.Left, currentNodeSyntaxType)
			if err != "" {
				return Invalid, err
			}

			nodeRightSyntaxType, err := TypeCheck(root.Right, currentNodeSyntaxType)
			if err != "" {
				return Invalid, err
			}

			//Both Left and Right Syntax Type should be same
			if nodeLeftSyntaxType == nodeRightSyntaxType {
				//For root node - expectedType will be ""
				if expectedType == "" {
					//If both child node are IntegerExpression and Operator is EqualTo or NotEqualTo - Valid case
					if (nodeLeftSyntaxType == Int && nodeRightSyntaxType == Int) && (currentOperatorToken == tokens.EqualTo || currentOperatorToken == tokens.NotEqualTo) {
						return Bool, ""
					}
					//If both child node are IntegerExpression and Operator is LogicalOr or LogicalAnd - Invalid case
					if (nodeLeftSyntaxType == Int && nodeRightSyntaxType == Int) && (currentOperatorToken == tokens.LogicalOr || currentOperatorToken == tokens.LogicalAnd) {
						return Invalid, "Error: Expression is not valid. Left and Right Node Mismatch with Operator " + string(nodeLeftSyntaxType) + string(nodeRightSyntaxType) + string(currentOperatorToken)
					}
					//return any of left or right syntax type as both are same
					return nodeLeftSyntaxType, ""
				}

				//If both child node are IntegerExpression and Operator is EqualTo or NotEqualTo - Valid case
				if (nodeLeftSyntaxType == Int && nodeRightSyntaxType == Int) && (currentOperatorToken == tokens.EqualTo || currentOperatorToken == tokens.NotEqualTo) {
					return expectedType, ""
				}
				//If both child node are IntegerExpression and Operator is LogicalOr or LogicalAnd - Invalid case
				if (nodeLeftSyntaxType == Int && nodeRightSyntaxType == Int) && (currentOperatorToken == tokens.LogicalOr || currentOperatorToken == tokens.LogicalAnd) {
					return Invalid, "Error: Expression is not valid. Left and Right Node Mismatch with Operator " + string(nodeLeftSyntaxType) + string(nodeRightSyntaxType) + string(currentOperatorToken)
				}
				// Both Left and Right Syntax should be equal to expected Syntax
				//If the expected type is BooleanExpression and both child are same type (IntegerExpression), it should return expected type based on currentOperatorToken
				//Result of the child node need to check if expected parent and expected child are same type
				if expectedType == nodeLeftSyntaxType && expectedType == nodeRightSyntaxType || expectedType == Bool {
					return expectedType, ""
				} else {
					return Invalid, "Error: Expression is not valid. Left and Right Node Mismatch with Expected Node" + string(nodeLeftSyntaxType) + string(nodeRightSyntaxType) + string(expectedType)
				}

			} else {
				return Invalid, "Error: Expression is not valid. Left and Right Node Mismatch " + string(nodeLeftSyntaxType) + string(nodeRightSyntaxType)
			}

		} else {
			return Invalid, "Error: Expression is not valid"
		}

	//If it is BooleanExpression - reached the leaf node
	case string(Bool):
		return Bool, ""
	//If it is IntegerExpression - reached the leaf node
	case string(Int):
		return Int, ""
	//Default Condition when nothing matches
	default:
		return Invalid, "Error: Expression is not valid"

	}

}
