package main

import "fmt"

func main() {
	age := 23

	agePointer := &age

	fmt.Println("Age", age)
	fmt.Println("Age Pointer", agePointer)

	// value from pointer
	fmt.Println("Value from age pointer", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) int {
	/* so here this is the first use case of pointer where we pass the pointer in
	 function it won't create a copy of that perticular variable but instead it
	use it's address and use that varible itself */

	return *age - 18
}

func editAgeToAdultYears(age *int) {
	// Using pointer for data mutation
	*age = *age - 10

	// fmt.Scan() is also internally mutate the value by taking pointer.
}
