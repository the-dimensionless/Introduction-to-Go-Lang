### Day 4

1. WAP to break the infinite for loop by using timeout?
2. Operations using slice methods
    - define a main() func
    - define an array int[]
    a. print the array
    - create a slice : any range
    b. slice the array : return the 1st element of the slice
    c. extend the length of slice
    d. drop the first 2 elements from the slice
    e. print the slice

3. Read about pointers

Go has pointers. A pointer holds the memory address of a value.
The type *T is a pointer to a T value. Its zero value is nil.

> var p *int

The & operator generates a pointer to its operand.

> i := 42
> p = &i

* Unlike C, Go has no pointer arithmetic.
* It is useful when using function with receivers also
* It helps in reflecting changes to original value even when using call by value 
