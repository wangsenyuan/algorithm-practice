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
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Fprintln(writer, "Impossible")
		return
	}
	for _, cur := range res {
		fmt.Fprintln(writer, cur)
	}
}

func drive(reader *bufio.Reader) (a [][]int, res []string) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([][]int, n)
	for i := range n {
		a[i] = make([]int, 3)
		for j := range 3 {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	res = solve(a)
	return
}

const inf = 1 << 60

type data struct {
	l int
	m int
	w int
}

func solve(tasks [][]int) []string {
	n := len(tasks)

	get := func(v int) string {
		switch v {
		case 0:
			return "MW"
		case 1:
			return "LW"
		default:
			return "LM"
		}
	}

	if n == 1 {
		dp := play(tasks)
		best := -1
		for state, v := range dp {
			if v.l == v.m && v.m == v.w {
				if best < 0 || dp[best].l < v.l {
					best = state
				}
			}
		}
		if best == -1 {
			return nil
		}
		res := make([]string, n)
		for i := range n {
			res[i] = get(best % 3)
			best /= 3
		}
		return res
	}

	mid := n / 2
	dp := play(tasks[:mid])
	fp := play(tasks[mid:])

	mem := make(map[data]int)
	for state, v := range fp {
		x := data{v.l - v.m, v.m - v.w, v.w - v.l}
		if y, ok := mem[x]; !ok || fp[y].l < v.l {
			mem[x] = state
		}
	}

	best := []int{-1, -1, -inf}

	for state, v := range dp {
		x := data{v.m - v.l, v.w - v.m, v.l - v.w}
		if next, ok := mem[x]; ok {
			w := fp[next]
			if best[2] < v.l+w.l {
				best[0] = state
				best[1] = next
				best[2] = v.l + w.l
			}
		}
	}

	if best[2] == -inf {
		return nil
	}
	res := make([]string, n)
	for i := range mid {
		res[i] = get(best[0] % 3)
		best[0] /= 3
	}

	for i := mid; i < n; i++ {
		res[i] = get(best[1] % 3)
		best[1] /= 3
	}

	return res
}

func play(arr [][]int) []data {
	n := len(arr)
	T := 1
	for range n {
		T *= 3
	}
	dp := make([]data, T)

	for state := range T {
		var l, m, w int

		cur := state
		for i := range n {
			switch cur % 3 {
			case 0:
				// MW
				m += arr[i][1]
				w += arr[i][2]
			case 1:
				// LW
				l += arr[i][0]
				w += arr[i][2]
			case 2:
				l += arr[i][0]
				m += arr[i][1]
			}
			cur /= 3
		}

		dp[state] = data{l, m, w}
	}

	return dp
}
