package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func producer(c chan<- int) {
	c <- rand.Int()
	wg.Done()
}

func consumer(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}

func main() {

	c := make(chan int)

	wg.Add(20)
	for i := 0; i < 10; i++ {
		go producer(c)
		go consumer(c)
	}

	wg.Wait()
	fmt.Println("Exit")
}
