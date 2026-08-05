package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	todo := make([][]int, n+1)
	tr := make(Fenwick, n+2)

	dp := make([]int, n+1)

	for j, v := range a {

		j++
		var best int
		if j-v-1 > 0 {
			best = tr.query(j - v - 1)
		}
		best += v
		dp[j] = best
		if j+v <= n {
			todo[j+v] = append(todo[j+v], j)
		}
		for _, i := range todo[j] {
			tr.update(i, dp[i])
		}
	}

	return slices.Max(dp)
}

type Fenwick []int

func (f Fenwick) update(i int, v int) {
	i++
	for i < len(f) {
		f[i] = max(f[i], v)
		i += i & -i
	}
}

func (f Fenwick) query(i int) int {
	i++
	res := 0
	for i > 0 {
		res = max(res, f[i])
		i -= i & -i
	}
	return res
}
