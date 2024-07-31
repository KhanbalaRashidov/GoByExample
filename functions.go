package main

import "fmt"

func minus(a, b int) int {
	return a - b
}

func plus(a, b int) int {
	return a + b
}

func multiplication(a, b int) int {
	return a * b
}

func main() {
	result := plus(1, 2)
	fmt.Println(result)

	result = minus(1, 2)
	fmt.Println(result)

	result = multiplication(1, 2)
	fmt.Println(result)
}
