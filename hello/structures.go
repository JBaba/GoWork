package main

import(
	"fmt"
)

// declare type
type Person struct{
	Name string
	Age int
}

func main() {

	// declare var by tyoe as struct
	naimish := Person{
		Name: "Naimish",
		Age: 25,
	}

	fmt.Println("Name: ",naimish.Name)
	fmt.Println("Age: ",naimish.Age)
	
}