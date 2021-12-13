package main

import "fmt"

func Add(values ...interface{}) (sum interface{}) {
	for _, v := range values {
		switch v.(type) {
		case int:
			if sum == nil {
				sum = 0
			}
			sum = sum.(int) + v.(int)

		case float64:
			if sum == nil {
				sum = 0.0
			}
			sum = sum.(float64) + v.(float64)

		}
	}
	return sum
}

func main() {

	//fmt.Println(Add(10,10.345)
	fmt.Println(Add(10.23, 10.12, 32.43))
	fmt.Println(Add(10, 10, 20, 30, 30, 42, 23))

}

// slice, map , chan, any pointer , interface
