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

	//slice with length of 5  capacity 5
	score_2 := make([]int,5)
	fmt.Println(score_2)

	//slice with length of 0 capacity 5
	score_3 := make([]int,0,5)
	fmt.Println(score_3)

	// this will crash due to size zero
	//score_3[0] = 9033
	score_3 = append(score_3,5)
	fmt.Println(score_3)

	scores_4 := make([]int, 0, 10)
	scores_4 = scores_4[0:6]
	scores_4[5] = 9033
	fmt.Println(scores_4)

	//dynemically expanding array
	score_5 := make([]int,0,5)
	c := cap(score_5)
	fmt.Println("cap:",c)
	fmt.Println(score_5)

	for i := 0; i < 7; i++ {
		score_5 = append(score_5,i)

		if cap(score_5) != c {
			c = cap(score_5)
			fmt.Println("cap:",c)
			fmt.Println(score_5)
		}
	}

	//append
	score_6 := make([]int,5)
	score_6 = append(score_6,9332)
	fmt.Println(score_6)

	// slice of struct array
	students := []student{
	  student{
	   name: "naimish",
	   size: 1,
	  },
	  student{
	   name: "vivek",
	   size: 2,
	  },
	  student{
	   name: "raj",
	   size: 3,
	  },
	}

	fmt.Println(extractSize(students))
}

type student struct{
	name string
	size int
}

func extractSize(students []student) []int{
	listofsize := make([]int,len(students))
	for index, aStudrnt := range students {
		listofsize[index] = aStudrnt.size
	}
	return listofsize
}
