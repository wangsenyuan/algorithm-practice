package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)[:m]
	}
	return solve(a)
}

var dd = []int{-1, 0, 1, 0, -1}
var xx = []int{3, 1, 0, 2}

func solve(a []string) int {
	s_pos := findPosition(a, 'S')
	t_pos := findPosition(a, 'T')

	n := len(a)
	m := len(a[0])
	dp := make([][][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([][]int, 4)
			// 0 from top, 1 from left, 2 from right, 3 from bottom
			for k := range 4 {
				// 0, 1, 2
				dp[i][j][k] = make([]int, 3)
				for u := range 3 {
					dp[i][j][k][u] = -1
				}
			}
		}
	}

	que := make([]int, n*m*4*3)
	var head, tail int

	push := func(i int, j int, x int, d int) {
		que[head] = i*m*12 + j*12 + x*3 + d
		head++
	}

	pop := func() (i int, j int, x int, d int) {
		cur := que[tail]
		tail++
		i = cur / (m * 12)
		j = (cur % (m * 12)) / 12
		x = (cur % 12) / 3
		d = cur % 3
		return
	}

	checkAndAdd := func(i int, j int, x int, d int, v int) {
		if d < 3 && i >= 0 && i < n && j >= 0 && j < m && a[i][j] != '#' && dp[i][j][x][d] < 0 {
			dp[i][j][x][d] = v
			push(i, j, x, d)
		}
	}

	for i := 0; i < 4; i++ {
		checkAndAdd(s_pos[0]+dd[i], s_pos[1]+dd[i+1], xx[i], 0, 1)
	}

	for tail < head {
		r, c, x0, d := pop()
		for i := 0; i < 4; i++ {
			x := xx[i]
			var nd int
			if x == x0 {
				nd = d + 1
			}
			checkAndAdd(r+dd[i], c+dd[i+1], x, nd, dp[r][c][x0][d]+1)
		}
	}

	ans := 1 << 60

	for x := range 4 {
		for d := 0; d < 3; d++ {
			if dp[t_pos[0]][t_pos[1]][x][d] > 0 {
				ans = min(ans, dp[t_pos[0]][t_pos[1]][x][d])
			}
		}
	}

	if ans == 1<<60 {
		return -1
	}

	return ans
}

func findPosition(a []string, x byte) []int {
	for i, r := range a {
		for j := range len(r) {
			if r[j] == x {
				return []int{i, j}
			}
		}
	}
	return nil
}
