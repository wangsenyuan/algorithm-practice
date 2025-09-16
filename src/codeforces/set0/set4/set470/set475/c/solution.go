package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) int {

	// 然后找出所有的X区域
	n := len(a)
	m := len(a[0])
	next := make([][][2]int, n)
	for i := range n {
		next[i] = make([][2]int, m)
	}
	// next[i][j][0] 最右边不是X的地方
	// next[i][j][1] 最下边不是X的地方

	for i := range n {
		row := m
		var cnt int
		for j := m - 1; j >= 0; j-- {
			if a[i][j] == '.' {
				row = j
				if j+1 < m && a[i][j+1] == '.' {
					next[i][j][0] = next[i][j+1][0]
				} else {
					next[i][j][0] = j + 1
				}
			} else {
				// 同一行的不能被断开
				if j+1 < m && a[i][j+1] == '.' && cnt > 0 {
					return -1
				}
				cnt++
				next[i][j][0] = row
			}
		}
	}

	for j := range m {
		col := n
		var cnt int
		for i := n - 1; i >= 0; i-- {
			if a[i][j] == '.' {
				col = i
				if i+1 < n && a[i+1][j] == '.' {
					next[i][j][1] = next[i+1][j][1]
				} else {
					next[i][j][1] = i + 1
				}
			} else {
				if i+1 < n && a[i+1][j] == '.' && cnt > 0 {
					return -1
				}
				cnt++
				next[i][j][1] = col
			}
		}
	}

	find1 := func(r0 int, c0 int, h int, w int) int {
		if r0+h == n || next[r0+h][0][0] == m {
			// 只有一个区域
			return 1
		}

		r1 := r0 + h
		c1 := next[r1][0][0]

		if c1 < c0 || c0+w <= c1 {
			// 没有连通起来
			return -1
		}

		return c0 + w - c1
	}

	find2 := func(r0 int, c0 int, h int, w int) int {

		if c0+w == m || next[0][c0+w][1] == n {
			return 1
		}
		c1 := c0 + w
		r1 := next[0][c1][1]
		if r1 < r0 || r0+h <= r1 {
			return -1
		}
		return r0 + h - r1
	}

	sum := make([][]int, n+1)
	for i := range n + 1 {
		sum[i] = make([]int, m+1)
	}

	play := func(r0 int, c0 int, h int, w int) bool {
		for i := range sum {
			clear(sum[i])
		}

		// 2维差分
		for r0+h <= n || c0+w <= m {
			// 出现了必须同时往两个方向运行的情况
			r1 := next[r0][c0][1] - h
			c1 := next[r0][c0][0] - w
			if r1 < r0 || c1 < c0 || r1 > r0+h && c1 > c0+w {
				// 这个状态不应该出现
				return false
			}
			sum[r0][c0]++
			if r1 == r0 && c1 == c0 {
				sum[r0][c0+w]--
				sum[r0+h][c0]--
				sum[r0+h][c0+w]--
				break
			}
			if r0 < r1 {
				// 往下运动
				sum[r1+h][c0]--
				sum[r0][c0+w]--
				sum[r1+h][c0+w]++
				r0 = r1
			} else {
				// 往右运动
				sum[r0+h][c0]--
				sum[r0][c1+w]--
				sum[r0+h][c1+w]++
				c0 = c1
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i > 0 {
					sum[i][j] += sum[i-1][j]
				}
				if j > 0 {
					sum[i][j] += sum[i][j-1]
				}
				if i > 0 && j > 0 {
					sum[i][j] -= sum[i-1][j-1]
				}
				if sum[i][j] > 0 && a[i][j] == '.' {
					return false
				}
				if sum[i][j] == 0 && a[i][j] == 'X' {
					return false
				}
			}
		}

		return true
	}

	rect := findFirstRegion(a)
	// x = h0, or y = w0
	r0, c0 := rect[0], rect[1]
	h0, w0 := rect[2], rect[3]

	best := n*m + 1
	// 如果 x = h0, 找到最小的y <= w0,
	w1 := find1(r0, c0, h0, w0)
	if w1 > 0 && play(r0, c0, h0, w1) {
		best = h0 * w1
	}
	h1 := find2(r0, c0, h0, w0)
	if h1 > 0 && play(r0, c0, h1, w0) {
		best = min(best, h1*w0)
	}
	if best > n*m {
		return -1
	}

	return best
}

func findFirstRegion(a []string) []int {
	for i, row := range a {
		for j := range len(row) {
			if row[j] == 'X' {
				// 这个左边不能有
				k := j
				for k < len(row) && row[k] == 'X' {
					k++
				}
				w := k - j
				k = i
				for k < len(a) && a[k][j] == 'X' {
					k++
				}
				h := k - i
				return []int{i, j, h, w}
			}
		}
	}

	return nil
}
