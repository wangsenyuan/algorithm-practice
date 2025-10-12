package p3715

import "slices"

func sumOfAncestors(n int, edges [][]int, nums []int) int64 {
	x := slices.Max(nums)

	var primes []int
	lpf := make([]int, x+1)
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > x {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	rep := make([]int, x+1)
	rep[1] = 1

	get := func(x int) int {
		if rep[x] != 0 {
			return rep[x]
		}
		rep[x] = 1
		for i := x; i > 1; {
			var cnt int
			j := lpf[i]
			for i%j == 0 {
				cnt++
				i /= j
			}
			if cnt&1 == 1 {
				rep[x] *= j
			}
		}
		return rep[x]
	}

	freq := make([]int, x+1)

	g := NewGraph(n, 2*n)
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	var dfs func(p int, u int)

	var res int

	dfs = func(p int, u int) {
		y := get(nums[u])
		res += freq[y]
		freq[y]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}
		freq[y]--
	}

	dfs(-1, 0)

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
