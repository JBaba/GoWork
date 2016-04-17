package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "the spice must flow"
	num := strings.Index(text[5:]," ")
	fmt.Println(num)
	// number to the ends
	fmt.Println(text[5:])
	// to the number text
	fmt.Println(text[:5])
	fmt.Println(strings.Index(text[5:],"p"))

	score := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(score[:len(score)-1])

	removeAtIndex(score,2)
}

func removeAtIndex(source []int,index int) []int {
	lastIndex := len(source) - 1
	// swap the last value and the value we want to remove
	fmt.Println(source)
	source[index],source[lastIndex] = source[lastIndex],source[index]
	fmt.Println(source)

	return source[:lastIndex]
}

