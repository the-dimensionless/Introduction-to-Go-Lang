package main

import (
	"fmt"
	"time"
)

func main() {
	task1()
	task2()
}

func task1() {
	//write a program to break the below loop by using timeout?
	start := time.Now()         // get current TS
	duration := time.Second * 5 // after how much time shall we stop the execution
	fmt.Println("TS: Loop Start => ", time.Now())
	for {
		if time.Since(start) > duration { // checking difference of TS....can be further optimized
			break
		}
	}
	fmt.Println("TS: Loop End => ", time.Now())
}

func task2() {
	//Operations using slice methods
	var s []int = []int{1, 2, 3}
	fmt.Println("\n", s, len(s), cap(s))
	s = append(s, 3)
	fmt.Println(s, len(s), cap(s))

	//print
	fmt.Println("Print => ", s)

	//slice
	first := s[0]
	slice1 := s[:1]
	fmt.Println("First element ", first)
	fmt.Println("slice => ", slice1, len(slice1), cap(slice1))

	//extend the length of slice
	slice1 = append(slice1, 10)
	slice1 = append(slice1, 20)
	slice1 = append(slice1, 30)
	fmt.Println("Extended slice under cap => ", slice1, len(slice1), cap(slice1))

	//drop first 2 elements from slice
	slice1 = slice1[2:]
	fmt.Println("slice => ", slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 50)
	slice1 = append(slice1, 60)
	slice1 = append(slice1, 70)
	fmt.Println("Extended slice under cap => ", slice1, len(slice1), cap(slice1))

	// Dynamic growth rate
	// newCapacity =  oldCapacity + ceil[length / 2]
}
