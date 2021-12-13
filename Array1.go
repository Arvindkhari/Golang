package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}
func main() {
	num := [...]int{5, 6, 7, 8, 8}

	fmt.Println("before passing to function ", num)

	changeLocal(num) //num is passed by value

	fmt.Println("after passing to function ", num)
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	fmt.Println(sum)

	sum = 0
	for index, value := range num {
		sum += value
		fmt.Println("Index:", index, " Value:", value)
	}
	fmt.Println("Sum using range loop", sum)

	// range can also be used for strings . Because string is an array

	str := "Hello, 世界"
	//str := "Hello World"
	// range in string always returns unicode / rune
	for _, value := range str {
		fmt.Printf("Type of Value: %T\n", value)
		fmt.Println(value)
	}
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
		fmt.Printf("Type of str of iteration: %T\n", str[i])
	}
}
