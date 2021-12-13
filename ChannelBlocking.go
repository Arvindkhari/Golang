// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

var ch chan int

func init() {
	if ch == nil {
		ch = make(chan int)
	}
	fmt.Println("Hello is it called")
}

func init() {
	fmt.Println("Hey")
}

func Incr(r string, ns int) {
	for {
		k := <-ch
		time.Sleep(time.Duration(ns))

		fmt.Println("The value of ch on", r, " :", k)
		if k < 10000 {
			k = k + 1
			ch <- k // blocked here when use only one function
		}
	}
}

func Incr2(r string, ns int) {
	for {
		time.Sleep(time.Second * 1)
		k := <-ch
		time.Sleep(time.Duration(ns))

		fmt.Println("The value of ch on", r, " :", k)
		if k < 10000 {
			k = k + 1
			ch <- k
		}
	}
}

func main() {

	go Incr("1st", 20)
	go Incr("2nd", 3)
	//go Incr("3rd", 8)
	//go Incr2("Otherfunc 4th", 7)
	//go func() {
	ch <- 1
	//}()
	fmt.Println("Hello --")
	time.Sleep(time.Second * 10)
}
