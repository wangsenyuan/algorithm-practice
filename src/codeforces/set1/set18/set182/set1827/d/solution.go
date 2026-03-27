package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &p[i])
	}
	return solve(n, p)
}

func solve(n int, p []int) []int {
	adj := make([][]int, n)
	for i := range n - 1 {
		adj[p[i]-1] = append(adj[p[i]-1], i+1)
	}

	H := bits.Len(uint(n))
	fa := make([][]int, n)
	dep := make([]int, n)
	pos := make([]int, n)
	sz := make([]int, n)

	var timer int

	var dfs func(u int)

	dfs = func(u int) {
		pos[u] = timer
		timer++
		sz[u] = 1
		fa[u] = make([]int, H)
		if u > 0 {
			fa[u][0] = p[u-1] - 1
			for i := 1; i < H; i++ {
				fa[u][i] = fa[fa[u][i-1]][i-1]
			}
		}

		for _, v := range adj[u] {
			dep[v] = dep[u] + 1
			dfs(v)
			sz[u] += sz[v]
		}
	}

	dfs(0)

	isAnc := func(u int, v int) bool {
		return pos[u] <= pos[v] && pos[v] < pos[u]+sz[u]
	}

	findKthAnc := func(u int, k int) int {
		for i := H - 1; i >= 0; i-- {
			if k&(1<<i) > 0 {
				u = fa[u][i]
			}
		}
		return u
	}

	tr := NewBIT(n)
	var ct int
	var maxSonSize int

	play := func(u int) int {
		tr.Update(pos[u], 1)
		// u + 1 - 2 * maxSonSize
		// 更新ct和maxSonSize
		son := -1
		sonSize := 0
		if isAnc(ct, u) {
			son = findKthAnc(u, dep[u]-dep[ct]-1)
			sonSize = tr.QueryRange(pos[son], pos[son]+sz[son]-1)
		} else {
			son = fa[ct][0]
			sonSize = u + 1 - tr.QueryRange(pos[ct], pos[ct]+sz[ct]-1)
		}
		if sonSize > maxSonSize {
			maxSonSize = sonSize
			if maxSonSize > (u+1)/2 {
				ct = son
				maxSonSize = (u + 1) / 2
			}
		}

		return u + 1 - 2*maxSonSize
	}

	ans := make([]int, n-1)
	for i := range n - 1 {
		ans[i] = play(i + 1)
	}
	return ans
}

type BIT []int

func NewBIT(n int) BIT {
	return make(BIT, n+3)
}

func (t BIT) Update(p int, v int) {
	p++
	for p < len(t) {
		t[p] += v
		p += p & -p
	}
}

func (t BIT) Query(p int) int {
	p++
	var res int
	for p > 0 {
		res += t[p]
		p -= p & -p
	}
	return res
}

func (t BIT) QueryRange(l int, r int) int {
	return t.Query(r) - t.Query(l-1)
}
