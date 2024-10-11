package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

var funcMap = make(map[string]*ast.FuncDecl)

func Pure(funcDecl *ast.FuncDecl) bool {
	if funcDecl.Recv != nil { // not a function but a method
		return false
	}
	pure := true
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		switch stmt := n.(type) {
		case *ast.CallExpr:
			if ident, ok := stmt.Fun.(*ast.Ident); ok {
				if calledFunc, exists := funcMap[ident.Name]; exists {
					if !Pure(calledFunc) {
						pure = false
						return false // stop iterating
					}
				} else {
					pure = false
					return false // stop iterating
				}
			}
		case *ast.AssignStmt:
			for _, lhs := range stmt.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok {
					if ident.Obj == nil || ident.Obj.Kind != ast.Var {
						pure = false // not pure
						return false // stop iterating
					}
				}
			}
		}
		return true // continue to iterate
	})
	return pure
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("Usage: %s <go_source_file>\n", os.Args[0])
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, file, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("Failed to parse file: %v", err)
	}

	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if Pure(funcDecl) {
			fmt.Println(funcDecl.Name.Name)
		}
	}

	return nil // success
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
