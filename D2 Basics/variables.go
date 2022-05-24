package main

import "fmt"

func demoVars() {
	var x int
	x = 3
	fmt.Println("X: ", x)

	var y int = 15 // static definition of type
	fmt.Println("Y: ", y)

	const z = 30 // dynamica detection of type
	fmt.Println("Z: ", z)

	// multiple variables
	var i, k = 200, "Hello"
	fmt.Printf("Value of i %d and k is %s", i, k)
}
