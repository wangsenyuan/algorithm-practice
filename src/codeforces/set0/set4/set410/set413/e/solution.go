package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func drive(reader *bufio.Reader) []int {
	_, m := readTwoNums(reader)
	a := make([]string, 2)
	a[0] = readString(reader)
	a[1] = readString(reader)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

func solve(a []string, queries [][]int) []int {
	n := len(a[0])

	belong := make([][]int, 2)
	dist := make([][]int, 2)
	marked := make([][]bool, 2)
	for i := range 2 {
		dist[i] = make([]int, n)
		belong[i] = make([]int, n)
		marked[i] = make([]bool, n)
		for j := range n {
			dist[i][j] = -1
		}
	}

	var id int

	que := make([]int, 2*n)

	work := func(r int, c int) {
		id++
		var head, tail int
		que[head] = r*n + c
		head++
		dist[r][c] = 0
		for tail < head {
			x, y := que[tail]/n, que[tail]%n
			tail++

			if !marked[x^1][y] {
				// y列先出现的那个格子，肯定在最优路径上
				marked[x][y] = true
			}

			belong[x][y] = id
			if y+1 < n && a[x][y+1] == '.' && dist[x][y+1] < 0 {
				dist[x][y+1] = dist[x][y] + 1
				que[head] = x*n + y + 1
				head++
			}

			if a[x^1][y] == '.' && dist[x^1][y] < 0 {
				dist[x^1][y] = dist[x][y] + 1
				que[head] = (x^1)*n + y
				head++
			}
		}
	}

	for i := range n {
		for j := range 2 {
			if a[j][i] == '.' && dist[j][i] < 0 {
				work(j, i)
			}
		}
	}
	// next[i][j] 是(i, j)右边，最近的，在最优路径上的格子
	next := make([][]int, 2)
	for i := range 2 {
		next[i] = make([]int, n+1)
		next[i][n] = n
		for j := n - 1; j >= 0; j-- {
			if a[i][j] == '.' {
				next[i][j] = next[i][j+1]
				if marked[i][j] {
					next[i][j] = j
				}
			} else {
				next[i][j] = j
			}
		}
	}

	ans := make([]int, len(queries))

	getPos := func(num int) []int {
		if num < n {
			return []int{0, num}
		}
		return []int{1, num - n}
	}

	find := func(first []int, second []int) int {
		if belong[first[0]][first[1]] != belong[second[0]][second[1]] {
			return -1
		}

		if first[1] == second[1] {
			return abs(first[0] - second[0])
		}

		if second[1] <= next[first[0]][first[1]] && first[0] == second[0] {
			// 同一行，中间没有X
			return second[1] - first[1]
		}

		if marked[first[0]][first[1]] {
			return dist[second[0]][second[1]] - dist[first[0]][first[1]]
		}
		ans := dist[second[0]][second[1]] - dist[first[0]^1][first[1]] + 1
		j := next[first[0]][first[1]]
		if j < second[1] && a[first[0]][j] == '.' {
			ans -= 2
		}
		return ans
	}

	for i, cur := range queries {
		a := getPos(cur[0] - 1)
		b := getPos(cur[1] - 1)
		if a[1] > b[1] {
			a, b = b, a
		}
		ans[i] = find(a, b)
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}
