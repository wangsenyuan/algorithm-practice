package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
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
	n, _ := readTwoNums(reader)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = readString(reader)
	}
	return solve(grid)
}

type state struct {
	hr, hc int
	body   int
}

const inf = 1 << 60

func solve(grid []string) int {
	n := len(grid)
	m := len(grid[0])
	var k int

	dd := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}

	findState := func(dr int, dc int) int {
		for i := range 4 {
			if dd[i][0] == dr && dd[i][1] == dc {
				return i
			}
		}
		return -1
	}

	rev := func(s int) int {
		if s == 0 {
			return 2
		}
		if s == 2 {
			return 0
		}
		if s == 1 {
			return 3
		}
		return 1
	}

	pos := make([][]int, 10)

	ar, ac := 0, 0

	for i := range grid {
		for j := range m {
			if grid[i][j] >= '1' && grid[i][j] <= '9' {
				x := int(grid[i][j] - '0')
				k = max(k, x)
				pos[x] = []int{i, j}
			}
			if grid[i][j] == '@' {
				ar, ac = i, j
			}
		}
	}

	var snake state

	snake.hr, snake.hc = pos[1][0], pos[1][1]

	for i := 2; i <= k; i++ {
		dr, dc := pos[i][0]-pos[i-1][0], pos[i][1]-pos[i-1][1]
		s := findState(dr, dc)
		snake.body |= s << (2 * (i - 2))
	}

	mask := 1<<(2*(k-1)) - 1

	check := func(r int, c int, cur int, s int) (ok bool, next state) {
		// 当前状态，是否可以移动到位置(i, j)作为头部
		dr, dc := dd[s][0], dd[s][1]
		x, y := r+dr, c+dc
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == '#' {
			return false, state{}
		}
		// buf[i][j] 不是墙
		// 要知道蛇的形状
		pos[1][0], pos[1][1] = r, c
		for i := 2; i <= k; i++ {
			s := (cur >> (2 * (i - 2))) & 3
			r, c = r+dd[s][0], c+dd[s][1]
			pos[i][0], pos[i][1] = r, c
			if i < k && r == x && c == y {
				return false, state{}
			}
		}

		tmp := (cur << 2) & mask
		tmp |= rev(s)
		return true, state{x, y, tmp}
	}

	dist := make([][][]int, n)
	for i := range n {
		dist[i] = make([][]int, m)
		for j := range m {
			dist[i][j] = make([]int, mask+1)
			for w := range mask + 1 {
				dist[i][j][w] = -1
			}
		}
	}

	dist[snake.hr][snake.hc][snake.body] = 0

	que := []state{snake}

	var tail int
	for tail < len(que) {
		cur := que[tail]
		tail++
		if cur.hr == ar && cur.hc == ac {
			return dist[cur.hr][cur.hc][cur.body]
		}

		for i := range 4 {
			ok, next := check(cur.hr, cur.hc, cur.body, i)
			if ok && dist[next.hr][next.hc][next.body] == -1 {
				dist[next.hr][next.hc][next.body] = dist[cur.hr][cur.hc][cur.body] + 1
				que = append(que, next)
			}
		}
	}

	return -1
}
