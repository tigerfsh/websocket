package main

import (
	"fmt"
	"time"
)

func main() {
	// runtime.GOMAXPROCS(1)
	ch := make(chan int)
	go func() {
		for {
			num := <-ch
			fmt.Printf("from go1 : %v\n", num)
			ch <- num + 1
		}
	}()

	go func() {
		ch <- 1

		for {
			num := <-ch
			fmt.Printf("from go2 : %v\n", num)
			ch <- num + 1
		}
	}()

	<-time.After(100 * time.Second)
}
