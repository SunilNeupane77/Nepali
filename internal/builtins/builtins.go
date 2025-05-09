// Package builtins implements built-in functions for the Nepali programming language
package builtins

import "fmt"

import "github.com/SunilNeupane77/nepali/internal/object"


var builtins = map[string]*object.Builtin{
	"लेन": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `लेन` must be STRING or ARRAY, got %s", args[0].Type())
			}
		},
	},
	"प्रिन्ट": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Printf("%s ", arg.Inspect())
			}
			fmt.Printf("\n")
			return object.NULL
		},
	},
	"प्रिन्टल": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Printf("%s\n", arg.Inspect())
			}
			return object.NULL
		},
	},
	"टाइप": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			return &object.String{Value: strings.ToLower(args[0].Type())}
		},
	},
	"स्ट्रिंग": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			return &object.String{Value: args[0].Inspect()}
		},
	},
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func Lookup(name string) object.Object {
	if builtin, ok := builtins[name]; ok {
		return builtin
	}
	return nil
}
