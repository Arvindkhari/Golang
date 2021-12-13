package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}

	b := a[1:4] // from one to four
	c := a[:]   // all elements
	d := a[:4]  // from zero to four
	e := a[2:]  // from 2 to the end

	fmt.Println("1-4", b)
	fmt.Println("all", c)
	fmt.Println("0-4", d)
	fmt.Println("2-the end", e)

	b[0] = 777
	fmt.Println(a)

	b = append(b, 81)
	b = append(b, 82)
	for i := 1; i < 100; i++ {
		b = append(b, rand.Intn(500))
	}
	fmt.Println("Slice b after append", b)
	fmt.Println("Array a after appending b ", a)

	//var slice []int
	slice := make([]int, 0)
	fmt.Println("Before appending length:", len(slice), " Cap:", cap(slice))
	//slice[0] = 100
	slice = append(slice, 100)
	for i := 1; i < 100; i++ {
		slice = append(slice, rand.Intn(500))
	}
	fmt.Println("After appending length:", len(slice), " Cap:", cap(slice))

	fmt.Println(slice)
	//  to create slice , it can be created from array or it can be created using make keyword.
	// for slice we dnt need to give the size
	// slice has len and cap
	// slice is ref type
	// slice can grow and shrink
	//  make can be used with slice, maps , chan
}
