// Go program to illustrate the
// concept of Expression switch
// statement
package main

import "fmt"

func main() {
	var value int = 2

	// Switch statement without an
	// optional statement and
	// expression
	switch {
	case value == 1:
		fmt.Println("Hello")
	case value == 2:
		fmt.Println("Bonjour")
	case value == 3:
		fmt.Println("Namstay")
	default:
		fmt.Println("Invalid")
	}

}
