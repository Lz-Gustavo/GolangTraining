package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func readAndIncrement(n *int32, wg *sync.WaitGroup) {

	y := atomic.LoadInt32(n)
	atomic.AddInt32(&y, 1)
	atomic.StoreInt32(n, y)
	fmt.Println("Value:", y, "\tNumGoroutines:", runtime.NumGoroutine())
	wg.Done()
}

func main() {

	var x int32
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go readAndIncrement(&x, &wg)
	}

	wg.Wait()
	fmt.Println("Value:", x, "\tNumGoroutines:", runtime.NumGoroutine())
}
