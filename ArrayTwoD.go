package main

import (
	"fmt"
)

func main() {
	//arr := [2]string{"name", "jiten"}

	arr := [2][2][3]string{
		{
			{
				"apple", "banana", "orange",
			},

			{
				"5", "6", "7",
			},
		},
		{
			{
				"car", "jeep", "van",
			},

			{
				"2", "3", "4",
			},
		},
	}

	arr[0][0][1] = "peach"

	for _, a := range arr {
		for _, e := range a {
			fmt.Println(e)
		}

	}

	//fmt.Println(arr)
}
