package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	// just to show the need of deep-copy on slices...
	xi2 := make([]int, len(xi))
	xs2 := make([]string, len(xs))

	copy(xi2, xi)
	copy(xs2, xs)

	sort.Ints(xi)
	fmt.Println("sorted:", xi)
	fmt.Println("unsorted:", xi2)

	sort.Strings(xs)
	fmt.Println("sorted:", xs)
	fmt.Println("unsorted:", xs2)
}
