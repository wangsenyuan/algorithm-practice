package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, v := range res {
		buf.WriteString(fmt.Sprintf("%d\n", v))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		var l, x int
		fmt.Fscan(reader, &l, &x)
		queries[i] = []int{l, x}
	}
	return solve(a, queries)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

type pair struct {
	first  int
	second int
}

func solve(a []int, queries [][]int) []int {
	n := len(a)

	pw := make([]int, n+1)
	pw[0] = 1
	for i := 1; i <= n; i++ {
		pw[i] = add(pw[i-1], pw[i-1])
	}

	qs := make([][]pair, n)

	for i, cur := range queries {
		j, x := cur[0], cur[1]
		qs[j-1] = append(qs[j-1], pair{x, i})
	}

	basis := make([]int, 20)

	var num int

	add := func(x int) {
		for i := 19; i >= 0; i-- {
			if (x>>i)&1 == 0 {
				continue
			}
			if basis[i] == 0 {
				basis[i] = x
				num++
				break
			}
			x ^= basis[i]
		}
	}

	check := func(x int) bool {
		for i := 19; i >= 0; i-- {
			if (x>>i)&1 == 0 {
				continue
			}
			if basis[i] == 0 {
				return false
			}
			x ^= basis[i]
		}
		return true
	}

	ans := make([]int, len(queries))

	for i := 0; i < n; i++ {
		add(a[i])

		for _, q := range qs[i] {
			x, j := q.first, q.second
			if check(x) {
				ans[j] = pw[i+1-num]
			}
		}
	}

	return ans
}
