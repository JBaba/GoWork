package main

import (
	"fmt"
)

// declare type
type Person struct{
	Name string
	Age int
}

// Fields of a Structure
type Student struct{
	Name string
	Age int
	Classmate *Student
}

//associate method with structure
func (s *Person) ageIncr(){
	s.Age += 10
}

func main() {
	naimish := &Person{"naimish",25}
	naimish.ageIncr()
	fmt.Println(naimish.Age)
	fmt.Println(naimish)

	// Use of Constructor
	vivek := NewPerson("vivek",24)
	fmt.Println(vivek)

	// use of new keywoird
	james := new(Person)
	fmt.Println(james)
	james.Name = "Jmaes"
	james.Age = 50
	fmt.Println(james)

	// Struct as a fields declaration
	naimish_1 := &Student{
		Name: "Naimish",
		Age: 25,
		Classmate: &Student{
			Name: "vivek",
			Age: 24,
			Classmate: nil,
		},
	}

	fmt.Println(naimish_1)
}

// Constructor sends copy
func NewPerson(Name string,Age int) Person{
	return Person{
		Name: Name,
		Age: Age,
	}
}





