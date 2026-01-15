package fib

import "sync"

// --- not memoized ---

func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// --- memoized ---

var fibCache sync.Map

func FibMemoized(n int) int {
	if n <= 1 {
		return 1
	}
	if result, ok := fibCache.Load(n); ok {
		return result.(int)
	}
	result := FibMemoized(n-1) + FibMemoized(n-2)
	fibCache.Store(n, result)
	return result
}
