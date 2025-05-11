package p3544

const inf = 1 << 60

func subtreeInversionSum(edges [][]int, nums []int, k int) int64 {
	n := len(nums)
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k)
		for j := 0; j < k; j++ {
			dp[i][j] = make([]int, 2)
			for d := 0; d < 2; d++ {
				dp[i][j][d] = -inf
			}
		}
	}

	var dfs func(p int, u int, x int, d int) int

	dfs = func(p int, u int, x int, d int) int {
		if dp[u][x][d] != -inf {
			return dp[u][x][d]
		}

		if x == 0 {
			sum := []int{nums[u], -nums[u]}
			if d == 1 {
				// d == 1 表示目前状态是反转过的
				sum[0], sum[1] = sum[1], sum[0]
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v {
					sum[0] += dfs(u, v, 0, d)
					sum[1] += dfs(u, v, 1%k, d^1)
				}
			}
			dp[u][x][d] = max(sum[0], sum[1])
		} else {
			// x > 0, 这里受限了
			sum := nums[u]
			if d == 1 {
				sum = -nums[u]
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v {
					sum += dfs(u, v, (x+1)%k, d)
				}
			}
			dp[u][x][d] = sum
		}
		return dp[u][x][d]
	}

	res := dfs(-1, 0, 0, 0)
	return int64(res)
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
