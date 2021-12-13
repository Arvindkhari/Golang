//Q.
//Insert(key,value)
//Update(key,value)(bool)
//Delete(key)(bool)
//Read(key)(string)
//ReadAll() --> Print everything
//ToStrig() (string)
//ToByte()([]byte)
//GetKeys()([]string)
//GetValues()([]string)
//ReadInOrder()

package main

import (
	//"demo/area"

	_ "demo/mymath/calc"
	"fmt"
	_ "myDemo/mymath"
)

func main() {
	println("Hello World")

	var myMap MyMap

	myMap = make(map[string]string)

	myMap.Insert("1001", "Jiten")
	myMap.Insert("1002", "Rahim")
	myMap.Insert("1003", "Johan")

	fmt.Println(myMap)
	fmt.Println(myMap.ToByte())

	fmt.Println("Keys of the map:", myMap.GetKeys())
	fmt.Println("Valuesof the map:", myMap.GetVals())

	fmt.Println(myMap.Update("1004", "Anila"))
	fmt.Println(myMap.Update("1002", "Anila"))
	fmt.Println("Valuesof the map:", myMap.GetVals())

}

type MyMap map[string]string

func (mm MyMap) ToByte() []byte {
	return []byte(fmt.Sprint(mm))
}

func (mm MyMap) Insert(key, val string) {
	if mm != nil {
		mm[key] = val
	}
}

func (mm MyMap) Update(key, val string) bool {
	if mm == nil {
		return false
	}
	_, ok := mm[key]
	if !ok {
		return ok
	}

	mm[key] = val
	return true

}

func (mm MyMap) GetKeys() []string {
	if mm == nil {
		return nil
	}
	var slice []string
	for k, _ := range mm {
		slice = append(slice, k)
	}
	return slice
}

func (mm MyMap) GetVals() []string {
	if mm == nil {
		return nil
	}
	var slice []string
	for _, val := range mm {
		slice = append(slice, val)
	}
	return slice
}
