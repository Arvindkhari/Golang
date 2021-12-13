//gorm which is called go object relationship mapping
package main

import "fmt"

func main() {
	var e1 Employee

	e1.id = 101
	e1.name = "jiten"
	e1.line1 = "Flat no 202" // promoted fields can be called directly
	e1.city = "Bangalore"
	e1.Address.city = "Bangalore-2"
	e1.country = "India"
	e1.Addresses = append(e1.Addresses, Address{line1: "line1", city: "Bangalore", country: "India"})
	e1.Addresses = append(e1.Addresses, Address{line1: "Times Sqare", city: "Madrid", country: "Spain"})

	fmt.Println(e1)
}

type Employee struct {
	id        int
	name      string
	city      string
	Address   // promoted field
	Addresses []Address
}

type Address struct {
	line1   string
	city    string
	country string
}
