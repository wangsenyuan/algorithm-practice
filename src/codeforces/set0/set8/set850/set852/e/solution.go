package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	return solve(n, edges)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
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

func solve(n int, edges [][]int) int {
	if n == 1 {
		return 1
	}
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
	}

	var leaf_cnt int
	for u := range n {
		if deg[u] == 1 {
			leaf_cnt++
		}
	}

	internal_cnt := n - leaf_cnt

	// 每个内部节点，可以有 pow(2, internal_cnt)中排列
	res1 := mul(pow(2, internal_cnt), internal_cnt)
	// 每个叶子节点，可以有 pow(2, internal_cnt + 1)中排列
	res2 := mul(pow(2, internal_cnt+1), leaf_cnt)

	return add(res1, res2)
}
