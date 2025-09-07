package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) []int {
	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)
	a := make([]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(a, b, queries)
}

func solve(a []int, b []int, queries [][]int) []int {
	var s0 int

	n := len(a)
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			s0 += a[i]
		} else {
			s0 -= a[i]
		}
	}

	m := len(b)
	var s1 int
	for i := range n {
		if i&1 == 0 {
			s1 -= b[i]
		} else {
			s1 += b[i]
		}
	}
	var arr []int
	arr = append(arr, s1)
	for i := n; i < m; i++ {
		s1 += b[i-n]
		s1 *= -1

		if n&1 == 1 {
			s1 -= b[i]
		} else {
			s1 += b[i]
		}
		arr = append(arr, s1)
	}
	sort.Ints(arr)

	calc := func(x int) int {
		i := sort.SearchInts(arr, x)
		var res int = inf
		if i < len(arr) {
			res = arr[i] - x
		}
		if i > 0 {
			res = min(res, x-arr[i-1])
		}
		return res
	}

	find := func() int {
		return calc(-s0)
	}

	ans := make([]int, len(queries)+1)
	ans[0] = find()

	for i, cur := range queries {
		l, r, x := cur[0], cur[1], cur[2]
		if (r-l+1)&1 == 0 {
			ans[i+1] = ans[i]
		} else {
			if l&1 == 1 {
				s0 += x
			} else {
				s0 -= x
			}
			ans[i+1] = find()
		}
	}
	return ans
}

const inf = 1e18
