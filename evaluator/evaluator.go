package evaluator

import (
	"github.com/valter4578/vlang/ast"
	"github.com/valter4578/vlang/object"
)

var (
	trueRef  = &object.Boolean{Value: true}
	falseRef = &object.Boolean{Value: false}
	nullRef  = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return trueRef
	} else {
		return falseRef
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	default:
		return nullRef
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case trueRef:
		return falseRef
	case falseRef:
		return trueRef
	case nullRef:
		return trueRef
	default:
		return falseRef
	}
}
