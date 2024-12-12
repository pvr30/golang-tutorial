package functionsarevalues

import (
	"fmt"
)

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	numbers2 := []int{5, 1, 2}

	doubleNumbers := transformNumbers(&numbers, double)
	tripleNumbers := transformNumbers(&numbers, triple)

	fmt.Println(doubleNumbers)
	fmt.Println(tripleNumbers)

	transformerFunc1 := getTransformerFuntion(&numbers)
	transformerFunc2 := getTransformerFuntion(&numbers2)

	transformNumbers1 := transformNumbers(&numbers, transformerFunc1)
	transformNumbers2 := transformNumbers(&numbers2, transformerFunc2)

	fmt.Println(transformNumbers1)
	fmt.Println(transformNumbers2)

}

// Take function as an argument
func transformNumbers(numbers *[]int, transform transformFunc) []int {
	newNumbers := []int{}

	for _, val := range *numbers {
		newNumbers = append(newNumbers, transform(val))
	}

	return newNumbers
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}

// return function as value
func getTransformerFuntion(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}
