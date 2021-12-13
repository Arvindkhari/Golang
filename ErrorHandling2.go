package main

import (
	"fmt"
)

func main() {
	sum, err := Add(12, 13, 14.45, 25.3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
	sum, err = Add(12, 13, true)
	if err != nil {
		fmt.Println(err)
	}

}

func Add(nums ...interface{}) (float64, error) {
	sum := 0.0
	for _, num := range nums {
		switch num.(type) {
		case int:
			sum = sum + float64(num.(int))
		case float64:
			sum = sum + num.(float64)
		default:
			//return 0.0, errors.New("invalid data")
			return 0.0, fmt.Errorf("Invalid Data.Provided type is %T and Value is %v", num, num)

		}

	}
	return sum, nil

}
