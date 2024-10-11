package purefunction

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// Pure examines if the given function is pure or not. No need to pass in existing pure function declarations, it can be skipped.
func Pure(funcDecl *ast.FuncDecl, existingPureFuncDecls ...*ast.FuncDecl) bool {
	if funcDecl.Recv != nil { // not a pure function but a method
		return false
	}
	get := func(name *ast.Ident) (*ast.FuncDecl, bool) {
		for _, v := range existingPureFuncDecls {
			if v.Name == name {
				return v, true
			}
		}
		return nil, false
	}
	pure := true
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		switch stmt := n.(type) {
		case *ast.CallExpr:
			if ident, ok := stmt.Fun.(*ast.Ident); ok {
				if calledFunc, exists := get(ident); exists {
					if !Pure(calledFunc, existingPureFuncDecls...) {
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

// PureFunctionDecls returns the pure *ast.FuncDecl functions from the given source file
func PureFunctionDecls(filename string) ([]*ast.FuncDecl, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []*ast.FuncDecl{}, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, file, parser.AllErrors)
	if err != nil {
		return []*ast.FuncDecl{}, fmt.Errorf("failed to parse file: %v", err)
	}
	var pureFunctions []*ast.FuncDecl
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if Pure(funcDecl) {
			pureFunctions = append(pureFunctions, funcDecl)
		}
	}
	return pureFunctions, nil // success
}
