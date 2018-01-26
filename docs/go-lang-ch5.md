# Functions

## Function Declarations
1. If the function returns one unnamed result or no results at all, parentheses are optional and usually omitted.
2. Like parameters, results may be named. In that case, each name declares a local variable ini- tialized to the zero value for its type.
3. A sequence of parameters or results of the same type can be factored so that the type itself is written only once.
4. The type of a function is sometimes called its signature. Two functions have the same type or signature if they have the same sequence of parameter types and the same sequence of result types.
5. Go has no concept of default parameter values, nor any way to specify arguments by name, so the names of parameters and results don’t matter to the caller except as documentation.
6. Parameters are local variables within the body of the function, with their initial values set to the arguments supplied by the caller. Function parameters and named results are variables in the same lexical block as the function’s outermost local variables.
7. Arguments are passed by value, so the function receives a copy of each argument; modifica- tions to the copy do not affect the caller. 
8. You may occasionally encounter a function declaration without a body, indicating that the function is implemented in a language other than Go. Such a declaration defines the function signature.

## Recursion
1. typical Go implementations use variable-size stacks that start small and grow as needed up to a limit on the order of a gigabyte. This lets us use recursion safely and without worrying about overflow

## Multiple Return Values
1. Go’s garbage collector recycles unused memory, but do not assume it will release unused operating system resources like open files and network connections. They should be closed explicitly.
2. In a function with named results, the operands of a return statement may be omitted. This is called a bare return. A bare return is a shorthand way to return each of the named result variables in order.
3. Bare returns can reduce code duplication, but they rarely make code easier to understand. For this reason, bare returns are best used sparingly.

## Errors
1. A function for which failure is an expected behavior returns an additional result, convention- ally the last one. If the failure has only one possible cause, the result is a `boolean`, usually called ok. More often, and especially for I/O, the failure may have a variety of causes for which the caller will need an explanation. In such cases, the type of the additional result is `error`.
2. The built-in type `error` is an interface type. An error may be nil or non-nil, that nil implies success and non-nil implies failure, and that a non-nil error has an error message string which we can obtain by calling its `Error` method.
3. Exceptions tend to entangle the description of an error with the control flow required to handle it, often leading to an undesirable outcome: routine errors are reported to the end user in the form of an incomprehensible stack trace, full of information about the structure of the program but lacking intelligible context about what went wrong. By contrast, Go programs use ordinary control-flow mechanisms like if and return to respond to errors. This style undeniably demands that more attention be paid to error-han- dling logic, but that is precisely the point.
4. We should build descriptive errors by successively prefixing additional context information to the original error message. Because error messages are frequently chained together, message strings should not be capitalized and newlines should be avoided.
5. In general, the call f(x) is responsible for reporting the attempted operation f and the argu- ment value x as they relate to the context of the error. The caller is responsible for adding fur- ther information that it has but the call f(x) does not.
6. For errors that represent transient or unpredictable problems, it may make sense to retry the failed operation, possibly with a delay between tries, and perhaps with a limit on the number of attempts or the time spent trying before giving up entirely.
7. Functions tend to exhibit a common structure, with a series of initial checks to reject errors, followed by the substance of the function at the end, minimally indented.

## Function Values
1. Functions are first-class values in Go: like other values, function values have types, and they may be assigned to variables or passed to or returned from functions.
2. The zero value of a function type is nil.
3. Function values may be compared with nil, but they are not comparable.
4. 