package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, best, res := drive(reader)
	fmt.Println(best)
	if best < 0 {
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func drive(reader *bufio.Reader) (n int, edges [][]int, best int, res []int) {
	n, m := readTwoNums(reader)
	edges = make([][]int, m)
	for i := range m {
		s := readString(reader)
		var u, v int
		pos := readInt([]byte(s), 0, &u) + 1
		pos = readInt([]byte(s), pos, &v) + 1
		if s[pos] == 'R' {
			edges[i] = []int{u, v, 0}
		} else {
			edges[i] = []int{u, v, 1}
		}
	}
	best, res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) (int, []int) {
	// RED is 0, BLUE is 1
	g := NewGraph(n, len(edges)*2)
	w1 := -1
	for _, cur := range edges {
		u, v, w := cur[0]-1, cur[1]-1, cur[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
		if w1 == -1 {
			w1 = w
		} else if w1 != w {
			w1 = -2
		}
	}

	if w1 >= 0 {
		// 已经一个颜色了
		return 0, nil
	}

	color := make([]int, n)
	for i := range n {
		color[i] = -1
	}

	buf := make([][]int, 2)

	var dfs func(u int, c int, x int) bool
	dfs = func(u int, c int, x int) bool {
		if color[u] != -1 {
			return color[u] == c
		}
		color[u] = c
		buf[c] = append(buf[c], u+1)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			nc := g.val[i] ^ c ^ x
			if !dfs(v, nc, x) {
				return false
			}
		}
		return true
	}

	var res []int

	play := func(x int) {
		for i := range n {
			color[i] = -1
		}
		var arr []int
		for u := range n {
			if color[u] == -1 {
				clear(buf)
				if !dfs(u, 0, x) {
					return
				}
				if len(buf[0]) <= len(buf[1]) {
					arr = append(arr, buf[0]...)
				} else {
					arr = append(arr, buf[1]...)
				}
			}
		}
		if len(res) == 0 || len(arr) < len(res) {
			res = arr
		}
	}

	play(0)
	play(1)

	if len(res) == 0 {
		return -1, nil
	}
	return len(res), res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
