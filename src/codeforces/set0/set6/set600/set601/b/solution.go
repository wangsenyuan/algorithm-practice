package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(a, queries)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	stack := make([]pair, n)
	L := make([]int, n)
	calc := func(l int, r int) int {
		if l == r {
			return 0
		}
		var top int
		for i := l; i < r; i++ {
			d := abs(a[i+1] - a[i])
			for top > 0 && stack[top-1].first < d {
				top--
			}
			if top > 0 {
				L[i] = stack[top-1].second + 1
			} else {
				L[i] = l
			}
			stack[top] = pair{d, i}
			top++
		}

		var res int
		top = 0
		for i := r - 1; i >= l; i-- {
			d := abs(a[i+1] - a[i])
			for top > 0 && stack[top-1].first <= d {
				top--
			}
			// d > stack[top-1].first
			R := r - 1
			if top > 0 {
				R = stack[top-1].second - 1
			}
			res += d * (i - L[i] + 1) * (R - i + 1)
			stack[top] = pair{d, i}
			top++
		}

		return res
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		ans[i] = calc(cur[0]-1, cur[1]-1)
	}
	return ans
}

func abs(num int) int {
	return max(num, -num)
}
