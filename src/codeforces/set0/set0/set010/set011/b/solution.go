package main

import (
	"fmt"
	"sort"
)

func main() {
	var x int
	fmt.Scanf("%d\n", &x)
	res := solve(x)
	fmt.Printf("%d\n", res)
}

func solve(x int) int {
	if x < 0 {
		x = -x
	}
	if x == 0 {
		return 0
	}
	n := sort.Search(x+1, func(i int) bool {
		return i*(i+1)/2 >= x
	})
	sum := n * (n + 1) / 2
	diff := sum - x
	if diff%2 == 0 {
		return n
	}
	// diff is odd
	for diff%2 == 1 {
		n++
		sum += n
		diff = sum - x
	}

	return n
}
