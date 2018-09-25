package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
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

type FuncBodyVisitor struct {
	pure         map[string]bool // for keeping the purity status of all known functions
	functionName string          // for the name of the function that is to be examined
	idents       []string        // for gathering all identifiers used in a function
	created      *[]string       // all identifiers that were created in this function
	verbose      bool
}

func NewFuncBodyVisitor(pureMap map[string]bool, functionName string, created *[]string, verbose bool) *FuncBodyVisitor {
	idents := []string{}
	return &FuncBodyVisitor{pureMap, functionName, idents, created, verbose}
}

func (v *FuncBodyVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return nil
	}
	switch t := node.(type) {
	case *ast.Ident:
		//fmt.Println("IDENT", t)
		// Gather all IDENTS in a slice
		v.idents = append(v.idents, fmt.Sprintf("%s", t))
	case *ast.CallExpr:
		//fmt.Println("case *ast.CallExpr")
		if fun, ok := t.Fun.(*ast.SelectorExpr); ok {
			fName := fun.Sel.Name
			if ok, isPure := v.pure[fName]; ok && isPure {
				if v.verbose {
					fmt.Println(v.functionName + " is calling an unproblematic function: " + fName)
				}
			} else {
				if v.verbose {
					fmt.Println(v.functionName + " is calling a function that might be unpure: " + fName)
				}
				v.pure[v.functionName] = false
				return nil
			}
		} else {
			fName := fmt.Sprintf("%s", t.Fun)
			if ok, isPure := v.pure[fName]; ok && isPure {
				if v.verbose {
					fmt.Println("pure call: " + fName)
				}
			} else {
				if !has(basicTypes, fName) {
					if v.verbose {
						fmt.Println(v.functionName + " is making a call that might be unpure: " + fName)
					}
					v.pure[v.functionName] = false
					return nil
				} else {
					//if v.verbose {
					//	fmt.Println("casting to " + fName + " is fine")
					//}
				}
			}
		}
	case *ast.AssignStmt:
		// Gather the new identifiers
		for _, newThing := range t.Lhs {
			//fmt.Println("CREATED " + fmt.Sprintf("%s", newThing))
			*v.created = append(*v.created, fmt.Sprintf("%s", newThing))
		}
	//case *ast.SelectorExpr:
	//	fmt.Println("case *ast.SelectorExpr")
	//	fName := t.Sel.Name
	//	if ok, isPure := v.pure[fName]; ok && isPure {
	//		fmt.Println(v.functionName + " is calling an unproblematic function: " + fName)
	//	} else {
	//		fmt.Println(v.functionName + " is calling a function that might be unpure: " + fName)
	//		v.pure[v.functionName] = false
	//		return nil
	//	}
	//case *ast.ReturnStmt:
	//	fmt.Println("case *ast.ReturnStmt")
	//case *ast.FuncDecl:
	//fmt.Println("case *ast.FuncDecl")
	//case *ast.ExprStmt:
	//	fmt.Println("case *ast.ExprStmt")
	//	if call, ok := t.X.(*ast.CallExpr); ok {
	//		if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
	//			fName := fun.Sel.Name
	//			if ok, isPure := v.pure[fName]; ok && isPure {
	//				fmt.Println(v.functionName + " is calling an unproblematic function: " + fName)
	//			} else {
	//				fmt.Println(v.functionName + " is calling a function that might be unpure: " + fName)
	//				v.pure[v.functionName] = false
	//				return nil
	//			}
	//		}
	//	} else {
	//		fmt.Println("UNHANDLED EXPRESSION STATEMENT", t)
	//		//return nil
	//	}
	default:
		//fmt.Printf("ignored: %T\n", node)
	}
	return v
}

// PureFunctions returns a slice of the names of the functions that are considered pure
func PureFunctions(filename string, verbose bool) []string {
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

	// Purity map per function, starts out as true, but can be set to false
	pure := make(map[string]bool)

	if verbose {
		fmt.Println("Functions:")
	}
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		functionName := fn.Name.Name
		if verbose {
			fmt.Println("--- " + functionName + " ---")
		}

		pure[functionName] = true

		// Examine the function signature. If non-simple types are passed in, they must not be modified in the function body.
		var argNames []string
		for _, arg := range fn.Type.Params.List {
			// output "name: type" or "name, name, ...: type"
			for _, argIdent := range arg.Names {
				argNames = append(argNames, argIdent.String())
			}
			ident, ok := arg.Type.(*ast.Ident)
			if !ok {
				fmt.Printf("%v contains a non-pure type: %T\n", argNames, arg.Type)
				pure[functionName] = false
				break
			}
			typeName := ident.Name
			if !has(basicTypes, typeName) {
				fmt.Println(typeName + ": impure")
				pure[functionName] = false
				break
			}
			if verbose {
				fmt.Printf("pure type: %s\n", arg.Type)
			}
		}
		//fmt.Printf("arg names: %v\n", argNames)

		//if verbose {
		//	fmt.Println("The "+functionName+" signature is pure:", pure[functionName])
		//}

		if pure[functionName] {

			created := []string{}
			// Examine the function body. If non-pure functions are called, the function is not pure.
			for _, stmt := range fn.Body.List {
				// Look for:
				// * function calls to non-pure or unencountered functions (functions must exist in the "pure" map and be marked as pure)
				// * use of globals
				if !pure[functionName] {
					break
				}
				v := NewFuncBodyVisitor(pure, functionName, &created, verbose)
				ast.Walk(v, stmt)
				//fmt.Println(functionName + " identifiers: ")
				for _, name := range v.idents {
					if !has(argNames, name) && !has(basicTypes, name) && !has(created, name) {
						if ok, isPure := pure[name]; ok && isPure {
							//if v.verbose {
							//	fmt.Println("Calling the pure function " + name + " is fine.")
							//}
						} else {
							if v.verbose {
								fmt.Println("not pure: " + name)
							}
							pure[functionName] = false
							break
						}
					}
				}
				if !pure[functionName] {
					break
				}
			}
		}

		fmt.Println("The "+functionName+" function is pure:", pure[functionName])
	}

	pureFunctionNames := []string{}
	for name, isPure := range pure {
		if isPure {
			pureFunctionNames = append(pureFunctionNames, name)
		}
	}
	return pureFunctionNames
}

func main() {
	filename := "data/main.go"
	pureFunctions := PureFunctions(filename, true)
	fmt.Printf("Pure functions in %s:\n", filename)
	for _, name := range pureFunctions {
		fmt.Println("\t" + name)
	}
}
