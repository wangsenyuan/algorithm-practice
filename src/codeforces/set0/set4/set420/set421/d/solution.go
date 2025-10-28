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
	var n, p int
	fmt.Fscan(reader, &n, &p)
	pairs := make([][]int, n)
	for i := range n {
		pairs[i] = make([]int, 2)
		fmt.Fscan(reader, &pairs[i][0], &pairs[i][1])
	}
	return solve(p, pairs)
}

func solve(p int, pairs [][]int) int {
	n := len(pairs)
	adj := make([][]int, n)
	deg := make([]int, n)
	for _, cur := range pairs {
		x, y := cur[0]-1, cur[1]-1
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
		deg[x]++
		deg[y]++
	}

	// 不排序似乎也可以的
	var res int
	freq := make(BIT, n+3)

	for u := range n {
		for _, v := range adj[u] {
			if v < u {
				freq.update(deg[v], -1)
				deg[v]--
				freq.update(deg[v], 1)
			}
		}
		// d + dv >= p
		res += freq.rangeQuery(max(0, p-deg[u]), n)

		for _, v := range adj[u] {
			if v < u {
				freq.update(deg[v], -1)
				deg[v]++
				freq.update(deg[v], 1)
			}
		}
		freq.update(deg[u], 1)
	}

	return res
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	var res int
	i++
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) rangeQuery(l int, r int) int {
	return bit.get(r) - bit.get(l-1)
}
