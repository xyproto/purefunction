# purefunction [![GoDoc](https://godoc.org/github.com/xyproto/purefunction?status.svg)](http://godoc.org/github.com/xyproto/purefunction) [![License](http://img.shields.io/badge/license-BSD-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/purefunction/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/purefunction)](https://goreportcard.com/report/github.com/xyproto/purefunction)

Given a Go source file, find the names of all functions that are known to be pure.

A "pure function" for this module is, a function that:

* Only calls functions that are known to be pure, if any.
* Does not read or write to any global variables.
* Does not have pointers or slices as function arguments.
* Does not read or write to any memory location using pointers.
* Ideally: Always returns the same answer, given the same input, but this is hard to test for (ref. halting problem).

### Examples of pure functions

Example of a pure function:

```go
func add(a, b int) int {
    return a + b
}
```

Another example of a pure function (even though it is a "method"):

```go
func (c *Config) Add(a, b int) int {
    return a + b
}
```

### Features and limitations

* Functions are filtered out if they have non-pure indicators. The ones that are left are considered pure.
* Functions that read from a file, read from the keyboard, uses randomness or the system time are not pure, but may be detected as such.

### Approach

* Uses [`go/ast`](http://golang.org/pkg/go/ast) extensively.

### Memoization

Pure functions, like the `fibonacci` function, has great potential for optimization by memoization.

Benchmark output for "fibonacci" vs "memoized fibonacci":

```
goos: linux
goarch: amd64
pkg: github.com/xyproto/purefunction/example
BenchmarkFib10-8                   30586             38894 ns/op
BenchmarkFibMemoized-8          84117813                14.2 ns/op
PASS
ok      github.com/xyproto/purefunction/example 2.798s
```

See issue #1.

### Requirements

* Go 1.10 or later.

### General info

* Version: 1.0.2
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
