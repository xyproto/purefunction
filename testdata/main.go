package main

import "fmt"

// pure
func add(x, y int) int {
	return x + y
}

// pure
func mul(x, y int) int {
	s := 0
	for i := 0; i < x; i++ {
		s = add(s, y)
	}
	return s
}

// pure
func mul3(a int, b, c float64) float64 {
	return float64(a) * b * c
}

// HappyInt is a basic type
type HappyInt int

// pure, uses a custom type that is a basic type
func add2(a, b HappyInt) HappyInt {
	return a + b
}

// not pure, uses pointer and slice arguments
func add3(a int, b *int, c []int) int {
	if len(c) == 0 {
		return 0
	}
	return a + *b + c[0]
}

var globalInt = 42

// not pure, uses a global variable
func add4(a int) int {
	retval := a + globalInt
	globalInt *= 2
	return retval
}

// UnhappyInt is a pointer to a basic type
type UnhappyInt *uint8

// not pure, uses *uint8
func add5(a UnhappyInt) HappyInt {
	return HappyInt(*a) + 2
}

// not pure, calls nonpure functions: fmt.Println, add3, add4
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(mul(2, 3))
	fmt.Println(mul3(2, 3, 4))
	fmt.Println(add2(HappyInt(2), HappyInt(3)))
	x := 3
	y := []int{4}
	fmt.Println(add3(2, &x, y))
	fmt.Println(add4(7))
	fmt.Println(add4(7)) // add4 uses a global variable, calling it twice is on purpose
	z := uint8(40)
	fmt.Println(add5(&z))
}
