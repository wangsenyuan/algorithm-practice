package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res, _ := process(reader)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d ", x))
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Println(buf.String())
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

func process(reader *bufio.Reader) (res []int, songs [][]int) {
	n := readNum(reader)
	songs = make([][]int, n)
	for i := range n {
		songs[i] = readNNums(reader, 2)
	}
	res = solve(songs)
	return
}

func solve(songes [][]int) []int {
	var xs []int
	var ys []int
	for _, cur := range songes {
		xs = append(xs, cur[0])
		ys = append(ys, cur[1])
	}
	xs = sortAndUnique(xs)
	ys = sortAndUnique(ys)

	n := len(xs) + len(ys)
	m := len(songes)

	type edge struct {
		u  int
		v  int
		id int
	}

	edges := make([]edge, m)

	marked := make([]bool, m)
	g := NewGraph(n, 2*m)

	deg := make([]int, n)

	for i, cur := range songes {
		u := sort.SearchInts(xs, cur[0])
		v := sort.SearchInts(ys, cur[1]) + len(xs)
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
		edges[i] = edge{u, v, i}
		deg[u]++
		deg[v]++
	}

	var odd int
	var start int
	for i := range n {
		if deg[i]%2 == 1 {
			odd++
			start = i
		}
	}
	if odd > 2 {
		return nil
	}

	pos := make([]int, n)
	for i := range n {
		pos[i] = g.nodes[i]
	}

	var arr []int
	var dfs func(u int)
	dfs = func(u int) {
		for pos[u] > 0 {
			v := g.to[pos[u]]
			w := g.val[pos[u]]
			pos[u] = g.next[pos[u]]
			if !marked[w] {
				marked[w] = true
				dfs(v)
			}
		}
		arr = append(arr, u)
	}

	dfs(start)

	if len(arr) != m+1 {
		return nil
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return cmp.Or(a.u-b.u, a.v-b.v)
	})

	res := make([]int, m)

	clear(marked)

	for i := range m {
		u := arr[i]
		v := arr[i+1]
		if u >= len(xs) {
			// v < len(xs)
			u, v = v, u
		}
		j := sort.Search(m, func(j int) bool {
			if edges[j].u > u || edges[j].u == u && edges[j].v >= v {
				return true
			}
			return false
		})

		if j == m || marked[j] || edges[j].u != u || edges[j].v != v {
			return nil
		}
		marked[j] = true
		res[i] = edges[j].id + 1
	}
	return res
}

func sortAndUnique(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	sort.Ints(res)
	var n int
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] != res[i-1] {
			res[n] = res[i-1]
			n++
		}
	}
	return res[:n]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.val = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
