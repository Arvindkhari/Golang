package main

import "fmt"

func main() {
	// var a int
	// var arr [3]int
	// var slice []int
	// var mymap map[string]int

	var ptr *int

	b := 255

	ptr = &b

	fmt.Println("Value of B", b)

	fmt.Println("Address of B", &b)

	fmt.Println("Address of B", ptr)

	fmt.Println("Value of Of through ptr", *ptr)

	*ptr = 256

	fmt.Println("Value of B", b)

	var ptr2 **int
	ptr2 = &ptr
	fmt.Println("Address that ptr2 holds:", ptr2)
	if ptr2 != nil {
		fmt.Println("dereference of ptr2", *ptr2)
	}
	**ptr2 = 259
	fmt.Println("Value of B", b)

}
