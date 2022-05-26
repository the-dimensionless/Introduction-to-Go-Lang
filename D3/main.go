package main

import "fmt"

func demoVars() {
	a := 10
	fmt.Println(a)

	//a := 20 // Error => NO NEW VARIABLES ON LEFT SIDE OF := (eat op)
	// a := "a" // ERROR => no new variables on left side of :=, Untyped Float assigned

	b := 20.0
	fmt.Println(b)
	// b := 5 // no new variables on left side of :=
	fmt.Println(b)

	const c = 5
	// const d := 5 // Error
	fmt.Println("Constant ", c)
	// c = 3 // Error cannot assign to c

	const d int = 50 // Error => missing constant value
}

func main() {
	demoLoops()
	demoSwitchCase()
}

func demoLoops() {
	// traditional loop
	for i := 0; i < 10; i++ {
		fmt.Println(i, " ", &i)
	}

	s := []string{"A", "B", "C"}
	// for range loop
	for _, item := range s {
		fmt.Println(item)
	}

	n := 1
	for n < 5 {
		fmt.Println("Less than 5")
		n++
	}

	// range loop over string
	for range "Hello" {
		fmt.Println("Yellow")
	}

	// infinte loop
	for {
		fmt.Println("Infinity")
		if n == 10 {
			break
		}
		n++
	}

}

func demoSwitchCase() {
	a, b := 1, 2

	switch a + b {
	case 1:
		fmt.Println("First case")

	case 2:
		fmt.Println("Second case")

	case 3:
		fmt.Println("Third case")
	default:
		fmt.Println("Default case")
	}
}
