# purefunction

Given a Go source file, find the names of all functions that are known to be pure.

Pure functions, like the `fibonacci` function, has great potential for optimization by memoization.

A "pure function" for this module is, a function that:

* Only calls functions that are known to be pure, if any.
* Does not read or write to any global variables.
* Does not have pointers of slices as function arguments.
* Does not read or write to any memory location using pointers.
* Ideally: Always returns the same answer, given the same input, but this is hard to test for (ref. halting problem).

Example of a pure function:

    func add(a, b int) int {
        return a + b
    }

Functions are filtered out if they have non-pure indicators. The ones that are left are considered pure. There may be false negatives, but not false positives.
