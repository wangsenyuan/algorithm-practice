package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var m int
	fmt.Fscan(reader, &m)
	q := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &q[i])
	}
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(q, a)
}

func solve(q []int, a []int) int {
	slices.Sort(a)
	slices.Reverse(a)

	var ans int
	w := slices.Min(q)

	for len(a) > 0 {
		v := len(a)
		for range min(w, v) {
			ans += a[0]
			a = a[1:]
		}
		if len(a) < 2 {
			break
		}
		a = a[2:]
	}

	return ans
}
