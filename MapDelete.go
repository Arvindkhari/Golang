package main

import "fmt"

func main() {
	var empmap map[int]float32

	empmap = make(map[int]float32, 0)

	empmap[1001] = 12345.54
	empmap[1002] = 13343.43
	empmap[1003] = 15131.90
	empmap[100100] = 15112331.90
	// delete an element from the map

	delete(empmap, 100100)
	fmt.Println(empmap)
	// iterate through the map
	var totSal float32
	for key, val := range empmap {
		fmt.Println("Employee ID:", key, " Sal:", val)
		totSal += val
	}
	fmt.Println("Total Sal", totSal)

	// to check key exists or not
	val, ok := empmap[12312] // first is value and second return "ok" is bool . If true key exists else does not
	fmt.Println(val, ok)

}

// println, complex, len, cap , make, append, copy , delete
// by using map perform add, update, delete and get opearations .. add(map, key, value) update(map, key, value) , delete has to delete the elements in the map ,
// deleteall

// take a slice .. apend 100 elements in the slice .. using rand package.
// find out how many times each number occured in the slice .. using maps.
