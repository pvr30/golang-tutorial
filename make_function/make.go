package main

import "fmt"

// type aliases
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// Creates a slice of strings with an initial length of 2 and capacity of 5.
	// Length determines the number of elements initially accessible, and capacity determines how many elements the slice can hold before resizing.
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	// Appends new elements to the slice, increasing its length.
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	// Creates a map with string keys and float64 values.
	// The second argument (3) provides a capacity hint for performance optimization.
	courseRatings := make(floatMap, 3)

	// Adds key-value pairs to the map.
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	// fmt.Println(courseRatings)
	courseRatings.output()

	// for loop on slice and maps

	for index, value := range userNames {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

	fmt.Println("")

	for key, value := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}
}
