package main

import (
	"fmt"
    "os"
)

func main() {
    fmt.Printf("hello, world\n")
    fmt.Printf("My name is naimish.\n")
    
    if len(os.Args) != 2 {
			os.Exit(1)
	}
	
	fmt.Println("It's over ",os.Args[1])
}
