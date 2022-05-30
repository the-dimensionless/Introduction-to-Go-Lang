package main

import "fmt"

func main() {
	// use timeout to refactor the code
	// for {
	// }
	demoFunctions()
}

func demoFunctions() {
	// return types & params  => optional
	sum, diff := calculator(5, 3)
	fmt.Printf("[%d, %d]: Sum is %d and diff is %d", 5, 3, sum, diff)
}

func calculator(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}
