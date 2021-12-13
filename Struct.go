package main

import "fmt"

func main() {
	e1 := Employee{string: "Jiten", int: 12, float32: 12323.45}
	fmt.Println(e1.string, e1.int)
}

type Employee struct {
	string
	int
	float32
}
