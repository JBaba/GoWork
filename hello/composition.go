package main

import (
	"fmt"
)

type Person struct{
	Name string
}

type Person2 struct{
	Name string
}

func (p *Person) Print(){
	fmt.Println("Name of student:",p.Name)
}

type Student struct{
	*Person
	marks int
}

type Student2 struct{
	*Person2
	Name string
	marks int
}

func main() {
	naimish := &Student{
		Person: &Person{
			Name: "Naimish",
		},
		marks: 20,
	}

	vivek := &Student2{
		Person2: &Person2{
			Name: "vivek",
		},
		Name: "Student name is vivek",
		marks: 50,

	}
	naimish.Print()
	fmt.Println(naimish.Name)
	fmt.Println(naimish.Person.Name)


	fmt.Println(vivek.Name)
	fmt.Println(vivek.Person2.Name)
}

