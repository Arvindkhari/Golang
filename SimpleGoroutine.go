package main

import "fmt"

func display(str string) {
	for w := 0; w < 6; w++ {
		fmt.Println(str)
	}
}

func main() {

	// Calling Goroutine
	go display("Welcome")

	// Calling normal function
	display("This is Goroutine")
}
