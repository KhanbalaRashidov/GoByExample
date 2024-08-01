package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{name, age}
}

func main() {

	fmt.Println(person{"Khan", 24})
	fmt.Println(person{name: "Khan2"})
	fmt.Println(&person{name: "Khan3", age: 24})
	fmt.Println(newPerson("Khan4", 24))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}
