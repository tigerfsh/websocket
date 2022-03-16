package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan string)

	go func() {
		<-time.After(5 * time.Second)
		close(done)
	}()

	what := <-done
	fmt.Printf("What is : %v\n", what)

	what = <-done
	fmt.Printf("What is : %v\n", what)

	ch := make(chan int, 1)
	ch <- 100

	close(ch)
	val, ok := <-ch
	fmt.Println(ok, val)

	val, ok = <-ch
	fmt.Println(ok, val)

	val, ok = <-ch
	fmt.Println(ok, val)




}
