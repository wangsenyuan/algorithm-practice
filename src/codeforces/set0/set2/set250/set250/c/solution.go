package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	a = slices.Compact(a)
	n := len(a)

	save := make([]int, k+1)

	for i := range n {
		if i > 0 {
			save[a[i]]++
		}
		if i+1 < n {
			save[a[i]]++
		}
		if i > 0 && i+1 < n && a[i-1] != a[i+1] {
			save[a[i]]--
		}
	}
	best := k
	for i := k - 1; i > 0; i-- {
		if save[i] >= save[best] {
			best = i
		}
	}
	return best
}
