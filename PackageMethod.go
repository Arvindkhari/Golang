package main

import (
	"fmt"
	"reflect"
)

func main() {

	//var i int = 100
	var i1 MyInt = 200

	str := (&i1).ToString()
	fmt.Println(reflect.TypeOf(str))

	//str1 := MyInt(i).ToString()
	//fmt.Println(str1)

}

type MyInt int

func (m *MyInt) ToString() string {
	return fmt.Sprint(m)
}

type MyMap map[string]string
