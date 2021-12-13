package main

import "fmt"

func main() {

	e1 := Employee{id: 101, name: "Jiten", address: Address{addressline1: "FLatno 202, Sriraj", city: "Bangalore", country: "India"}}
	fmt.Println(e1)

}

type Employee struct {
	id      int
	name    string
	address Address
}

type Address struct {
	addressline1 string
	city         string
	country      string
}
