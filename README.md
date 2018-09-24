# purefunction

# WORK IN PROGRESS

Given Go code, find all pure functions.

Pure functions, like the `fibonacci` function, has great potential for optimization by memoization.

A "pure function" for this module is, a function that:

* Only calls any non-pure functions.
* Does not read or write to any global variables.
* Does not modify any of the argument variables.
* Does not read or write to any memory location using pointers.
* Always returns the same answer, given the same input.

Example of a pure function:

    func add(a, b int) int {
        return a + b
    }

If this module believes that a function is pure, it must be correct, but if in doubt, it should mark a function as non-pure. False negatives are okay, but not false positives.
