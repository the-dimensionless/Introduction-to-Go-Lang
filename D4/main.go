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
	dummy()
	fmt.Println("--------- Demo Defer---------")
	demoDefer()
}

func calculator(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}

func dummy() {
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
	fmt.Println(" slice => ", slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 50)
	slice1 = append(slice1, 60)
	slice1 = append(slice1, 70)
	fmt.Println("Extended slice under cap => ", slice1, len(slice1), cap(slice1))

	// newCap =  oldCap + ceil[length / 2]

}

func demoDefer() {
	// delay execution till near by functions are completed
	// can be used at statement/function/method level
	defer fmt.Println("Defer 1")
	fmt.Println("No defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("No defer 2")
	dummy1()
	defer fmt.Println("defer 3")
	fmt.Println("No defer 3")
}

func dummy1() {
	fmt.Println("Dummy 1")
}

func dummy2() {
	fmt.Println("Dummy 2")
}

func dummy3() {
	fmt.Println("Dummy 3")
}
