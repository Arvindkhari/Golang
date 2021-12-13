package main

import "fmt"

func main() {
	printStack(6)

}

func printStack(num int) {
	fmt.Println("Input number ", num)
	var i, j int
	for i = 0; i < num; i++ {
		for j = i; j < num; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
