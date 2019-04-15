package main

import "log"

type customErr struct {
	err string
}

func (c *customErr) Error() string {
	return c.err
}

func foo(er error) {
	log.Println(er)
}

func main() {

	er := &customErr{"My custom error"}
	foo(er)
}
