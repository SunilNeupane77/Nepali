// Package interpreter implements the interpreter for the Nepali programming language
package interpreter

import (
	"fmt"

	"github.com/SunilNeupane77/nepali/internal/ast"
	"github.com/SunilNeupane77/nepali/internal/object"
)

// Interpreter represents the interpreter
var Interpreter = &interpreter{}

type interpreter struct{}

// Eval evaluates an AST node
func (i *interpreter) Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return i.evalProgram(node, env)
	case *ast.LetStatement:
		return i.evalLetStatement(node, env)
	case *ast.ReturnStatement:
		return i.evalReturnStatement(node, env)
	case *ast.ExpressionStatement:
		return i.Eval(node.Expression, env)
	case *ast.Identifier:
		return i.evalIdentifier(node, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return i.nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		return i.evalPrefixExpression(node, env)
	case *ast.InfixExpression:
		return i.evalInfixExpression(node, env)
	case *ast.IfExpression:
		return i.evalIfExpression(node, env)
	case *ast.BlockStatement:
		return i.evalBlockStatement(node, env)
	case *ast.FunctionLiteral:
		params := i.evalFunctionParameters(node.Parameters)
		return &object.Function{
			Parameters: params,
			Body:      node.Body,
			Env:       env,
		}
	case *ast.CallExpression:
		return i.evalCallExpression(node, env)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.ArrayLiteral:
		return i.evalArrayLiteral(node, env)
	case *ast.IndexExpression:
		return i.evalIndexExpression(node, env)
	case *ast.HashLiteral:
		return i.evalHashLiteral(node, env)
	}

	return nil
}

func (i *interpreter) evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = i.Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func (i *interpreter) evalLetStatement(let *ast.LetStatement, env *object.Environment) object.Object {
	val := i.Eval(let.Value, env)
	if isError(val) {
		return val
	}

	env.Set(let.Name.Value, val)
	return nil
}

func (i *interpreter) evalReturnStatement(ret *ast.ReturnStatement, env *object.Environment) object.Object {
	val := i.Eval(ret.ReturnValue, env)
	if isError(val) {
		return val
	}

	return &object.ReturnValue{Value: val}
}

func (i *interpreter) evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: " + node.Value)
}

func (i *interpreter) evalPrefixExpression(node *ast.PrefixExpression, env *object.Environment) object.Object {
	right := i.Eval(node.Right, env)
	if isError(right) {
		return right
	}

	switch node.Operator {
	case "!":
		return i.evalBangOperatorExpression(right)
	case "-":
		return i.evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: " + node.Operator + right.Type())
	}
}

func (i *interpreter) evalInfixExpression(node *ast.InfixExpression, env *object.Environment) object.Object {
	left := i.Eval(node.Left, env)
	if isError(left) {
		return left
	}

	right := i.Eval(node.Right, env)
	if isError(right) {
		return right
	}

	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return i.evalIntegerInfixExpression(node, left, right)
	case node.Operator == "==":
		return i.nativeBoolToBooleanObject(left == right)
	case node.Operator == "!=":
		return i.nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: " + left.Type() + " " + node.Operator + " " + right.Type())
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return i.evalStringInfixExpression(node, left, right)
	case left.Type() == object.ARRAY_OBJ && node.Operator == "==":
		return i.nativeBoolToBooleanObject(left == right)
	case left.Type() == object.ARRAY_OBJ && node.Operator == "!=":
		return i.nativeBoolToBooleanObject(left != right)
	case left.Type() == object.HASH_OBJ && node.Operator == "==":
		return i.nativeBoolToBooleanObject(left == right)
	case left.Type() == object.HASH_OBJ && node.Operator == "!=":
		return i.nativeBoolToBooleanObject(left != right)
	default:
		return newError("unknown operator: " + left.Type() + " " + node.Operator + " " + right.Type())
	}
}

func (i *interpreter) evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := i.Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return i.Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return i.Eval(ie.Alternative, env)
	}

	return nil
}

func (i *interpreter) evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = i.Eval(statement, env)

		if result != nil {
			switch result := result.(type) {
			case *object.ReturnValue:
				return result.Value
			case *object.Error:
				return result
			}
		}
	}

	return result
}

func (i *interpreter) evalFunctionParameters(params []*ast.Identifier) []*object.Parameter {
	result := make([]*object.Parameter, len(params))

	for i, param := range params {
		result[i] = &object.Parameter{Name: param.Value}
	}

	return result
}

func (i *interpreter) evalCallExpression(node *ast.CallExpression, env *object.Environment) object.Object {
	function := i.Eval(node.Function, env)
	if isError(function) {
		return function
	}

	args := i.evalExpressions(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	return i.applyFunction(function, args)
}

func (i *interpreter) evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		result = append(result, i.Eval(e, env))
		if isError(result[len(result)-1]) {
			return result
		}
	}

	return result
}

func (i *interpreter) evalArrayLiteral(node *ast.ArrayLiteral, env *object.Environment) object.Object {
	elements := i.evalExpressions(node.Elements, env)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}

	return &object.Array{Elements: elements}
}

func (i *interpreter) evalIndexExpression(node *ast.IndexExpression, env *object.Environment) object.Object {
	left := i.Eval(node.Left, env)
	if isError(left) {
		return left
	}

	right := i.Eval(node.Index, env)
	if isError(right) {
		return right
	}

	return i.evalIndex(left, right)
}

func (i *interpreter) evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := i.Eval(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: " + key.Type())
		}

		value := i.Eval(valueNode, env)
		if isError(value) {
			return value
		}

		pairs[hashKey.HashKey()] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

func (i *interpreter) evalIntegerInfixExpression(node *ast.InfixExpression, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch node.Operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return i.nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return i.nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return i.nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return i.nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: " + left.Type() + " " + node.Operator + " " + right.Type())
	}
}

func (i *interpreter) evalStringInfixExpression(node *ast.InfixExpression, left, right object.Object) object.Object {
	if node.Operator != "" {
		return newError("unknown operator: " + left.Type() + " " + node.Operator + " " + right.Type())
	}

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
}

func (i *interpreter) evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case nil:
		return TRUE
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	default:
		return FALSE
	}
}

func (i *interpreter) evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -" + right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func (i *interpreter) nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func (i *interpreter) isTruthy(obj object.Object) bool {
	switch obj {
	case nil:
		return false
	case FALSE:
		return false
	default:
		return true
	}
}

func (i *interpreter) applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := i.extendFunctionEnv(fn, args)
		evaluated := i.Eval(fn.Body, extendedEnv)
		return i.unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: " + fn.Type())
	}
}

func (i *interpreter) extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Name, args[paramIdx])
	}

	return env
}

func (i *interpreter) unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func (i *interpreter) evalIndex(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return i.evalArrayIndex(left, index)
	case left.Type() == object.HASH_OBJ:
		return i.evalHashIndex(left, index)
	default:
		return newError("index operator not supported: " + left.Type())
	}
}

func (i *interpreter) evalArrayIndex(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value

	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return nil
	}

	return arrayObject.Elements[idx]
}

func (i *interpreter) evalHashIndex(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)
	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: " + index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return nil
	}

	return pair.Value
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

// TRUE represents the boolean true value
var TRUE = &object.Boolean{Value: true}

// FALSE represents the boolean false value
var FALSE = &object.Boolean{Value: false}

// NULL represents the null value
var NULL = &object.Null{}
