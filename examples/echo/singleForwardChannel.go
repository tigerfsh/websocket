package main

import "fmt"

func main() {

	ch := make(chan int)
	go func(ch <-chan int) {
		for {
			val, ok := <-ch
			if !ok {
				return
			}
			fmt.Printf("val : %v\n", val)
		}
	}(ch)

	for i := 0; i < 20; i++ {
		ch <- i
	}
	close(ch)

}
