package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	p := readNNums(reader, n-1)
	return solve(n, p)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, p []int) int {
	g := NewGraph(n, n)
	for i := 1; i < n; i++ {
		g.AddEdge(p[i-1]-1, i)
	}

	cnt := make([]int, n)

	var dfs func(u int, d int)
	dfs = func(u int, d int) {
		cnt[d]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v, d+1)
		}
	}
	dfs(0, 0)
	// 只有根节点的时候
	res := 1
	prod := 1
	for i := 1; i < n && cnt[i] > 0; i++ {
		res = add(res, mul(prod, cnt[i]))
		prod = mul(prod, cnt[i]-1)
	}
	return res
}

func solve1(n int, p []int) int {
	g := NewGraph(n, n)
	for i := 1; i < n; i++ {
		g.AddEdge(p[i-1]-1, i)
	}

	dp := make([]int, n)
	dp[0] = 1
	que := make([]int, n)
	var head, tail int
	que[head] = 0
	head++

	var res int
	for tail < head {
		// 这个是同一层的
		var sum int
		for i := tail; i < head; i++ {
			sum = add(sum, dp[que[i]])
		}
		res = add(res, sum)
		tmp := head
		for i := tail; i < tmp; i++ {
			u := que[i]
			for j := g.nodes[u]; j > 0; j = g.next[j] {
				v := g.to[j]
				if u == 0 {
					dp[v] = 1
				} else {
					dp[v] = sub(sum, dp[u])
				}
				que[head] = v
				head++
			}
		}
		tail = tmp
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
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
