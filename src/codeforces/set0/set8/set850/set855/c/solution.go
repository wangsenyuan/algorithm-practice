package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, k, x int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	fmt.Fscan(reader, &k, &x)
	return solve(n, m, k, x, edges)
}

const mod = 1000000007

func add(nums ...int) int {
	var res int
	for _, num := range nums {
		res += num
		if res >= mod {
			res -= mod
		}
	}
	return res
}

func mul(a, b int) int {
	return (a * b) % mod
}

func solve(n int, m int, k int, x int, edges [][]int) int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	// dp[u][i][0/1/2] 表示u节点使用了i个k节点，且u节点的状态是0/1/2
	dp := make([][][3]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([][3]int, x+1)
	}

	var fp [3][12]int
	var gp [3][12]int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		leaf := true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				leaf = false
				dfs(u, v)
			}
		}
		if leaf {
			dp[u][0][0] = k - 1
			dp[u][1][1] = 1
			dp[u][0][2] = m - k
			return
		}

		for i := range 3 {
			for j := range x + 1 {
				fp[i][j] = 0
				gp[i][j] = 0
			}
			fp[i][0] = 1
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				for j := range 3 {
					for l := range x + 1 {
						for r := range x - l + 1 {
							// l + r <= x
							switch j {
							case 0:
								tmp := add(dp[v][r][0], dp[v][r][1], dp[v][r][2])
								gp[j][l+r] = add(gp[j][l+r], mul(fp[j][l], tmp))
							case 1:
								gp[j][l+r] = add(gp[j][l+r], mul(fp[j][l], dp[v][r][0]))
							default:
								// j == 2
								tmp := add(dp[v][r][0], dp[v][r][2])
								gp[j][l+r] = add(gp[j][l+r], mul(fp[j][l], tmp))
							}
						}
					}
				}

				for j := range 3 {
					for l := range x + 1 {
						fp[j][l] = gp[j][l]
						gp[j][l] = 0
					}
				}
			}
		}

		for l := range x + 1 {
			dp[u][l][0] = mul(fp[0][l], k-1)
			if l > 0 {
				dp[u][l][1] = fp[1][l-1]
			}
			dp[u][l][2] = mul(fp[2][l], m-k)
		}
	}

	dfs(-1, 0)

	var res int

	for j := range 3 {
		for l := range x + 1 {
			res = add(res, dp[0][l][j])
		}
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
