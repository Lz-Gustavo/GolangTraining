package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := generateValues()

	wg.Add(1)
	go receive(c)
	wg.Wait()

	fmt.Println("about to exit")
}

func generateValues() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	wg.Done()
}
