package main

import (
	"fmt"
	"runtime"
	"sync"
)

func readAndIncrement(n *int, wg *sync.WaitGroup) {

	y := *n
	runtime.Gosched()
	y++
	*n = y
	wg.Done()
}

func main() {

	x := 0
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go readAndIncrement(&x, &wg)
		fmt.Println("Value:", x, "\tNumGoroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Value:", x, "\tNumGoroutines:", runtime.NumGoroutine())
}
