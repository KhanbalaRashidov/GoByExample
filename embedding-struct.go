package main

import "fmt"

type base struct {
	firstname string
}

func (b *base) PrintName() string {
	return fmt.Sprintf("%s", b.firstname)
}

type container struct {
	base
	lastname string
}

func main() {
	name := container{base: base{
		firstname: "Khanbala",
	},
		lastname: "Rashidov",
	}

	fmt.Printf("fistname: %s , lastname: %s", name.firstname, name.lastname)
	fmt.Println(name.base.firstname)
	fmt.Println("printName:", name.PrintName())

	type printer interface {
		PrintName() string
	}

	var p printer = &name
	fmt.Println(p.PrintName())
}
