package main

import (
	"fmt"
)

func main() {
	// use timeout to refactor the code
	// for {
	// }
}

func demoFunctions() {
	// return types & params  => optional
	// sum, diff := calculator(5, 3)
	// fmt.Printf("[%d, %d]: Sum is %d and diff is %d", 5, 3, sum, diff)
	// dummy()
	//fmt.Println("--------- Demo Defer---------")
	//demoDefer()
	// demoPts()
	demoStructs()
}

func calculator(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2

	defer fmt.Println("Done")
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
	// defer with function or method call
	/* defer fmt.Println("Defer 1")
	fmt.Println("No defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("No defer 2")
	dummy1()
	defer fmt.Println("defer 3")
	fmt.Println("No defer 3") */

	fmt.Println("counting")
	i := 0
	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }
	for {
		defer fmt.Println(i)
		i++
		if i >= 10 {
			break
		}
	}
	fmt.Println("done")
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

func demoPts() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer; deferencing or indirecting
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

}

func demoStructs() {
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})

	emp1 := Employee{
		name:    "Sumit Singh",
		age:     24,
		address: "10880 MBP",
	}
	emp2 := Employee{
		name:    "Sumit Pandit",
		age:     14,
		address: "10881 MBP",
	}
	emp3 := Employee{
		name:    "Sumit Sharma",
		age:     29,
		address: "10882 MBP",
	}
	var emp4 Employee
	emp4.name = "Mr. Bean"
	emp4.age = 65
	emp4.address = "London, UK"
	emp4.print()

	emp5 := Employee{"Chris Evans", 36, "Downtown Seattle"}
	emp5.print()

	fmt.Println("---------Emps----------")
	emp1.print()
	emp2.print()
	emp3.print()

	print(emp1)

}

type Employee struct {
	name    string
	age     int
	address string `default:"Home Sweet Home"`
}

func (e Employee) print() {
	fmt.Printf("Emp Name : %v, Age: %v, Address: %v\n", e.name, e.age, e.address)
}

func print(e Employee) {
	fmt.Printf("Emp Name : %v, Age: %v, Address: %v", e.name, e.age, e.address)
}
