package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	less = iota
	equal
	greater
)

const inf = 1 << 60

type parent struct {
	px int
	py int
	pb int
	qb int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		res := solve(x, y)
		fmt.Fprintln(writer, res[0], res[1])
	}
}

func solve(x int, y int) []int {
	dp := make([][3][3]int, 32)
	par := make([][3][3]parent, 32)

	for i := range 32 {
		for a := range 3 {
			for b := range 3 {
				dp[i][a][b] = inf
			}
		}
	}

	dp[0][equal][equal] = 0

	for step := 0; step < 31; step++ {
		bit := 30 - step
		xb := (x >> bit) & 1
		yb := (y >> bit) & 1
		w := 1 << bit

		for a := range 3 {
			for b := range 3 {
				dp[step+1][a][b] = inf
			}
		}

		for sx := range 3 {
			for sy := range 3 {
				cur := dp[step][sx][sy]
				if cur == inf {
					continue
				}

				for _, choice := range [][2]int{{0, 0}, {1, 0}, {0, 1}} {
					pb, qb := choice[0], choice[1]
					nx, cx := transition(sx, xb, pb, w)
					ny, cy := transition(sy, yb, qb, w)
					val := cur + cx + cy
					if val < dp[step+1][nx][ny] {
						dp[step+1][nx][ny] = val
						par[step+1][nx][ny] = parent{sx, sy, pb, qb}
					}
				}
			}
		}
	}

	best := inf
	bx, by := equal, equal
	for sx := range 3 {
		for sy := range 3 {
			if dp[31][sx][sy] < best {
				best = dp[31][sx][sy]
				bx, by = sx, sy
			}
		}
	}

	var p, q int
	for step := 31; step > 0; step-- {
		prev := par[step][bx][by]
		bit := 31 - step
		p |= prev.pb << bit
		q |= prev.qb << bit
		bx, by = prev.px, prev.py
	}

	return []int{p, q}
}

func transition(state int, want int, have int, weight int) (int, int) {
	if state == less {
		return less, (want - have) * weight
	}
	if state == greater {
		return greater, (have - want) * weight
	}
	if have < want {
		return less, weight
	}
	if have > want {
		return greater, weight
	}
	return equal, 0
}
