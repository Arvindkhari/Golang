package main

import (
	"fmt"
)

func Sender(sender chan int) {
	sender <- 100
}

func Receiver(receiver chan int) int {
	return <-receiver
}
func main() {
	fmt.Println("Hello")

	ch := make(chan int)
	done := make(chan bool)

	go Sender(ch)

	go func() {
		b := Receiver(ch)
		fmt.Println(b)
		done <- true
	}()
	<-done
}
