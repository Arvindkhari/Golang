package main

import (
	"fmt"
)

func main() {
	a := 99
	changeValue(a)
	fmt.Println("Has a chaged?", a)
}

func changeValue(v int) {
	fmt.Println("The value of V", v)
	v = 100
	fmt.Println("The value of V after mutating", v)
}
