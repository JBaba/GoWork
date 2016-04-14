package main

import (
	"fmt"
)

func main() {
	var score [10]int
	score[0] = 10
	fmt.Println(score)

	score_1 := [4]int{10,22,31,45}
	fmt.Println(score_1)
	// Use of len() func
	fmt.Println("Length of socre_1:",len(score_1))

	//for loop
	for i := 0; i < len(score_1); i++ {
		fmt.Println(score_1[i])
	}

	//for loop
	for index, value := range score_1 {
		fmt.Println(index," : ",value)
	}
}

