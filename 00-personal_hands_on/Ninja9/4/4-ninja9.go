package main

import (
	"fmt"
	"runtime"
	"sync"
)

var sem sync.Mutex

func readAndIncrement(n *int, wg *sync.WaitGroup) {

	sem.Lock()
	defer sem.Unlock()

	y := *n
	y++
	*n = y
	fmt.Println("Value:", *n, "\tNumGoroutines:", runtime.NumGoroutine())
	wg.Done()
}

func main() {

	x := 0
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go readAndIncrement(&x, &wg)
	}

	wg.Wait()
	fmt.Println("Value:", x, "\tNumGoroutines:", runtime.NumGoroutine())
}
