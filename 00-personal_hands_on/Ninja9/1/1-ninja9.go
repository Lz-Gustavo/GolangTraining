package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func increment(x *int) {
	*x++
	wg.Done()
}

func decrement(x *int) {
	*x--
	wg.Done()
}

func main() {

	n := 4
	fmt.Println("N =", n)
	wg.Add(2)

	go increment(&n)
	go decrement(&n)

	wg.Wait()
	fmt.Println("N =", n)
}
