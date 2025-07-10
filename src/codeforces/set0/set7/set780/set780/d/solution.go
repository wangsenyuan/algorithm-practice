package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	teams := make([]string, n)

	for i := 0; i < n; i++ {
		teams[i] = readString(reader)
	}

	ok, res := solve(teams)

	if !ok {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func getNameByRules(teams []string) [][]string {
	n := len(teams)
	names := make([][]string, n)

	for i, cur := range teams {
		names[i] = []string{
			rule1(cur),
			rule2(cur),
		}
	}

	return names
}

func solve(teams []string) (ok bool, res []string) {
	names := getNameByRules(teams)
	n := len(teams)
	n2 := n * 2

	g := NewGraph(n2, n2*n2*4)
	r := NewGraph(n2, n2*n2*4)

	connect := func(u int, v int) {
		g.AddEdge(u, v)
		g.AddEdge(v^1, u^1)
		r.AddEdge(v, u)
		r.AddEdge(u^1, v^1)
	}

	for i := 0; i < n; i++ {
		var found bool
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			// 如果i按照规则2取名称，且j按照规则1和i的规则1名称相同时，j必须也按照规则2
			if names[i][0] == names[j][0] {
				// connect(2*i+1, 2*j+1)
				found = true
			}
			for u := 0; u < 2; u++ {
				for v := 0; v < 2; v++ {
					if names[i][u] == names[j][v] {
						// 不能有相同的名称
						connect(2*i+u, 2*j+v^1)
					}
				}
			}
		}
		if found || names[i][0] == names[i][1] {
			// 规则2
			connect(2*i, 2*i+1)
		}
	}

	vis := make([]bool, n2)
	var order []int32

	var dfs func(u int32)

	dfs = func(u int32) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		order = append(order, u)
	}

	for i := 0; i < n2; i++ {
		if !vis[i] {
			dfs(int32(i))
		}
	}
	comp := make([]int32, n2)

	var dfs2 func(u int32, cid int32)
	dfs2 = func(u int32, cid int32) {
		comp[u] = cid
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			if comp[v] == 0 {
				dfs2(v, cid)
			}
		}
	}
	var cid int32
	for i := n2 - 1; i >= 0; i-- {
		if comp[order[i]] == 0 {
			cid++
			dfs2(order[i], cid)
		}
	}
	res = make([]string, n)
	for i := 0; i < n; i++ {
		if comp[2*i] == comp[2*i+1] {
			return false, nil
		}
		if comp[2*i] > comp[2*i+1] {
			res[i] = names[i][0]
		} else {
			res[i] = names[i][1]
		}
	}

	return true, res
}

func rule1(s string) string {
	return s[:3]
}

func rule2(s string) string {
	a := s[:2]
	var i int
	for i < len(s) && s[i] != ' ' {
		i++
	}
	i++
	return a + s[i:i+1]
}

type Graph struct {
	nodes []int32
	next  []int32
	to    []int32
	cur   int32
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int32, n)
	next := make([]int32, e+3)
	to := make([]int32, e+3)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = int32(v)
}
