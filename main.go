package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	//"go/printer"
	//"os"
)

func main() {
    filename := "data/main.go"

	fset := token.NewFileSet()

	//node, err := parser.ParseFile(fset, filename, nil, parser.Trace)
	//node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	node, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println("Imports:")
	//for _, i := range node.Imports {
	//	fmt.Println(i.Path.Value)
	//}

	fmt.Println("Functions:")
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println(fn.Name.Name)

		// Examine the function type. If non-simple types are passed in, they must not be modified in the function body.
		fmt.Println(fn.Type) // FuncType

		// Examine the function body. If non-pure functions are called, the function is not pure.
		fmt.Println(fn.Body) // BlockStmt
	}
}
