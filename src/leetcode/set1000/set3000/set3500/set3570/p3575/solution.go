package p3575

import "slices"

const inf = 1 << 60
const mod = 1e9 + 7

func goodSubtreeSum(vals []int, par []int) int {
	// 每个子树u种，最多选9个点(1, ... 9) 最多出现一次
	// 假设对于p来说，p在u中的选择，肯定也是u的一个最优选择
	// 所以假设p目前来说选择了一个set，那么合并v（v是p的一个子节点）
	// 可以用dp？
	// dp[state] = state表示的数的一个集合的最大值
	n := len(vals)
	dp := make([][]int, n)

	mask := 1 << 10
	for i := range n {
		dp[i] = make([]int, mask)
		for j := range mask {
			dp[i][j] = -inf
		}
	}
	g := NewGraph(n, n-1)
	for i := 1; i < n; i++ {
		g.AddEdge(par[i], i)
	}

	var dfs func(u int)

	dfs = func(u int) {
		dp[u][0] = 0
		x, ok := encode(vals[u])
		if ok {
			dp[u][x] = vals[u]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			// 使用v来更新u
			// 假设s的某个子集是x， 那么 dp[u][s] = max(dp[u][s ^ x] + dp[v][x])
			for s := range mask {
				for vs := s; vs > 0; vs = (vs - 1) & s {
					dp[u][s] = max(dp[u][s], dp[u][s^vs]+dp[v][vs])
				}
			}
		}
	}

	dfs(0)

	var ans int
	for u := range n {
		tmp := slices.Max(dp[u])
		ans = (ans + tmp) % mod
	}
	return ans
}

func encode(x int) (state int, ok bool) {
	for x > 0 {
		r := x % 10
		if state&(1<<r) > 0 {
			return 0, false
		}
		state |= 1 << r
		x /= 10
	}
	return state, true
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
