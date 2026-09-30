package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()

	reader := bufio.NewReader(r)
	fmt.Fprintln(w, drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	spiders := make([][][2]int, n)
	for i := range n {
		var beads int
		fmt.Fscan(reader, &beads)
		edges := make([][2]int, beads-1)
		for j := range edges {
			fmt.Fscan(reader, &edges[j][0], &edges[j][1])
		}
		spiders[i] = edges
	}
	return solve(spiders)
}

func solve(spiders [][][2]int) int {
	var res int
	for _, cur := range spiders {
		res += getDimeter(cur)
	}

	return res
}

func getDimeter(g [][2]int) int {
	n := len(g) + 1
	adj := make([][]int, n)
	for _, cur := range g {
		u, v := cur[0]-1, cur[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var res int

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		var far []int
		for _, v := range adj[u] {
			if p != v {
				c := dfs(u, v)
				for i := range far {
					if c >= far[i] {
						far[i], c = c, far[i]
					}
				}
				if len(far) < 2 {
					far = append(far, c)
				}
			}
		}

		if len(far) == 2 {
			res = max(res, far[0]+1+far[1])
		} else if len(far) == 1 {
			res = max(res, far[0]+1)
		} else {
			res = max(res, 1)
			far = append(far, 0)
		}
		return far[0] + 1
	}

	dfs(-1, 0)

	return res - 1
}
