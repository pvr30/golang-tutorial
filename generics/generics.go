// This program demonstrates the use of generics in Go.
// Generics allow functions or types to operate on different data types
// while ensuring type safety and code reusability.
// The generic function 'add' accepts a type parameter 'T', which is constrained
// to specific types (int, float64, string) using the type constraints syntax.

package main

import "fmt"

func main() {
	result := add(1, 2) // Calls the generic 'add' function with 'int' type.
	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T {
	// 'add' is a generic function with a type parameter 'T'.
	// The type parameter 'T' is constrained to specific types: 'int', 'float64', and 'string'.
	// This allows the function to operate on these types while maintaining type safety.

	return a + b
	// The '+' operator works for the types listed in the constraint (int, float64, string).
}
