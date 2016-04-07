package main

import (
	"fmt"
)

// function return multipale values
func power(name string) (int,bool) {
	return 0,false
}

// functions argument same type
func add(a,b int) int {
	sum := a+b
	return sum
}

func main() {
	fmt.Println("Its functions.go file")

	flag := true

	// Two return can be saved like this
	value, flag := power("naimish")

	if flag == false {
		fmt.Println("value:",value," flag:",flag)
	}

	// we only care about one value then use '_'
	_, flag2 := power("vivek")

	if flag2 == false {
		fmt.Println("value:",value," flag:",flag2)
	}

	fmt.Println("Sum:",add(1,2))

}