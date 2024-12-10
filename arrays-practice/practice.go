package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Sport", "Cooking", "Reading"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)

	// creating list for product struct
	products := []Product{
		{
			1,
			"First-Product",
			10.99,
		},
		{
			2,
			"Second-Product",
			20.12,
		},
	}
	fmt.Println(products)

	newProduct := Product{
		3,
		"Third-Product",
		100.0,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}
