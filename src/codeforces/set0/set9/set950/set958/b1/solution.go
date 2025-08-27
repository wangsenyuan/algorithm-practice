package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		var u, v int
		fmt.Scan(&u, &v)
		edges[i] = []int{u, v}
	}
	res := solve(n, edges)
	fmt.Println(res)
}

func solve(n int, edges [][]int) int {
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
	}

	ans := 0
	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			ans++
		}
	}
	return ans
}
