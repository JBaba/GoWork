package main

import (
	"fmt"
)

func main(){
	
	// Diclare variable using types
	var power int
	var flag bool
	var txt string
	
	fmt.Printf("Default value for int %d\n",power)
	fmt.Printf("Default value for booleans %d\n",flag)
	fmt.Printf("Default value for string %d\n",txt)
	
	power = 9
	
	fmt.Printf("It's over %d\n",power)
	
	// Single diclare variable without using type
	x := 50
	
	fmt.Printf("It's over %d\n",x)
	
	name := "Vivek"
	
	// multiple variable declaration
	name, y := "Naimish", 5

	// %s is used to print name(string) 
	// %d is used to print x(int)
	fmt.Printf("%s's age is %d\n",name,y)

	
}
