package main

import (
	"fmt"
)

func change(ptr *int, v int) {
	*ptr = v // the value that the pointer points to ..
	*ptr++
}

func main() {
	val := 100

	fmt.Println("Value=", val)

	change(&val, 101)

	fmt.Println("Value after calling the func with pointer", val)
}
