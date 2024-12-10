package main

import (
	"fmt"
)

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.99)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...) // unpacking list value by using ...
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "A Carpet"
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:3]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)

// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]

// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
