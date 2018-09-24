package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"log"
	//"go/printer"
	//"os"
)

var basicTypes = []string{
	"bool",
	"string",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"uintptr",
	"byte",
	"rune",
	"float32",
	"float64",
	"complex64",
	"complex128",
}

// Check if a given slice of strings has the given element
func has(sl []string, e string) bool {
	for _, s := range sl {
		if s == e {
			return true
		}
	}
	return false
}

func main() {
	filename := "data/main.go"

	fset := token.NewFileSet()

	// parser.Trace and ParseComments is also possible flags
	node, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		log.Fatalln(err)
	}

	// First gather all defined types
	for _, f := range node.Decls {
		gd, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		// the token, which may be a TypeSpec
		tok := gd.Tok
		if tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			// For each TypeSpec
			name := ts.Name // Ident
			typ := ts.Type

			// TODO: Find a smoother way to convert to string
			nameString := fmt.Sprintf("%s", name)
			typeString := fmt.Sprintf("%s", typ)

			// Collect all type aliases for basic types into the basicTypes slice
			if has(basicTypes, typeString) {
				basicTypes = append(basicTypes, nameString)
			}
		}
	}

	fmt.Println("Functions:")
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		functionName := fn.Name.Name
		fmt.Println("--- " + functionName + " ---")

		// Examine the function type. If non-simple types are passed in, they must not be modified in the function body.
		for _, arg := range fn.Type.Params.List {
			// output "name: type" or "name, name, ...: type"
			var argNames []string
			for _, argIdent := range arg.Names {
				argNames = append(argNames, argIdent.String())
			}
			ident, ok := arg.Type.(*ast.Ident)
			if !ok {
				fmt.Println("TYPE IS NOT *ast.Ident!")
				fmt.Printf("%s %v %T\n", arg.Type, arg.Type, arg.Type)
				panic("TO IMPLEMENT")
			}
			typeName := ident.Name
			if has(basicTypes, typeName) {
				fmt.Println("BASIC TYPE", typeName)
			} else {
				// TODO: Also look at typedefs, perhaps ast has support for this?
				fmt.Println("NOT A BASIC TYPE", typeName)
			}
			fmt.Println(strings.Join(argNames, ", ") + ": " + fmt.Sprint(arg.Type))
		}

		// Examine the function body. If non-pure functions are called, the function is not pure.
		for _, s := range fn.Body.List {
			fmt.Printf("%s %v %T\n", s, s, s)
		}
	}
}
