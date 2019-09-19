package purefunction

import (
	"fmt"
	"log"
	"sort"
)

func ExamplePureFunctions() {
	filename := "test/main.go"
	pureFunctions, err := PureFunctions(filename)
	if err != nil {
		log.Fatalln(err)
	}
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
