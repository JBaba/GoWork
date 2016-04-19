package main

import (
	"fmt"
)

func main() {
	lookup := make(map[string]int)
	lookup["name"] = 9001
	power, exists := lookup["veg"]

	fmt.Println(power,exists)
}