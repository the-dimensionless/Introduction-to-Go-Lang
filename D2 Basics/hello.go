package main // pkg definition

import (
	sayings "basics/sayings" // pkg with alias
	"fmt"                    // pkg import
)

// func definition
func main() {
	fmt.Println("Hello World!") // IO
	sayings.Pithy()
}
