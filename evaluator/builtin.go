package evaluator

import (
	"fmt"
	"github.com/valter4578/vlang/object"
)

var builtins = map[string]*object.BuiltFunc{
	"len": &object.BuiltFunc{
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.BuiltFunc{
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.BuiltFunc{
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			lenght := len(arr.Elements)
			if lenght > 0 {
				return arr.Elements[lenght-1]
			}

			return NULL
		},
	},
	"tail": &object.BuiltFunc{
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `tail` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])

				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"append": &object.BuiltFunc{
		Func: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `tail` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			newElement := args[1]
			arr.Elements = append(arr.Elements, newElement)

			return arr
		},
	},

	"print": &object.BuiltFunc{
		Func: func(args ...object.Object) object.Object {
			for _, argument := range args {
				fmt.Println(argument.Inspect())
			}
			return NULL
		},
	},
}
