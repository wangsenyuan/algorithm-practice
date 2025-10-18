package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%d:%d\n", res[0], res[1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	var m int
	fmt.Fscan(reader, &m)
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}

	return solve(a, b)
}

func solve(a []int, b []int) []int {
	n := len(a)
	m := len(b)
	sort.Ints(a)
	sort.Ints(b)
	c := make([]int, n+m)
	copy(c, a)
	copy(c[n:], b)
	sort.Ints(c)
	// 2 * n + n - x - (2 * m + m - y) 最大
	// d不一定是整数
	best := n - m
	var d int
	var i, j int
	for _, v := range c {
		for i < n && a[i] <= v {
			i++
		}
		for j < m && b[j] <= v {
			j++
		}
		if (n-i)-(m-j) > best {
			best = (n - i) - (m - j)
			d = v
		}
	}
	ans := make([]int, 2)
	ans[0] += 2 * n
	ans[1] += 2 * m

	x := sort.Search(n, func(i int) bool {
		return a[i] > d
	})
	ans[0] += n - x
	y := sort.Search(m, func(i int) bool {
		return b[i] > d
	})
	ans[1] += m - y

	return ans
}
