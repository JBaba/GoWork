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

	fmt.Println("Slice ex:")

	scores := []int {1,2,3,4,5}
	slice := scores [2:4]
	fmt.Println(scores)
	fmt.Println(slice)
	slice [0] = 999
	fmt.Println(scores)
	fmt.Println(slice)

	// array of struct pointer  
	students_2 := []*student_1{
	  &student_1{
	   name: "naimish",
	   size: 1,
	  },
	  &student_1{
	   name: "vivek",
	   size: 2,
	  },
	  &student_1{
	   name: "raj",
	   size: 3,
	  },
	}

	extractSize2(students_2)
}

type student struct{
	name string
	size int
}

type student_1 struct{
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

// Func which arg is array of struct pointer
func extractSize2(students []*student_1) []int{
	listofsize := make([]int,len(students))
	for index, aStudrnt := range students {
		listofsize[index] = aStudrnt.size
	}
	return listofsize
}
