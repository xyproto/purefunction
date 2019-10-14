package fib

// --- not memoizied ---

func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// --- memoizied ---

type fibType map[int]int

var fibCache fibType

func FibMemoized(n int) int {
	if fibCache == nil {
		fibCache = make(fibType)
	}
	if result, ok := fibCache[n]; ok {
		return result
	}
	if n <= 1 {
		fibCache[n] = 1
		return fibCache[n]
	}
	fibCache[n] = FibMemoized(n-1) + FibMemoized(n-2)
	return fibCache[n]
}
