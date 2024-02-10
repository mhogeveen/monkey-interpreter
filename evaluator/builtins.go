package evaluator

import "monkey-interpreter/object"

var builtins = map[string]*object.Builtin{
	"len": newBuiltin(builtinLen),
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
	default:
		return newError("argument to 'len' not supported, received %s", args[0].Type())
	}
}
