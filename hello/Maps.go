package main

import (
	"fmt"
)

func main() {
	lookup := make(map[string]int)
	lookup["name"] = 9001
	power, exists := lookup["veg"]

	fmt.Println(power,exists)

	total := len(lookup)
	fmt.Println("total size of map:",total)
	fmt.Println(lookup)

	lookup["age"] = 81	

	total = len(lookup)
	fmt.Println("total size of map:",total)
	fmt.Println(lookup)

	delete(lookup,"age")
	total = len(lookup)
	fmt.Println("total size of map:",total)
	fmt.Println(lookup)

}