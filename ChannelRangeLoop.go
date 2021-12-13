//range loop for channels
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello")
	ch := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			time.Sleep(time.Millisecond * 1)
			ch <- i
		}
		//i := 0
		/*for {
			time.Sleep(time.Millisecond * 1)
			i++
			ch <- i
		}*/
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
		if c == 10 {
			os.Exit(0)
			//break
			//return
		}
	}
}
