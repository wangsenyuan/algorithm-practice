package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer

	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}

	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) [][]int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	f := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &f[i])
	}
	w := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &w[i])
	}
	return solve(f, w, k)
}

func solve(f []int, w []int, k int) [][]int {
	// f[i] = i 也有可能存在
	n := len(f)
	deg := make([]int, n)
	for i := range n {
		deg[f[i]]++
	}

	que := make([]int, n)

	belong := make([]int, n)

	var head, tail int
	for i := range n {
		belong[i] = -1

		if deg[i] == 0 {
			que[head] = i
			head++
		}
	}
	adj := make([][]int, n)

	for tail < head {
		u := que[tail]
		tail++
		belong[u] = u
		// f[u] = u 不可能出现
		v := f[u]
		adj[v] = append(adj[v], u)

		deg[v]--
		if deg[v] == 0 {
			que[head] = v
			head++
		}
	}
	// 剩下的都是在圈上
	var cycle [][]int

	for i := range n {
		if belong[i] == -1 {
			var cur []int
			j := i
			for belong[j] == -1 {
				cur = append(cur, j)
				belong[j] = len(cycle)
				j = f[j]
			}
			cycle = append(cycle, cur)
		}
	}

	tr1 := NewSegTree(2*n, 0, plus)
	tr2 := NewSegTree(2*n, inf, func(a, b int) int {
		return min(a, b)
	})

	pos := make([]int, n)
	for i := range n {
		// -1 not in cycle
		pos[i] = -1
	}
	var begin, end int
	for i, cur := range cycle {
		if i > 0 {
			begin = end + 1
		}
		end = begin + 2*len(cur) - 1
		for j, u := range cur {
			pos[u] = j + begin
			tr1.Update(pos[u], w[u])
			tr1.Update(pos[u]+len(cur), w[u])
			tr2.Update(pos[u], w[u])
			tr2.Update(pos[u]+len(cur), w[u])
		}
	}

	fa := make([][]int, n)
	h := bits.Len(uint(n)) + 1
	dep := make([]int, n)
	dp1 := make([][]int, n)
	dp2 := make([][]int, n)

	for i := range n {
		fa[i] = make([]int, h)
		dp1[i] = make([]int, h)
		dp2[i] = make([]int, h)
		for j := range h {
			fa[i][j] = -1
			dp1[i][j] = 0
			dp2[i][j] = inf
		}
	}

	var dfs func(u int)
	dfs = func(u int) {
		for j := 1; j < h; j++ {
			p := fa[u][j-1]
			fa[u][j] = fa[p][j-1]
			dp1[u][j] = dp1[u][j-1] + dp1[p][j-1]
			dp2[u][j] = min(dp2[u][j-1], dp2[p][j-1])
		}
		for _, v := range adj[u] {
			dep[v] = dep[u] + 1
			fa[v][0] = u
			dp1[v][0] = w[v]
			dp2[v][0] = w[v]
			dfs(v)
		}
	}

	for u := range n {
		// 这是在圈上的点
		if pos[u] >= 0 {
			fa[u][0] = u
			dp1[u][0] = 0
			dp2[u][0] = inf
			dfs(u)
		}
	}

	getPath := func(u int, d int) []int {
		res := []int{0, inf, -1}
		for i := h - 1; i >= 0; i-- {
			if (d>>i)&1 == 1 {
				res[0] += dp1[u][i]
				res[1] = min(res[1], dp2[u][i])
				u = fa[u][i]
			}
		}
		res[2] = u
		return res
	}

	find := func(u int) []int {
		k1 := k
		res := []int{0, inf}
		if pos[u] < 0 {
			if dep[u] >= k1 {
				return getPath(u, k1)
			}
			// dep[u] < k
			tmp := getPath(u, dep[u])
			copy(res, tmp)
			k1 -= dep[u]
			u = tmp[2]
		}
		// k1 > 0
		id := belong[u]
		m := len(cycle[id])
		x, y := k1/m, k1%m
		if x > 0 {
			res[0] += x * tr1.Get(pos[u], pos[u]+m)
			res[1] = min(res[1], tr2.Get(pos[u], pos[u]+m))
		}
		if y > 0 {
			res[0] += tr1.Get(pos[u], pos[u]+y)
			res[1] = min(res[1], tr2.Get(pos[u], pos[u]+y))
		}
		return res
	}

	ans := make([][]int, n)

	for u := range n {
		ans[u] = find(u)
	}

	return ans
}

const inf = 1 << 60

func plus(a, b int) int {
	return a + b
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
