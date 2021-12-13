package main

import (
	"fmt"
	_ "time"
)

func display(str string) {
	fmt.Println(str)
}

func square(n int) int { // unnamed return type
	return n * n
}

func cube(n int) (i int) { // named return type
	i = square(n) * n
	return i
}

func areaOfRect(l float32, b float32) (area float32) {
	area = l * b
	return area
}

func perimeterOfRect(l, b float32) float32 {
	peri := 2 * (l + b)
	return peri
}

func areaAndperiOfRect(l, b float32) (area, peri float32) {
	area = areaOfRect(l, b)
	peri = perimeterOfRect(l, b)
	return area, peri
}

func main() {
	display("Hello World")
	a := areaOfRect(12.4, 15.5)
	p := perimeterOfRect(12.4, 15.5)
	fmt.Println("Area and Perimeter Of Rect-1", a, p)
	a1, p1 := areaAndperiOfRect(12.4, 15.5) // this function returns two values
	fmt.Println("Area and Perimeter Of Rect-2", a1, p1)

	a1, _ = areaAndperiOfRect(12.4, 15.5) // _ (blank identifier)
	_, p2 := areaAndperiOfRect(12.4, 15.5)
	fmt.Println("Area and Perimeter Of Rect-3", a1, p2)
	n, _ := fmt.Println("Hello World")
	println(n)

}
