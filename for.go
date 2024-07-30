package main

import "fmt"

func main() {

	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}

	for n := range 10 {
		fmt.Println("range", n)
	}

	//infinity loop
	i = 1
	for {
		if i > 10 {
			fmt.Println("break")
			break
		}
		i++
	}

	for j := range 10 {
		if j%2 != 0 {
			continue
		}
		fmt.Println(j)
	}
}
