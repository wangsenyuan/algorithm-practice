package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(len(res))
	if len(res) > 0 {
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges, a)
}

func solve(n int, edges [][]int, a []int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	freq := make([]int, n)
	for _, v := range a {
		freq[v]++
	}

	if freq[0] > 1 {
		// 不能有超过1个0
		return nil
	}

	valid := make(map[Hash]bool)

	bases := make([]Hash, n)
	bases[0] = NewHash(1)
	for i := 1; i < n; i++ {
		bases[i] = bases[i-1].MulInt(n)
	}

	var target Hash
	for i := range n {
		if freq[i] > 0 {
			target = target.Add(bases[i].MulInt(freq[i]))
		}
	}
	// 剩下的那个数随便设置（0...n-1)
	for i := n - 1; i >= 0; i-- {
		valid[target.Add(bases[i])] = true
	}

	dp := make([]Hash, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)
				dp[u] = dp[u].Add(dp[v].MulInt(n))
			}
		}
		dp[u] = dp[u].AddInt(1)
	}

	dfs(-1, 0)

	var res []int

	var dfs2 func(p int, u int, w Hash)
	dfs2 = func(p int, u int, w Hash) {
		cur := dp[u]
		if p >= 0 {
			cur = cur.Add(w.MulInt(n))
		}
		if valid[cur] {
			res = append(res, u+1)
		}
		for _, v := range adj[u] {
			if p != v {
				tmp := cur.Sub(dp[v].MulInt(n))
				dfs2(u, v, tmp)
			}
		}
	}

	dfs2(-1, 0, NewHash(0))

	slices.Sort(res)

	return res
}

var MOD = [...]uint{1000000007, 1000000009}

type Hash struct {
	h [2]uint
}

func NewHash(x uint) Hash {
	h := [2]uint{uint(x) % MOD[0], uint(x) % MOD[1]}
	return Hash{h}
}

func (this Hash) Sub(that Hash) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + MOD[i] - that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Add(that Hash) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) AddInt(x int) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + uint(x)%MOD[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Mul(that Hash) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) MulInt(x int) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * uint(x)) % MOD[i]
	}
	return Hash{h}
}
