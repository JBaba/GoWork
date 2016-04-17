package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "the spice must flow"
	num := strings.Index(text[5:]," ")
	fmt.Println(num)
	fmt.Println(text[5:])
	fmt.Println(strings.Index(text[5:],"p"))
}

