package main

import "fmt"

func numNormal(num int) {
	num = 3
}

func numPointer(num *int) {
	*num = 3
}

func main() {

	num := 1
	fmt.Println("num:", num)

	numNormal(num)
	fmt.Println("num:", num)

	numPointer(&num)
	fmt.Println("num:", &num)
}
