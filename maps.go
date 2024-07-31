package main

import (
	"fmt"
	"maps"
)

func main() {
	numbers := make(map[string]int)

	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3

	fmt.Println("numbers:", numbers)

	one := numbers["one"]
	fmt.Println("one:", one)

	two := numbers["two"]
	fmt.Println("two:", two)

	four := numbers["four"]
	fmt.Println("four:", four)

	delete(numbers, "three")
	fmt.Println("numbers:", numbers)

	//clear maps
	clear(numbers)
	fmt.Println("numbers:", numbers)

	numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("numbers:", numbers)

	numbers2 := map[string]int{"one": 1, "two": 2, "three": 3}
	if maps.Equal(numbers, numbers2) {
		fmt.Println("maps are equal")
	}

}
