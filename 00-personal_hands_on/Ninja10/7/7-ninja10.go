package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func add10Values(start int, c chan<- int) {

	for i := start; i < start+10; i++ {
		c <- i
	}
	wg.Done()
}

func main() {

	// if uncomment this line, numbers will be output in order (sequential execution)
	//runtime.GOMAXPROCS(1)

	c := make(chan int)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go add10Values(i*10, c)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	// bellow is not gonna consume any data because channel is already empty
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
