package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("possible")
	} else {
		fmt.Println("impossible")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	forbidden := make([][]int, m)
	for i := 0; i < m; i++ {
		forbidden[i] = make([]int, 2)
		fmt.Fscan(reader, &forbidden[i][0], &forbidden[i][1])
	}
	return solve(n, k, forbidden)
}

type pair struct {
	first  int
	second int
}

func solve(n int, k int, forbiden [][]int) bool {

	adj := make([]map[int]bool, n)
	for i := range n {
		adj[i] = make(map[int]bool)
	}
	for _, e := range forbiden {
		u, v := e[0]-1, e[1]-1
		adj[u][v] = true
		adj[v][u] = true
	}

	var unmarked []int
	for i := 1; i < n; i++ {
		unmarked = append(unmarked, i)
	}

	marked := make([]bool, n)

	firstLayer := make([]bool, n)

	play := func(s int) {
		var que []int
		que = append(que, s)
		marked[s] = true
		for len(que) > 0 {
			u := que[0]
			que = que[1:]

			var next []int
			for _, v := range unmarked {
				if !adj[u][v] {
					// 可以构造一条边
					que = append(que, v)
					if u == 0 {
						firstLayer[v] = true
					}
					marked[v] = true
				} else {
					// v 不能连接到u上去
					next = append(next, v)
				}
			}
			unmarked = next
		}
	}

	play(0)

	if len(unmarked) > 0 {
		return false
	}
	clear(marked)
	var todo []int
	for i := 1; i < n; i++ {
		unmarked = append(unmarked, i)
		if firstLayer[i] {
			// not allow to use 0 again
			adj[i][0] = true
			todo = append(todo, i)
		}
	}

	d := n - 1 - len(adj[0])

	var c int
	for len(unmarked) > 0 {
		u := todo[0]
		todo = todo[1:]
		if marked[u] {
			continue
		}
		c++
		play(u)
	}

	return c <= k && k <= d
}
