package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		res := drive(reader)
		if res {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) bool {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	roads := make([][]int, m)
	for i := range m {
		roads[i] = make([]int, 2)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1])
	}
	var w int
	fmt.Fscan(reader, &w)
	calenders := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &calenders[i])
	}
	return solve(n, roads, calenders)
}

func solve(n int, roads [][]int, calenders []string) bool {
	w := len(calenders[0])

	adj := make([][]int, n*w)

	for id, cal := range calenders {
		for i := range w {
			j := (i + 1) % w
			if cal[i] == 'o' && cal[j] == 'o' {
				adj[id*w+i] = append(adj[id*w+i], id*w+j)
			}
		}
	}

	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		for i := range w {
			j := (i + 1) % w
			if calenders[u][i] == 'o' && calenders[v][j] == 'o' {
				adj[u*w+i] = append(adj[u*w+i], v*w+j)
			}
			if calenders[v][i] == 'o' && calenders[u][j] == 'o' {
				adj[v*w+i] = append(adj[v*w+i], u*w+j)
			}
		}
	}

	mark := make([]int, n*w)

	var dfs func(u int, i int) bool
	dfs = func(u int, i int) bool {
		if mark[u*w+i] == 1 {
			return true
		}
		mark[u*w+i]++
		for _, cur := range adj[u*w+i] {
			v := cur / w
			j := cur % w
			if mark[cur] < 2 && dfs(v, j) {
				return true
			}
		}

		mark[u*w+i]++
		return false
	}

	for u := range n {
		for i := range w {
			if mark[u*w+i] == 0 && dfs(u, i) {
				return true
			}
		}
	}

	return false
}
