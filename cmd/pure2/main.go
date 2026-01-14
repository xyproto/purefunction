package main

import (
	"fmt"
	"os"

	"github.com/xyproto/purefunction"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: pure2 FILE")
		os.Exit(1)
	}

	filename := os.Args[1]

	pureFuncDecls, err := purefunction.Functions(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "%v", err)
		os.Exit(1)
	}
	for _, funcDecl := range pureFuncDecls {
		fmt.Println(funcDecl.Name.Name)
	}
}
