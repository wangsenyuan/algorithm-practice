package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(k, a, b)
}

func solve(k int, a []int, b []int) int {
	n := len(a)
	// m := len(b)

	// a[0] 是第一个 icon
	pos := make([]int, n)
	for i := range n {
		a[i]--
		pos[a[i]] = i
	}

	lauch := func(x int) int {
		i := pos[x]
		res := i / k
		if i > 0 {
			a[i-1], a[i] = a[i], a[i-1]
			pos[a[i-1]] = i - 1
			pos[a[i]] = i
		}

		return res + 1
	}
	var ans int
	for _, x := range b {
		ans += lauch(x - 1)
	}
	return ans
}
