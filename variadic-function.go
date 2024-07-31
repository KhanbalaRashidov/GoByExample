package main

import "fmt"

func hello(text ...string) {
	fmt.Println(text)
	return
}

func main() {

	hello("Hello", "world", "gopher")
}
