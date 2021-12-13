package main

import (
	"fmt"
)

func main() {

	// print numbers from 1-100 which are divisible by 7

	for i := 1; i <= 100; i++ {

		if i%3 != 0 {
			continue
		}

		if i%7 == 0 {
			fmt.Print(i, " ")
		}

		/*if i%3 == 0 && i%7==0{
		fmt.Println(i," divisible by 3 and 7")
		}*/

	}

}
