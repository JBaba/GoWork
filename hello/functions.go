package main

import (
	"fmt"
)

func power(name string) (int,bool) {
	return 0,false
}

func main() {
	fmt.Println("Its functions.go file")

	// Two return can be saved like this
	value, flag := power("naimish")

	if flag == false {
		fmt.Println("value:",value," flag:",flag)
	}

}