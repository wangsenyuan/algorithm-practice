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
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		edges[i] = []int{u, v, w}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	g := make([][]Pair, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], Pair{v, w})
		g[v] = append(g[v], Pair{u, w})
	}

	parent := make([]int, n)
	parentWeight := make([]int, n)
	for i := range n {
		parent[i] = -1
	}

	order := make([]int, 0, n)
	stack := []int{0}
	parent[0] = 0
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, u)
		for _, e := range g[u] {
			if e.first == parent[u] {
				continue
			}
			parent[e.first] = u
			parentWeight[e.first] = e.second
			stack = append(stack, e.first)
		}
	}

	a := make([]int, n)
	b := make([]int, n)
	for i := n - 1; i > 0; i-- {
		u := order[i]
		if len(g[u]) == 1 {
			continue
		}
		sumA := 0
		sumB := parentWeight[u] % mod
		for _, e := range g[u] {
			v := e.first
			if v == parent[u] {
				continue
			}
			sumA = add(sumA, a[v])
			sumB = add(sumB, add(e.second%mod, b[v]))
		}

		den := sub(len(g[u])%mod, sumA)
		inv := pow(den, mod-2)
		a[u] = inv
		b[u] = mul(sumB, inv)
	}

	sumA := 0
	sumB := 0
	for _, e := range g[0] {
		v := e.first
		sumA = add(sumA, a[v])
		sumB = add(sumB, add(e.second%mod, b[v]))
	}

	den := sub(len(g[0])%mod, sumA)
	return mul(sumB, pow(den, mod-2))
}

const mod = 1000000007

type Pair struct {
	first  int
	second int
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	a -= b
	if a < 0 {
		a += mod
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}
