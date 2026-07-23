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

	cost, grid, _, _ := drive(reader)
	fmt.Fprintln(writer, cost)
	for _, row := range grid {
		fmt.Fprintln(writer, row)
	}
}

func drive(reader *bufio.Reader) (cost int, res []string, a [][]int, important [][]int) {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a = make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	important = make([][]int, k)
	for i := range k {
		important[i] = make([]int, 2)
		fmt.Fscan(reader, &important[i][0], &important[i][1])
	}
	cost, res = solve(a, important)
	return
}

type state struct {
	val      int
	fromMask int
	fromPos  int
}

const inf = 1 << 60

func solve(a [][]int, important [][]int) (int, []string) {
	n := len(a)
	m := len(a[0])
	k := len(important)

	var dd = []int{-1, 0, 1, 0, -1}

	dist := make([][]int, n*m)
	next := make([][]int, n*m)
	for i := range n * m {
		dist[i] = make([]int, n*m)
		next[i] = make([]int, n*m)
		for j := range n * m {
			dist[i][j] = inf
			next[i][j] = -1
		}
		dist[i][i] = a[i/m][i%m]
		next[i][i] = i
		r, c := i/m, i%m
		for d := range 4 {
			nr, nc := r+dd[d], c+dd[d+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m {
				j := nr*m + nc
				dist[i][j] = a[nr][nc] + a[r][c]
				next[i][j] = j
			}
		}
	}

	for k := range n * m {
		for i := range n * m {
			for j := range n * m {
				cand := dist[i][k] + dist[k][j] - a[k/m][k%m]
				if cand < dist[i][j] {
					dist[i][j] = cand
					next[i][j] = next[i][k]
				}
			}
		}
	}

	dp := make([][]state, 1<<k)
	for i := range 1 << k {
		dp[i] = make([]state, n*m)
		for j := range n * m {
			dp[i][j] = state{val: inf, fromMask: -1, fromPos: -1}
		}
	}

	for i := range k {
		r, c := important[i][0]-1, important[i][1]-1
		dp[1<<i][r*m+c].val = a[r][c]
	}

	for mask := 1; mask < 1<<k; mask++ {
		for sub := mask; sub > 0; sub = (sub - 1) & mask {
			if sub != mask {
				other := mask ^ sub
				for v := range n * m {
					if dp[sub][v].val+dp[other][v].val-a[v/m][v%m] < dp[mask][v].val {
						dp[mask][v].val = dp[sub][v].val + dp[other][v].val - a[v/m][v%m]
						dp[mask][v].fromMask = sub
						dp[mask][v].fromPos = v
					}
				}
			}
		}
		for v := range n * m {
			best := dp[mask][v].val
			bestPos := v
			for u := range n * m {
				if dp[mask][u].val+dist[u][v]-a[u/m][u%m] < best {
					best = dp[mask][u].val + dist[u][v] - a[u/m][u%m]
					bestPos = u
				}
			}
			if best != dp[mask][v].val {
				dp[mask][v].val = best
				dp[mask][v].fromMask = mask
				dp[mask][v].fromPos = bestPos
			}
		}
	}

	all := (1 << k) - 1
	best := inf
	bestPos := -1
	for v := range n * m {
		if dp[all][v].val < best {
			best = dp[all][v].val
			bestPos = v
		}
	}

	buf := make([][]byte, n)
	for i := range n {
		buf[i] = make([]byte, m)
		for j := range m {
			buf[i][j] = '.'
		}
	}

	markPath := func(from int, to int) {
		buf[from/m][from%m] = 'X'
		for from != to {
			from = next[from][to]
			buf[from/m][from%m] = 'X'
		}
	}

	var f func(mask int, pos int)
	f = func(mask int, pos int) {
		if mask == 0 {
			return
		}
		buf[pos/m][pos%m] = 'X'
		cur := dp[mask][pos]
		if cur.fromMask == -1 {
			return
		}
		if cur.fromMask == mask {
			markPath(cur.fromPos, pos)
			f(mask, cur.fromPos)
			return
		}
		sub := cur.fromMask
		f(sub, pos)
		f(mask^sub, pos)
	}

	f(all, bestPos)

	ans := make([]string, n)
	for i := range n {
		ans[i] = string(buf[i])
	}
	return best, ans
}
