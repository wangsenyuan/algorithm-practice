package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	if len(res) == 0 {
		fmt.Fprintln(writer, "NO")
		return
	}

	fmt.Fprintln(writer, "YES", len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Fprintln(writer, s[1:len(s)-1])
}

func drive(reader *bufio.Reader) (rects [][]int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	rects = make([][]int, n)
	for i := range n {
		rects[i] = make([]int, 4)
		fmt.Fscan(reader, &rects[i][0], &rects[i][1], &rects[i][2], &rects[i][3])
	}
	res = solve(slices.Clone(rects))
	return
}

func solve(rects [][]int) []int {
	var mx, my int
	for _, rect := range rects {
		mx = max(mx, rect[2])
		my = max(my, rect[3])
	}

	dp := make([][]int, mx+1)
	fp := make([][]int, mx+1)
	for i := range mx + 1 {
		dp[i] = make([]int, my+1)
		fp[i] = make([]int, my+1)
		for j := range my + 1 {
			fp[i][j] = -1
		}
	}

	for i, rect := range rects {
		x2, y2 := rect[2], rect[3]
		fp[x2][y2] = i
	}

	for i := range mx + 1 {
		for j := range my + 1 {
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
			if i > 0 && j > 0 {
				dp[i][j] -= dp[i-1][j-1]
			}
			if fp[i][j] != -1 {
				v := fp[i][j]
				x1, y1 := rects[v][0], rects[v][1]
				dp[i][j] += (i - x1) * (j - y1)
			}
		}
	}

	up := make([][]int, mx+1)
	rg := make([][]int, mx+1)
	for i := range mx + 1 {
		up[i] = make([]int, my+1)
		rg[i] = make([]int, my+1)
		for j := range my + 1 {
			up[i][j] = j
			rg[i][j] = i
		}
	}

	// 从右到左处理
	slices.SortFunc(rects, func(a, b []int) int {
		return b[0] - a[0]
	})

	getUp := func(x, y int) int {
		y1 := y
		for up[x][y1] != y1 {
			y1 = up[x][y1]
		}
		for y != y1 {
			up[x][y], y = y1, up[x][y]
		}
		return y1
	}

	getRg := func(x int, y int) int {
		x1 := x
		for rg[x1][y] != x1 {
			x1 = rg[x1][y]
		}
		for x != x1 {
			rg[x][y], x = x1, rg[x][y]
		}
		return x1
	}

	for _, rect := range rects {
		x1, y1, x2 := rect[0], rect[1], rect[2]
		rg[x1][y1] = getRg(x2, y1)
	}

	slices.SortFunc(rects, func(a, b []int) int {
		return b[1] - a[1]
	})

	for _, rect := range rects {
		x1, y1, y2 := rect[0], rect[1], rect[3]
		up[x1][y1] = getUp(x1, y2)
	}

	play := func(x1, y1, x2, y2 int) []int {
		var res []int
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				if fp[i][j] != -1 {
					res = append(res, fp[i][j]+1)
				}
			}
		}
		slices.Sort(res)

		return res
	}

	for _, rect := range rects {
		x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
		w := max(x2-x1, y2-y1)
		// 至少要这么大
		for w <= x2 && w <= y2 {
			// sz :=
			// 这里还不够, 有可能某个位置把区域外的面积给包括进来了
			// 怎么避免这种情况呢？
			x1 := x2 - w
			y1 := y2 - w

			sz2 := dp[x2][y2] - dp[x2][y1] - dp[x1][y2] + dp[x1][y1]
			if sz2 == w*w {
				t1 := getUp(x1, y1)
				t2 := getRg(x1, y1)
				if t1 >= y2 && t2 >= x2 {
					return play(x1+1, y1+1, x2, y2)
				}
			}
			if sz2 < w*w {
				break
			}

			w++
		}
	}

	return nil
}
