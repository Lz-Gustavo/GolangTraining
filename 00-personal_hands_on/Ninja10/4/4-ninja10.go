package main

import (
	"fmt"
)

func main() {
	quit := make(chan int)
	c := generateValues(quit)

	receive(c, quit)

	fmt.Println("about to exit")
}

func generateValues(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
	}()
	return c
}

func receive(c, q <-chan int) {

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
