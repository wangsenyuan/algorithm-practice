package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", res[0])
	fmt.Println(s[1 : len(s)-1])
	s = fmt.Sprintf("%v", res[1])
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (n int, edges [][]int, res [][]int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {
	m := len(edges)
	if m == n*(n-1)/2 {
		// 完全图， 没有答案
		return nil
	}
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		slices.Sort(adj[i])
		a[i] = i + 1
		b[i] = i + 1
	}

	for i := range n {
		var c int
		for _, j := range adj[i] {
			if c == i {
				c++
			}
			if c != j {
				break
			}
			c++
		}
	
		if c == i {
			c++
		}
		if c < n {
			b[c] = a[i]
			for j := i + 1; j < c; j++ {
				a[j]++
				b[j]++
			}
			a[c] = a[i] + 1
			return [][]int{a, b}
		}
	}

	return nil
}
