package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		time.Sleep(3 * time.Second)
		close(c)
	}()

	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Println(v, ok)
	}
}
