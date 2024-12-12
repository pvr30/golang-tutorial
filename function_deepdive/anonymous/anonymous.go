package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	// anonymous function
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// `createTransformer` is a **factory function**.
// It generates and returns a closure based on the given factor.
// A factory function creates and returns specific instances of functions or objects,
// in this case, a transformation function.
func createTransformer(factor int) func(int) int {
	// This is a **closure**: it captures and remembers the `factor` variable
	// from its surrounding environment even after `createTransformer` returns.
	return func(number int) int {
		return number * factor // `factor` is enclosed within this function
	}
}
