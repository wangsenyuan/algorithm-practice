package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
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

func drive(reader *bufio.Reader) int {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

var dd = []int{-1, 0, 1, 0, -1}
var dir = "v<^>"

func solve(a []string) int {
	n := len(a)
	m := len(a[0])
	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, m)
		for j := range m {
			dist[i][j] = -1
		}
	}
	que := make([]int, n*m)
	var head, tail int
	for i := range n {
		for j := range m {
			if a[i][j] == '#' {
				dist[i][j] = 0
				que[head] = i*m + j
				head++
			}
		}
	}

	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++

		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && dist[x][y] < 0 && dir[i] == a[x][y] {
				// 还必须能从(x, y)到达位置(r, c)
				dist[x][y] = dist[r][c] + 1
				que[head] = x*m + y
				head++
			}
		}
	}

	var best int

	for i := range n {
		for j := range m {
			if dist[i][j] < 0 {
				// 这里肯定有个环
				return -1
			}
			best = max(best, dist[i][j])
		}
	}

	if best == 0 {
		return 0
	}

	// 即使有多个best，也必须保证它们不会撞在一起
	head = 0
	tail = 0
	vis := make([][]int, n)
	for i := range n {
		vis[i] = make([]int, m)
	}

	for i := range n {
		for j := range m {
			if dist[i][j] == best {
				que[head] = i*m + j
				head++
				vis[i][j] = 1
			}
		}
	}

	var cnt int

	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++
		if dist[r][c] == 0 {
			// #没有关系
			cnt++
			continue
		}
		if vis[r][c] > 1 {
			// 撞在一起了
			continue
		}
		switch a[r][c] {
		case '>':
			vis[r][c+1]++
			que[head] = r*m + c + 1
			head++
		case '<':
			vis[r][c-1]++
			que[head] = r*m + c - 1
			head++
		case 'v':
			vis[r+1][c]++
			que[head] = (r+1)*m + c
			head++
		default:
			vis[r-1][c]++
			que[head] = (r-1)*m + c
			head++
		}
	}

	if cnt >= 2 {
		return 2 * best
	}

	return 2*best - 1
}
