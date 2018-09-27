# purefunction [![Build Status](https://travis-ci.org/xyproto/purefunction.svg?branch=master)](https://travis-ci.org/xyproto/purefunction) [![GoDoc](https://godoc.org/github.com/xyproto/purefunction?status.svg)](http://godoc.org/github.com/xyproto/purefunction) [![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/xyproto/purefunction/master/LICENSE) [![Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat)](http://goreportcard.com/report/xyproto/purefunction)


Given a Go source file, find the names of all functions that are known to be pure.

Pure functions, like the `fibonacci` function, has great potential for optimization by memoization.

Uses [`go/ast`](http://golang.org/pkg/go/ast) extensively.

A "pure function" for this module is, a function that:

* Only calls functions that are known to be pure, if any.
* Does not read or write to any global variables.
* Does not have pointers of slices as function arguments.
* Does not read or write to any memory location using pointers.
* Ideally: Always returns the same answer, given the same input, but this is hard to test for (ref. halting problem).

Example of a pure function:

```go
func add(a, b int) int {
    return a + b
}
```

Functions are filtered out if they have non-pure indicators. The ones that are left are considered pure. There may be false negatives, but not false positives.

### General info

* Version: 1.0.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;

