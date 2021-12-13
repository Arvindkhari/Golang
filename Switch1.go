// Go program to illustrate the
// concept of Expression switch
// statement
package main

import "fmt"

func main() {

	// Switch statement with both
	// optional statement, i.e, day:=4
	// and expression, i.e, day
	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

}
