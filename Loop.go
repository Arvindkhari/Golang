// Go program to illustrate the
// use of simple range loop
package main

import "fmt"

// Main function
func main() {

	// Here rvariable is a array
	rvariable := []string{"GFG", "Geeks", "GeeksforGeeks"}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	for i, j := range rvariable {
		fmt.Println(i, j)
	}

}
