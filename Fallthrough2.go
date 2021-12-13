// Golang program to show the
// uses of fallthrough keyword
package main

// Here "fmt" is formatted IO which
// is same as C’s printf and scanf.
import "fmt"

// Main function
func main() {
	gfg := "Geek"

	// Use switch on the day variable.
	switch {
	case gfg == "Geek":
		fmt.Println("Geek")
		fallthrough
	case gfg == "For":
		fmt.Println("For")
		fallthrough
	case gfg == "Geeks":
		fmt.Println("Geeks")
	}
}
