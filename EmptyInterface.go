package main

import "fmt"

func Add(values ...interface{}) (sum interface{}) {
	for _, v := range values {
		fmt.Println("Value is ", v)
		//sum = sum + v
	}
	return sum
}

func main() {

	/*Add(10)
	Add("Hello", " World")
	Add(12.34, 43123, 43.34)
	Add(true, false, false, 12, "Hello", 43.34)*/

	//var n int
	//var f float32
	//var s string

	var i interface{} = 12

	//n = i.(int) // interface{}.(T)
	// cannot use type casting ,instead have to use type assertion
	//fmt.Println(n)

	//i = "Hello World"

	//s = i.(string)

	//fmt.Println(s)
	var f32 float32 = 12.3
	i = f32

	switch i.(type) {
	case int:
		fmt.Println("It is of type int")
	case float32:
		fmt.Println("It is float32")
	case float64:
		fmt.Println("It is float64")
	case string:
		fmt.Println("It is string")
	default:
		fmt.Println("It is some other type")

	}

}
