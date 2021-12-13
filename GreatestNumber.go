package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3
	if a > b {
		if a > c {
			fmt.Println("Biggest is a")
		} else if b > c {
			fmt.Println("Biggest is b")
		}
	} else if b > c {
		fmt.Println("Biggest is b")
	} else {
		fmt.Println("Biggest is c")
	}
}
