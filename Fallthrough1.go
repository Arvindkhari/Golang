// Golang program to show the
// uses of fallthrough keyword
package main

// Here "fmt" is formatted IO which
// is same as C’s printf and scanf.
import "fmt"

// Main function
func main() {
	day := "Tue"

	// Use switch on the day variable.
	switch {
	case day == "Mon":
		fmt.Println("Monday")
		fallthrough
	case day == "Tue":
		fmt.Println("Tuesday")
		fallthrough
	case day == "Wed":
		fmt.Println("Wednesday")
	}
}
