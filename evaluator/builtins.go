package evaluator

import "monkey-interpreter/object"

var builtins = map[string]*object.Builtin{
	"len":   newBuiltin(builtinLen),
	"first": newBuiltin(builtinFirst),
	"last":  newBuiltin(builtinLast),
	"rest":  newBuiltin(builtinRest),
	"push":  newBuiltin(builtinPush),
}

func newBuiltin(fn object.BuiltinFunction) *object.Builtin {
	return &object.Builtin{
		Fn: fn,
	}
}

func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. received=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to 'len' not supported, received %s", args[0].Type())
	}
}

func builtinFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. received=%q, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to 'first' must be ARRAY, received %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func builtinLast(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. received=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to 'last' must be ARRAY, received %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return NULL
}

func builtinRest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. received=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to 'rest' must be ARRAY, received %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}

	return NULL
}

func builtinPush(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. received=%d, want=2", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to 'push' must be ARRAY, received %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	newElements := make([]object.Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}
