package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(time.Millisecond * 10)
		ch <- "from server1"
	}
}
func server2(ch chan string) {
	for {
		time.Sleep(time.Millisecond * 10)
		ch <- "from server2"
	}

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	timeout := make(chan bool)

	go server1(output1)
	go server2(output2)
	go func() {
		time.Sleep(time.Second * 1)
		timeout <- true
	}()

	time.Sleep(1 * time.Millisecond)
outer:
	for {
		select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
		case <-timeout:
			fmt.Println("Timedout ------")
			break outer

		}
	}
	fmt.Println("This is done")
}
