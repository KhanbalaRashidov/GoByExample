package main

import "fmt"

func incrementFunc() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {

	increment := incrementFunc()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	newIncrement := incrementFunc()
	fmt.Println(newIncrement())
}
