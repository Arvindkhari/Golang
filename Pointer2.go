package main

import "fmt"

func main() {

	var e1 Employee
	e1 = Employee{id: 101, age: 37, name: "jiten"}
	fmt.Println(e1.name)

	e2 := Employee{id: 102, age: 38, name: "Rajesh", address: "Bangalore", email: "JitenP@outlook.com"}
	fmt.Println(e2)

	e3 := Employee{}
	e3.id = 103
	e3.age = 40
	e3.name = "Rahim"
	fmt.Println(e3)

	e4 := &Employee{id: 104, name: "Jessy", age: 42}
	fmt.Println(e4)

	var e5 *Employee

	e5 = &Employee{id: 106, name: "Happy Singh"}
	fmt.Println(e5)

	e6 := new(Employee)
	e6.id = 107
	fmt.Println(e6)

	var e7 Employee
	e7.name = "Johan"
	fmt.Println("-->", e7)

	var e8 *Employee
	fmt.Println(e8)

}

type Employee struct {
	id      int
	age     uint8
	name    string
	address string
	email   string
}
