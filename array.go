package main

import "fmt"

func main() {

	var numbers [10]int
	fmt.Println("numbers:", numbers)

	numbers[3] = 303
	fmt.Println("numbers:", numbers)
	fmt.Println("numbers[3]:", numbers[3])

	numbers2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("numbers2:", numbers2)

	numbers2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("numbers2:", numbers2)

	numbers2 = [...]int{1, 3: 7, 5}
	fmt.Println("numbers2:", numbers2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)

	twoD = [2][3]int{{1, 2, 3}, {3, 4, 5}}
	fmt.Println("twoD:", twoD)
}
