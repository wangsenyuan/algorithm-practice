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
	n, t := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	readString(reader)
	b := make([]string, n)
	for i := range n {
		b[i] = readString(reader)
	}
	return solve(n, t, a, b)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(n int, t int, a []string, b []string) int {
	N := n*n*2 + 2
	src := N - 2
	snk := N - 1

	que := make([]int, n*n)
	var head, tail int
	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, n)
		for j := range n {
			dist[i][j] = INF
			if a[i][j] == 'Z' {
				que[head] = i*n + j
				head++
				dist[i][j] = 0
			}
		}
	}

	for tail < head {
		r, c := que[tail]/n, que[tail]%n
		tail++
		if dist[r][c] == t {
			break
		}

		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && dist[nr][nc] == INF && a[nr][nc] != 'Y' {
				dist[nr][nc] = dist[r][c] + 1
				que[head] = nr*n + nc
				head++
			}
		}
	}

	// N == 100
	g := NewGraph(N, N*N*2)

	addEdge := func(u, v, w int) {
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, 0)
	}

	dist2 := make([][]int, n)
	for i := range n {
		dist2[i] = make([]int, n)
	}
	// 为a[x][y]的科学家寻找安全仓
	play := func(x int, y int) {
		head, tail = 0, 0
		for i := range n {
			for j := range n {
				dist2[i][j] = INF
			}
		}
		dist2[x][y] = 0
		que[head] = x*n + y
		head++

		for tail < head {
			r, c := que[tail]/n, que[tail]%n
			tail++

			if b[r][c] >= '0' && b[r][c] <= '9' {
				// 安全仓
				addEdge(x*n+y, n*n+r*n+c, INF)
			}
			if dist2[r][c] == t || dist2[r][c] == dist[r][c] {
				// 被赶上了
				continue
			}

			for i := range 4 {
				nr, nc := r+dd[i], c+dd[i+1]
				if nr >= 0 && nr < n && nc >= 0 && nc < n &&
					dist2[nr][nc] == INF && b[nr][nc] >= '0' && b[nr][nc] <= '9' &&
					dist[nr][nc] >= dist2[r][c]+1 {
					// 可以安全到达
					dist2[nr][nc] = dist2[r][c] + 1
					que[head] = nr*n + nc
					head++
				}
			}
		}
	}

	for i := range n {
		for j := range n {
			if a[i][j] >= '0' && a[i][j] <= '9' {
				play(i, j)
				addEdge(src, i*n+j, int(a[i][j]-'0'))
			}
			if b[i][j] >= '0' && b[i][j] <= '9' {
				addEdge(n*n+i*n+j, snk, int(b[i][j]-'0'))
			}
		}
	}

	return dinic(src, snk, N, g)
}

const INF = 1 << 60

func dinic(src, snk int, n int, g *Graph) int {
	level := make([]int, n)
	all_minus_one := make([]int, n)
	for i := 0; i < n; i++ {
		all_minus_one[i] = -1
	}

	que := make([]int, n)

	bfs := func() bool {
		var front, end int
		copy(level, all_minus_one)
		level[src] = 0
		que[end] = src
		end++
		for front < end {
			u := que[front]
			front++
			for i := g.node[u]; i > 0; i = g.next[i] {
				if g.limit[i] > g.flow[i] && level[g.to[i]] == -1 {
					v := g.to[i]
					level[v] = level[u] + 1
					que[end] = v
					end++
				}
			}
		}
		return level[snk] > 0
	}

	pos := make([]int, n)

	var dfs func(u int, flow int) int
	dfs = func(u int, flow int) int {
		if flow == 0 {
			return 0
		}
		if u == snk {
			return flow
		}

		for pos[u] > 0 {
			i := pos[u]
			v := g.to[i]
			if level[v] == level[u]+1 && g.flow[i] < g.limit[i] {
				tr := dfs(v, min(flow, g.limit[i]-g.flow[i]))
				if tr > 0 {
					g.flow[i] += tr
					g.flow[i^1] -= tr
					return tr
				}
			}

			pos[u] = g.next[i]
		}
		return 0
	}
	var flow int
	for bfs() {
		for i := 0; i < n; i++ {
			pos[i] = g.node[i]
		}
		for {
			cur := dfs(src, INF)
			if cur == 0 {
				break
			}
			flow += cur
		}
	}
	return flow
}

type Graph struct {
	node  []int
	next  []int
	to    []int
	flow  []int
	limit []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.flow = make([]int, e+3)
	g.limit = make([]int, e+3)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.limit[g.cur] = w
	g.flow[g.cur] = 0
}
