package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	R := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &R[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(n, R, queries)
}

func solve(n int, R []int, queries [][]int) []int {
	adj := make([][]int, n)
	var roots []int
	for i := range n {
		if R[i] == 0 {
			roots = append(roots, i)
		} else {
			adj[R[i]-1] = append(adj[R[i]-1], i)
		}
	}

	h := bits.Len(uint(n))
	fa := make([][]int, n)

	for i := range n {
		fa[i] = make([]int, h)
	}

	var layers [][]int

	in := make([]int, n)
	sz := make([]int, n)
	dep := make([]int, n)
	var timer int

	var dfs func(u int, d int)
	dfs = func(u int, d int) {
		dep[u] = d
		in[u] = timer
		timer++
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		if len(layers) == d {
			layers = append(layers, []int{})
		}

		layers[d] = append(layers[d], u)

		sz[u] = 1
		for _, v := range adj[u] {
			fa[v][0] = u
			dfs(v, d+1)
			sz[u] += sz[v]
		}
	}

	for _, r := range roots {
		fa[r][0] = r
		dfs(r, 0)
	}

	kthAncestor := func(u int, k int) int {
		for i := h - 1; i >= 0; i-- {
			if k&(1<<i) > 0 {
				u = fa[u][i]
			}
		}
		return u
	}

	find := func(v int, k int) int {
		if dep[v] < k {
			return 0
		}
		// dep[v] >= k
		p := kthAncestor(v, k)

		d := dep[v]

		l := sort.Search(len(layers[d]), func(i int) bool {
			return in[layers[d][i]] > in[p]
		})

		// 还有v后面，但是也是它p-cousin的，怎么处理？

		r := sort.Search(len(layers[d]), func(i int) bool {
			return in[layers[d][i]] >= in[p]+sz[p]
		})

		return r - l - 1
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0]-1, cur[1])
	}

	return ans
}
