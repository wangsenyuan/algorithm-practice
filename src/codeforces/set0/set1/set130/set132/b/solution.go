package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(m, a)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(tot int, a []string) int {
	n := len(a)
	m := len(a[0])

	// T/R, R/T, R/D, D/R, D/L, L/D, L/T, T/L
	cur := make([]int, 8)

	updateTR := func(x int, y int) {
		if cur[0] == -1 {
			cur[0] = x*m + y
			return
		}
		r, c := cur[0]/m, cur[0]%m
		if x < r || x == r && c < y {
			cur[0] = x*m + y
		}
	}

	updateRT := func(x int, y int) {
		if cur[1] == -1 {
			cur[1] = x*m + y
			return
		}
		r, c := cur[1]/m, cur[1]%m
		if c < y || c == y && x < r {
			cur[1] = x*m + y
		}
	}

	updateRD := func(x int, y int) {
		if cur[2] == -1 {
			cur[2] = x*m + y
			return
		}
		r, c := cur[2]/m, cur[2]%m
		if c < y || c == y && r < x {
			cur[2] = x*m + y
		}
	}

	updateDR := func(x int, y int) {
		if cur[3] == -1 {
			cur[3] = x*m + y
			return
		}
		r, c := cur[3]/m, cur[3]%m
		if r < x || r == x && c < y {
			cur[3] = x*m + y
		}
	}

	updateDL := func(x int, y int) {
		if cur[4] == -1 {
			cur[4] = x*m + y
			return
		}
		r, c := cur[4]/m, cur[4]%m
		if r < x || r == x && y < c {
			cur[4] = x*m + y
		}
	}

	updateLD := func(x int, y int) {
		if cur[5] == -1 {
			cur[5] = x*m + y
			return
		}
		r, c := cur[5]/m, cur[5]%m
		if y < c || y == c && r < x {
			cur[5] = x*m + y
		}
	}

	updateLT := func(x int, y int) {
		if cur[6] == -1 {
			cur[6] = x*m + y
			return
		}
		r, c := cur[6]/m, cur[6]%m
		if y < c || y == c && x < r {
			cur[6] = x*m + y
		}
	}

	updateTL := func(x int, y int) {
		if cur[7] == -1 {
			cur[7] = x*m + y
			return
		}
		r, c := cur[7]/m, cur[7]%m
		if x < r || x == r && y < c {
			cur[7] = x*m + y
		}
	}

	updates := []func(int, int){
		updateTR, updateRT, updateRD, updateDR, updateDL, updateLD, updateLT, updateTL,
	}

	var dfs func(x int, y int, num byte)

	buf := make([][]byte, n)

	belong := make([][]int, n)

	for i := range buf {
		belong[i] = make([]int, m)
		buf[i] = []byte(a[i])
		for j := range m {
			belong[i][j] = -1
		}
	}

	var blocks [][]int

	dfs = func(x int, y int, num byte) {
		for _, fn := range updates {
			fn(x, y)
		}
		belong[x][y] = len(blocks)

		buf[x][y] = '0'

		for i := range 4 {
			r, c := x+dd[i], y+dd[i+1]
			if r >= 0 && r < n && c >= 0 && c < m && buf[r][c] == num {
				dfs(r, c, num)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if buf[i][j] != '0' {
				for k := range cur {
					cur[k] = -1
				}
				dfs(i, j, buf[i][j])
				blocks = append(blocks, slices.Clone(cur))
			}
		}
	}

	k := len(blocks)

	dp := make([][][]int, k)
	for i := range k {
		dp[i] = make([][]int, 4)
		for j := range 4 {
			dp[i][j] = make([]int, 2)
			for u := range 2 {
				dp[i][j][u] = -1
			}
		}
	}

	// R, D, L, T
	// L, R

	var dfs2 func(b int, d int, c int, w int, moves int, keep bool) int

	dfs2 = func(b int, d int, c int, w int, moves int, keep bool) int {
		if w == 0 {
			first := blocks[b][0]
			i, j := first/m, first%m
			return int(a[i][j] - '0')
		}

		if dp[b][d][c] != -1 && keep {
			cycle_length := moves - dp[b][d][c]
			left := (tot - moves) % cycle_length
			return dfs2(b, d, c, left, moves, false)
		}
		if keep {
			dp[b][d][c] = moves
		}

		if d == 0 {
			if c == 0 {
				x, y := blocks[b][0]/m, blocks[b][0]%m
				if y+1 < m && belong[x][y+1] >= 0 {
					return dfs2(belong[x][y+1], d, c, w-1, moves+1, keep)
				}
				return dfs2(b, d, 1, w-1, moves+1, keep)
			}
			x, y := blocks[b][3]/m, blocks[b][3]%m
			if y+1 < m && belong[x][y+1] >= 0 {
				return dfs2(belong[x][y+1], d, c, w-1, moves+1, keep)
			}
			// c == 1
			return dfs2(b, 1, 0, w-1, moves+1, keep)
		}
		if d == 1 {
			if c == 0 {
				x, y := blocks[b][2]/m, blocks[b][2]%m
				if x+1 < n && belong[x+1][y] >= 0 {
					return dfs2(belong[x+1][y], d, c, w-1, moves+1, keep)
				}
				return dfs2(b, d, 1, w-1, moves+1, keep)
			}
			// c == 1
			x, y := blocks[b][4]/m, blocks[b][4]%m
			if x+1 < n && belong[x+1][y] >= 0 {
				return dfs2(belong[x+1][y], d, c, w-1, moves+1, keep)
			}
			return dfs2(b, 2, 0, w-1, moves+1, keep)
		}
		if d == 2 {
			if c == 0 {
				x, y := blocks[b][5]/m, blocks[b][5]%m
				if y-1 >= 0 && belong[x][y-1] >= 0 {
					return dfs2(belong[x][y-1], d, c, w-1, moves+1, keep)
				}
				return dfs2(b, d, 1, w-1, moves+1, keep)
			}
			// c == 1
			x, y := blocks[b][6]/m, blocks[b][6]%m
			if y-1 >= 0 && belong[x][y-1] >= 0 {
				return dfs2(belong[x][y-1], d, c, w-1, moves+1, keep)
			}
			return dfs2(b, 3, 0, w-1, moves+1, keep)
		}
		// d == 3
		if c == 0 {
			x, y := blocks[b][7]/m, blocks[b][7]%m
			if x-1 >= 0 && belong[x-1][y] >= 0 {
				return dfs2(belong[x-1][y], d, c, w-1, moves+1, keep)
			}
			return dfs2(b, d, 1, w-1, moves+1, keep)
		}
		// c == 1
		x, y := blocks[b][1]/m, blocks[b][1]%m
		if x-1 >= 0 && belong[x-1][y] >= 0 {
			return dfs2(belong[x-1][y], d, c, w-1, moves+1, keep)
		}
		return dfs2(b, 0, 0, w-1, moves+1, keep)
	}

	return dfs2(0, 0, 0, tot, 0, true)
}
