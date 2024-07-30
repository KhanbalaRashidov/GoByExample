package main

import (
	"fmt"
	"math"
)

const (
	firstName string = "test"
	age       int    = 1
)

func main() {

	fmt.Println(firstName, age)

	const PI float64 = 3.14
	fmt.Println(PI)

	const binaryNumber = 2e20 / PI
	fmt.Println(binaryNumber)

	fmt.Println(math.Cos(binaryNumber))
}
