package main

import "fmt"

type rect struct {
	width, height int
}

// reference type
func (r *rect) area() int {
	return r.width * r.height
}

// value type
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

}
