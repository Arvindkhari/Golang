package main

import "fmt"

func main() {

	e := &Employee{No: 1001, Name: "Jiten"}
	s := &Student{No: 1111, Name: "Karthi", Grade: "1st"}

	var iStringer Stringer

	iStringer = e
	fmt.Println(iStringer.ToString())

	iStringer = s
	fmt.Println(iStringer.ToString())

}

type Stringer interface {
	ToString() string
}

type Employee struct {
	No   int
	Name string
}

func (e *Employee) SetDetails(no int, name string) {
	e.No = no
	e.Name = name
}

func (e *Employee) ToString() string {
	return fmt.Sprint(e)
}

type Student struct {
	No    int
	Name  string
	Grade string
}

func (s *Student) ToString() string {
	return fmt.Sprint(s)
}
