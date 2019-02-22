package main

import "fmt"

func main() {

	a := 10
	b := 42
	c := 10
	d := 24

	e := a == c
	f := a <= d
	g := b >= a
	h := d != c
	i := a < b
	j := b > c

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
}
