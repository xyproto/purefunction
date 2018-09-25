package purefunction

import (
	"fmt"
	"sort"
)

func ExamplePureFunctions() {
	filename := "test/main.go"
	pureFunctions := PureFunctions(filename, false)
	fmt.Printf("Pure functions in %s:\n", filename)
	sort.Strings(pureFunctions)
	for _, name := range pureFunctions {
		fmt.Println(name)
	}
	// Output: Pure functions in test/main.go:
	// add
	// add2
	// mul
	// mul3
}
