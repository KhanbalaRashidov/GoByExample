package main

import "fmt"

func mathOperation(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	a, b := 3, 33
	fmt.Println(a, b)

	fmt.Println(mathOperation(a, b))
}
