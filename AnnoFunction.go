package main

import "fmt"

func main() {

	// without argument
	func() {
		fmt.Println("Hello World by anon")
	}()

	// with string as arg

	func(str string) {
		fmt.Println(str)
	}("Hello Folks")

	// anon func with return
	square := func(n int) int {
		return n * n
	}(100)
	fmt.Println("Square", square)

	// anon with multiple returns

	a, p := func(l, b float32) (float32, float32) {
		return l * b, 2 * (l + b)
	}(12.5, 15.5)
	fmt.Println("Area Of Rect and Perimeter of Rect", a, p)

	execute(func() {
		fmt.Println("Hello World,Execute the function with func parameter")
	})

	execute(display)
}

func display() {
	fmt.Println("Calling display function")
}

func execute(f func()) {
	f()
}
