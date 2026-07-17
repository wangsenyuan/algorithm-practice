package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
		if queries[i][0] == 1 {
			fmt.Fscan(reader, &queries[i][2])
		}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []int {
	fa := make([]int, n)
	R := make([]int, n)
	color := make([]int, n)
	for i := range n {
		fa[i] = i
		R[i] = i
		color[i] = i
	}

	find := func(x int) int {
		y := fa[x]
		for y != fa[y] {
			y = fa[y]
		}

		for fa[x] != y {
			fa[x], x = y, fa[x]
		}

		return y
	}

	// color i
	freq := make([]int, n)
	for i := range n {
		freq[i] = 1
	}

	var ans []int

	change := func(x int, c int) {
		l := find(x)
		if color[l] == c {
			// no change
			return
		}

		r := R[l]

		freq[color[l]] -= r - l + 1

		if l > 0 && color[find(l-1)] == c {
			// merge
			l1 := find(l - 1)
			R[l1] = r
			fa[l] = l1
		}
		if r < n-1 && color[r+1] == c {
			r1 := R[r+1]
			R[find(l)] = r1
			fa[r+1] = r
		}
		freq[c] += r - l + 1
		color[find(l)] = c
	}

	for _, cur := range queries {
		if cur[0] == 1 {
			x, c := cur[1], cur[2]
			// change contiguous cells with the same color
			change(x-1, c-1)
		} else {
			c := cur[1]
			ans = append(ans, freq[c-1])
		}
	}

	return ans
}
