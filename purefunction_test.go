package purefunction

import (
	"fmt"
	"log"
	"sort"
)

func ExamplePureFunctions() {
	filename := "testdata/main.go"
	pureFunctions, err := PureFunctions(filename)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Pure functions in %s:\n", filename)
	sort.Strings(pureFunctions)
	for _, name := range pureFunctions {
		fmt.Println(name)
	}
	// Output: Pure functions in testdata/main.go:
	// add
	// add2
	// mul
	// mul3
}
