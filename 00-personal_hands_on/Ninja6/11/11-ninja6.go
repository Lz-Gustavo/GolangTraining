// solution of URI online judge problem 1161

package main

import (
	"fmt"
	"io"
)

func numFactorial(n int) []int64 {

	xi := make([]int64, 0)

	//fmt.Printf("type: %T len: %d cap: %d\n", xi, len(xi), cap(xi))

	xi = append(xi, 1)
	for i := 1; i <= n; i++ {
		xi = append(xi, xi[i-1]*int64(i))
	}

	return xi
}

func main() {

	xi := numFactorial(20)
	//for i, v := range xi {
	//	fmt.Println(i, "||\t", v)
	//}

	var x int
	var y int
	var err error

	for {
		_, err = fmt.Scanln(&x, &y)
		if err == io.EOF {
			break
		}

		fmt.Println(xi[x] + xi[y])
	}
}
