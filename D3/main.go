package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	// demoLoops()
	// demoSwitchCase()
	// demoArrays()

	// demoSlices()
	demoMath()
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
	// var a, b = 1, 2.2
	var a, b = 1, 2

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

func demoArrays() {
	fmt.Println("\n ARRAYS")
	var arr [3]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2.0 // but 2.3 wont work
	// arr[3] = 4 // error

	fmt.Println("Length of arr => ", len(arr))
	fmt.Println("Arr is => ", arr)

	b := []int{1, 2, 3}
	fmt.Println("Length of b => ", len(b))
	for _, item := range b {
		fmt.Println(item)
	}

	array := []int{10, 20, 30}
	array = append(array, 1)
	array = append(array, 2)
	fmt.Printf("Dynamic Array: %v\n", array)
}

func demoSlices() {
	arr := [4]string{
		"Stark",
		"Rogers",
		"Banner",
		"Natasha",
	}
	fmt.Printf("arr after init no slice : %v\n", arr)

	slice1 := arr[1:3]
	fmt.Printf("Slice [1:3] is %v\n", slice1)

	slice2 := arr[:3]
	slice2[0] = "dummy"
	fmt.Printf("Slice [:3] is %v\n", slice2)

	slice3 := arr[1:]
	fmt.Printf("Slice [1:] is %v\n", slice3)

	fmt.Printf("arr after slices : %v\n", arr)

	slice1 = append(slice1, slice2...)
	fmt.Printf("arr after slices : %v\n", slice1)

}

func demoMath() {
	// Randomizers

	source := rand.NewSource(int64(time.Now().UnixNano())) // Create new source/seed
	r := rand.New(source)

	random := r.Intn(20)
	fmt.Println(random)

	// pseudo randomizers
	fmt.Println("default seed", rand.Intn(10), rand.Intn(30))

}
