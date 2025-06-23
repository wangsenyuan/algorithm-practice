package p3590

const H = 20

func kthSmallest(par []int, vals []int, queries [][]int) []int {
	n := len(par)
	g := NewGraph(n, n)
	for i := 1; i < n; i++ {
		g.AddEdge(par[i], i)
	}

	qs := make([][]int, n)
	for i, q := range queries {
		u := q[0]
		qs[u] = append(qs[u], i)
	}

	ans := make([]int, len(queries))

	var dfs func(u int) *Node

	dfs = func(u int) *Node {
		res := NewNode()

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			vals[v] ^= vals[u]
			tmp := dfs(v)
			if res.Count() <= tmp.Count() {
				// 使用少的去合并多的
				res.Merge(H-1, tmp)
			} else {
				tmp.Merge(H-1, res)
				res = tmp
			}
		}

		res.Add(H-1, vals[u])
		for _, i := range qs[u] {
			k := queries[i][1]
			ans[i] = -1
			if res.Count() >= k {
				ans[i] = res.FindKth(H-1, k)
			}
		}

		return res
	}

	dfs(0)

	return ans
}

type Node struct {
	children [2]*Node
	cnt      int
}

func NewNode() *Node {
	return new(Node)
}

func (node *Node) Count() int {
	if node == nil {
		return 0
	}
	return node.cnt
}

func (node *Node) Add(i int, v int) {
	if i < 0 {
		if node.cnt == 0 {
			node.cnt++
		}
		return
	}
	x := (v >> i) & 1
	if node.children[x] == nil {
		node.children[x] = NewNode()
	}
	node.children[x].Add(i-1, v)
	node.cnt = node.children[0].Count() + node.children[1].Count()
}

func (node *Node) Merge(i int, other *Node) {
	if i < 0 || other == nil {
		return
	}

	if node.children[0] == nil {
		node.children[0] = other.children[0]
	} else {
		node.children[0].Merge(i-1, other.children[0])
	}
	if node.children[1] == nil {
		node.children[1] = other.children[1]
	} else {
		node.children[1].Merge(i-1, other.children[1])
	}
	node.cnt = node.children[0].Count() + node.children[1].Count()
}

func (node *Node) FindKth(i int, k int) int {
	if i < 0 || node == nil {
		return 0
	}
	if node.children[0].Count() >= k {
		return node.children[0].FindKth(i-1, k)
	}
	k -= node.children[0].Count()
	return (1 << i) | node.children[1].FindKth(i-1, k)
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
