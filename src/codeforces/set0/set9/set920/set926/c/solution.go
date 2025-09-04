package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res := solve(a)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(a []int) bool {
	n := len(a)

	x := slices.Min(a)
	y := slices.Max(a)
	if x == y {
		return true
	}
	// x == 0 and y == 1
	var w int
	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		if w == 0 {
			w = i - j
		}
		if i-j != w {
			return false
		}
	}
	return true
}
