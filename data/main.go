package main

import "fmt"

// a pure function, does not deal with global variables, calls only pure functions (if any)
func add(x, y int) int {
	return x + y
}

// a pure function, does not deal with global variables, calls only pure functions (if any)
func mul(x, y int) int {
	s := 0
	for i := 0; i < x; i++ {
		s = add(s, y)
	}
	return s
}

// not a pure function, calls a function that is not known to be pure: fmt.Println
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(mul(2, 3))
}
