package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, q := drive(reader)
	if len(q) == 0 {
		fmt.Println("-1")
		return
	}
	s := fmt.Sprintf("%v", q)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (p []int, q []int) {
	var n int
	fmt.Fscan(reader, &n)
	p = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	q = solve(p)
	return
}

func solve(p []int) []int {
	n := len(p)
	for i := range n {
		p[i]--
	}

	var cycles [][]int
	marked := make([]bool, n)

	for i := range n {
		if !marked[i] {
			j := i
			var cur []int
			for !marked[j] {
				cur = append(cur, j)
				marked[j] = true
				j = p[j]
			}
			cycles = append(cycles, cur)
		}
	}

	slices.SortFunc(cycles, func(a, b []int) int {
		return cmp.Or(len(b)-len(a), a[0]-b[0])
	})

	// 奇数圈保留，偶数圈合并

	q := make([]int, n)

	buf := make([]int, n+3)

	odd := func(cur []int) {
		m := len(cur)

		for i, j := 0, 0; i < m; i, j = i+1, (j+2)%m {
			buf[j] = cur[i]
		}

		for i := range m {
			q[buf[i]] = buf[i+1]
		}
		q[buf[m-1]] = buf[0]
	}

	// 还是不对
	even := func(a, b []int) {
		m := len(a)
		last := a[0]
		var l, r int
		for j := range 2*m - 1 {
			if j&1 == 0 {
				q[last] = b[r]
				last = b[r]
				r++
			} else {
				l++
				q[last] = a[l]
				last = a[l]
			}
		}
		q[last] = a[0]
	}

	for i := 0; i < len(cycles); {
		if len(cycles[i])%2 == 1 {
			odd(cycles[i])
			i++
		} else {
			// even
			if i == len(cycles)-1 || len(cycles[i+1]) != len(cycles[i]) {
				return nil
			}

			even(cycles[i], cycles[i+1])
			i += 2
		}
	}

	for i := range n {
		p[i]++
		q[i]++
	}

	return q
}
