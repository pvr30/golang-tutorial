package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}
	sum := sumup(1, 10, 15, 40, -5)
	fmt.Println(sum)

	anotherSum := sumup(1, numbers...)
	fmt.Println(anotherSum)
}

// `sumup` is a **variadic function**.
// It can accept a variable number of arguments for the `numbers` parameter.

func sumup(startingValue int, numbers ...int) int {

	fmt.Println(startingValue)

	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
