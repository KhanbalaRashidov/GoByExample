package main

import "fmt"

func main() {

	number := 9

	if number%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	if number%2 != 0 && number%9 == 0 {
		fmt.Println("Logical and operator")
	}

	if number < 0 {
		fmt.Println("Number is negative")
	} else if number == 0 {
		fmt.Println("Number is zero")
	} else {
		fmt.Println("Number is positive")
	}
}
