package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, s := range res {
		buf.WriteString(s)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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
	bs, _ := reader.ReadBytes('\n')
	return strings.TrimSpace(string(bs))
}

func process(reader *bufio.Reader) []string {
	m := readNum(reader)
	friendships := make([]string, m)
	for i := range m {
		friendships[i] = readString(reader)
	}
	return solve(friendships)
}

const inf = 1 << 60

func solve(friendships []string) []string {
	id := make(map[string]int)
	var arr []string

	get := func(name string) int {
		if v, ok := id[name]; ok {
			return v
		}
		arr = append(arr, name)
		id[name] = len(id)
		return id[name]
	}

	m := len(friendships)
	edges := make([][]int, m)
	for i, cur := range friendships {
		ss := strings.Split(cur, " ")
		u := get(ss[0])
		v := get(ss[1])
		edges[i] = []int{u, v}
	}

	n := len(id)

	g := NewGraph(n, m*2)
	for _, cur := range edges {
		u, v := cur[0], cur[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	cnt := make([]int, n)
	res := make([]string, n)

	for p := 0; p < n; p++ {
		clear(cnt)
		cnt[p] = -inf
		for i := g.nodes[p]; i > 0; i = g.next[i] {
			u := g.to[i]
			// 直接相连的
			cnt[u] = -inf
			for j := g.nodes[u]; j > 0; j = g.next[j] {
				v := g.to[j]
				cnt[v]++
			}
		}
		x := slices.Max(cnt)
		var tmp int
		if x >= 0 {
			for i := 0; i < n; i++ {
				if cnt[i] == x {
					tmp++
				}
			}
		}

		res[p] = fmt.Sprintf("%s %d", arr[p], tmp)
	}

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
