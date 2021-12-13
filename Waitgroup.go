package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func AnotherProcess(wg *sync.WaitGroup) {
	fmt.Println("This is a another method")
	time.Sleep(time.Second * 1)
	fmt.Println("This another method is done")
	wg.Done()
}

func main() {
	no := 4
	wg := &sync.WaitGroup{}
	wg.Add(no)
	for i := 0; i < no; i++ {
		go process(i, wg)
	}
	wg.Add(1)
	go AnotherProcess(wg)
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
