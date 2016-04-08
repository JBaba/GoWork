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
	
	// Empty struct
	vivek := Person{}

	fmt.Println("Name: ",vivek.Name)
	fmt.Println("Age: ",vivek.Age)

	// Half init
	akash := Person{
		Name: "Akash",
	}
	akash.Age = 20

	fmt.Println("Name: ",akash.Name)
	fmt.Println("Age: ",akash.Age)

	// Init without "TYPE" name
	xyz := Person {"xyz",78}

	fmt.Println("Name: ",xyz.Name)
	fmt.Println("Age: ",xyz.Age)

	//Pass copy into function
	passValue(naimish)
	fmt.Println(naimish.Age)

	//Pass by value
	abc := &Person{"abc",10}
	passValue2(abc)
	fmt.Println(abc.Age)	
}

func passValue2(s *Person){
	s.Age += 100
}

func passValue(s Person){
	s.Age += 100
}