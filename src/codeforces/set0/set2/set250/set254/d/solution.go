package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Fprintln(w, "-1")
		return
	}
	fmt.Fprintln(w, res[0], res[1], res[2], res[3])
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

func drive(reader *bufio.Reader) (d int, a []string, res []int) {
	n, _, d := readThreeNums(reader)
	a = make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	res = solve(d, a)
	return
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 30

var dd = []int{-1, 0, 1, 0, -1}

func solve(d int, a []string) []int {
	n := len(a)
	m := len(a[0])
	var rats []pair

	for i := range n {
		for j := range m {
			if a[i][j] == 'R' {
				rats = append(rats, pair{i, j})
			}
		}
	}

	que := make([]int, n*m)

	bfs := func(x int, y int) [][]int {
		dist := make([][]int, n)
		for i := range n {
			dist[i] = make([]int, m)
			for j := range m {
				dist[i][j] = inf
			}
		}
		dist[x][y] = 0
		var head, tail int
		que[head] = x*m + y
		head++
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++

			for i := range 4 {
				nr, nc := r+dd[i], c+dd[i+1]
				if nr >= 0 && nr < n && nc >= 0 && nc < m && dist[nr][nc] == inf && a[nr][nc] != 'X' {
					dist[nr][nc] = dist[r][c] + 1
					que[head] = nr*m + nc
					head++
				}
			}
		}
		return dist
	}

	d1 := bfs(rats[0].first, rats[0].second)

	rat1 := 0
	for i, cur := range rats {
		if d1[cur.first][cur.second] > d1[rats[rat1].first][rats[rat1].second] {
			rat1 = i
		}
	}

	x1, y1 := rats[rat1].first, rats[rat1].second

	d2 := bfs(x1, y1)

	getCandidates := func(dist [][]int) [][]int {
		var res [][]int
		for i := range n {
			for j := range m {
				if dist[i][j] <= d {
					res = append(res, []int{i, j})
				}
			}
		}
		return res
	}

	p1 := getCandidates(d2)

	rat2 := 0
	for i, cur := range rats {
		if d2[cur.first][cur.second] > d2[rats[rat2].first][rats[rat2].second] {
			rat2 = i
		}
	}

	x2, y2 := rats[rat2].first, rats[rat2].second
	d3 := bfs(x2, y2)

	for _, cur := range rats {
		x, y := cur.first, cur.second
		if d2[x][y] > 2*d && d3[x][y] > 2*d {
			// 存在第3个很远的rat
			return nil
		}
	}

	p2 := getCandidates(d3)

	marked := make([][]int, n)
	for i := range n {
		marked[i] = make([]int, m)
		for j := range m {
			d1[i][j] = inf
		}
	}

	que1 := make([]int, n*m)
	que2 := make([]int, n*m)

	bfs2 := func(x int, y int, w int, que []int) int {
		var head, tail int
		que[head] = x*m + y
		head++
		marked[x][y] |= w
		d1[x][y] = 0
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++
			if d1[r][c] == d {
				continue
			}
			for i := range 4 {
				nr, nc := r+dd[i], c+dd[i+1]
				if nr >= 0 && nr < n && nc >= 0 && nc < m && a[nr][nc] != 'X' && d1[nr][nc] == inf {
					d1[nr][nc] = d1[r][c] + 1
					marked[nr][nc] |= w
					que[head] = nr*m + nc
					head++
				}
			}
		}

		for i := range head {
			r, c := que[i]/m, que[i]%m
			d1[r][c] = inf
		}
		return head
	}

	check := func() bool {
		for _, cur := range rats {
			if marked[cur.first][cur.second] == 0 {
				return false
			}
		}
		return true
	}

	for _, f := range p1 {
		k1 := bfs2(f[0], f[1], 1, que1)
		if check() {
			// 全部被覆盖到了
			for i := range n {
				for j := range m {
					if (f[0] != i || f[1] != j) && a[i][j] != 'X' {
						return []int{f[0] + 1, f[1] + 1, i + 1, j + 1}
					}
				}
			}
		}
		for _, s := range p2 {
			if f[0] == s[0] && f[1] == s[1] {
				continue
			}
			k2 := bfs2(s[0], s[1], 2, que2)

			// 检查rats的状态
			if check() {
				return []int{f[0] + 1, f[1] + 1, s[0] + 1, s[1] + 1}
			}
			for i := range k2 {
				r, c := que2[i]/m, que2[i]%m
				marked[r][c] ^= 2
			}
		}
		for i := range k1 {
			r, c := que1[i]/m, que1[i]%m
			marked[r][c] ^= 1
		}
	}

	return nil
}
