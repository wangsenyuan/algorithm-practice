package p3887

func numberOfEdgesAdded(n int, edges [][]int) int {
	set := NewDSU(n)

	adj := make([][]int, n)

	var res int
	var todo []int
	for i, e := range edges {
		u, v := e[0], e[1]
		if set.Union(u, v) {
			res++
			adj[u] = append(adj[u], i)
			adj[v] = append(adj[v], i)
		} else {
			todo = append(todo, i)
		}
	}
	color := make([]int, n)
	for i := range n {
		color[i] = -1
	}

	que := make([]int, n)
	bfs := func(u int) {
		var head, tail int
		color[u] = 0
		que[head] = u
		head++
		for tail < head {
			u := que[tail]
			tail++
			for _, ei := range adj[u] {
				v := edges[ei][1] ^ edges[ei][0] ^ u
				w := edges[ei][2]
				if color[v] == -1 {
					color[v] = w ^ color[u]
					que[head] = v
					head++
				}
			}
		}
	}

	for i, v := range color {
		if v == -1 {
			bfs(i)
		}
	}

	for _, i := range todo {
		u, v, w := edges[i][0], edges[i][1], edges[i][2]
		if w == color[u]^color[v] {
			res++
		}
	}
	return res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
