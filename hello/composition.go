package main

import (
	"fmt"
)

type Person struct{
	Name string
}

func (p *Person) Print(){
	fmt.Println("Name of student:",p.Name)
}

type Student struct{
	*Person
	marks int
}

func main() {
	naimish := &Student{
		Person: &Person{
			Name: "Naimish",
		},
		marks: 20,
	}

	naimish.Print()
	fmt.Println(naimish.Name)
	fmt.Println(naimish.Person.Name)
}

