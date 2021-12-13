package main

import (
	"fmt"
)

func main() {
	a, b := 10, 30
	fmt.Println("Before Swap process")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	a, b = b, a
	fmt.Println("After Swap process")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

}
