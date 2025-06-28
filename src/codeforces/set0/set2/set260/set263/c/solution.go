package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math/bits"
	"os"
	"reflect"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := process(reader)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) (n int, edges [][]int, res []int) {
	n = readNum(reader)
	edges = make([][]int, n*2)
	for i := range len(edges) {
		edges[i] = readNNums(reader, 2)
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) []int {
	if n <= 6 {
		return bruteForce(n, edges)
	}
	// n > 6, 这样子，前后之间的连线才能断开
	g := NewGraph(n+1, 4*n)
	deg := make([]int, n+1)

	var arr []int

	for i, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
		deg[u]++
		deg[v]++
		if u == 1 {
			arr = append(arr, v)
		} else if v == 1 {
			arr = append(arr, u)
		}
	}
	if len(arr) != 4 {
		return nil
	}

	que := make([]int, n)
	marked := make([]bool, 2*n)
	d := make([]int, n+1)

	disconnectFrom := func(u int, from []int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if slices.Contains(from, v) {
				marked[w] = true
				d[u]--
			}
		}
	}

	bfs := func(a []int, b []int) bool {
		copy(d, deg)
		clear(marked)
		disconnectFrom(1, b)
		for _, u := range a {
			disconnectFrom(u, b)
		}

		var head, tail int
		que[head] = 1
		head++
		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := g.val[i]
				if !marked[w] {
					marked[w] = true
					d[v]--
					if d[v] <= 2 {
						que[head] = v
						head++
					}
				}
			}
		}
		return head == n
	}

	for state := 3; state < 16; state++ {
		w := bits.OnesCount(uint(state))
		if w == 2 {
			var a []int
			var b []int
			for i := range 4 {
				if (state>>i)&1 == 1 {
					a = append(a, arr[i])
				} else {
					b = append(b, arr[i])
				}
			}
			if bfs(a, b) {
				return que
			}
		}
	}

	return nil
}

func bruteForce(n int, edges [][]int) []int {
	// n <= 6
	fn := func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	}

	for _, e := range edges {
		e[0], e[1] = min(e[0], e[1]), max(e[0], e[1])
	}

	slices.SortFunc(edges, fn)

	arr := make([]int, n)

	check := func() bool {
		var tmp [][]int
		for i := range n {
			u := arr[i]
			v := arr[(i+1)%n]
			tmp = append(tmp, []int{u, v})
			v = arr[(i+2)%n]
			tmp = append(tmp, []int{u, v})
		}
		for _, e := range tmp {
			e[0], e[1] = min(e[0], e[1]), max(e[0], e[1])
		}
		slices.SortFunc(tmp, fn)
		return reflect.DeepEqual(tmp, edges)
	}

	var dfs func(i int, flag int) bool

	dfs = func(i int, flag int) bool {
		if i == n {
			return check()
		}

		for j := range n {
			if (flag>>j)&1 == 0 {
				arr[i] = j + 1
				if dfs(i+1, flag|(1<<j)) {
					return true
				}
			}
		}
		return false
	}

	if dfs(0, 0) {
		return arr
	}
	return nil
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
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v int, val int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = val
}
