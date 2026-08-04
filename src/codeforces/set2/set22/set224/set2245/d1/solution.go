package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		ok, a := drive(reader)
		if !ok {
			fmt.Fprintln(writer, "NO")
			continue
		}
		fmt.Fprintln(writer, "YES")
		for i, v := range a {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) (ok bool, a []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	ops := make([][]int, m)
	for i := range m {
		ops[i] = make([]int, 3)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1], &ops[i][2])
	}
	return solve(n, ops)
}

func solve(n int, ops [][]int) (ok bool, a []int) {
	adj := make([][]int, n)

	b := make([][]int, n)
	for i := range n {
		b[i] = make([]int, n)
	}

	for _, cur := range ops {
		o, i, j := cur[0], cur[1]-1, cur[2]-1
		b[i][j] = o
		b[j][i] = o
	}

	for i := range n {
		for j := range n {
			if b[i][i] == 1 && b[j][j] == 2 {
				if b[i][j] == 1 {
					adj[j] = append(adj[j], i)
				} else {
					adj[i] = append(adj[i], j)
				}
			}
		}
	}

	// topo sort
	deg := make([]int, n)
	for i := range n {
		for _, j := range adj[i] {
			deg[j]++
		}
	}

	ord := make([]int, n)
	var que []int
	for i := range n {
		ord[i] = -1
		if deg[i] == 0 {
			ord[i] = len(que)
			que = append(que, i)
		}
	}
	ans := make([]int, n)
	for it := 0; it < len(que); {
		u := que[it]
		it++
		if b[u][u] == 1 {
			ans[u] = it
		} else {
			ans[u] = -it
		}
		for _, v := range adj[u] {
			deg[v]--
			if deg[v] == 0 {
				ord[v] = len(que)
				que = append(que, v)
			}
		}
	}

	if len(que) != n {
		return false, nil
	}

	for i := range n {
		for j := range n {
			if b[i][j] == 1 && ans[i]+ans[j] < 0 {
				return false, nil
			}
			if b[i][j] == 2 && ans[i]+ans[j] >= 0 {
				return false, nil
			}
		}
	}
	return true, ans
}
