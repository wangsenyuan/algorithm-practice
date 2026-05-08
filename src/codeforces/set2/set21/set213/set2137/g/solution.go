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
		res := drive(reader)
		for _, cur := range res {
			if cur {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}
}

func drive(reader *bufio.Reader) []bool {
	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(n, edges, queries)
}

type data struct {
	id    int
	state int
}

func solve(n int, edges [][]int, queries [][]int) []bool {
	adj := make([][]int, n)
	deg := make([]int, n)
	// cnt[u] 表示有多少个后继节点，dp2[v] = false（如果到了这些节点，因为是River操作）
	cnt := make([]int, n)
	dp1 := make([]bool, n)
	dp2 := make([]bool, n)
	// dp1[i] 表示在节点i，且Cry操作时，cry是否可以获胜
	// dp2[i] 表示在节点i，且目前由River操作时，Cry是否可以获胜
	for i := range n {
		dp1[i] = true
		dp2[i] = true
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[v] = append(adj[v], u)
		deg[u]++
	}

	que := make([]data, n*2)

	flip := func(s int) {
		var head, tail int
		if dp1[s] {
			dp1[s] = false
			que[head] = data{s, 1}
			head++
		}
		if dp2[s] {
			dp2[s] = false
			que[head] = data{s, 2}
			head++
		}

		for tail < head {
			u, w := que[tail].id, que[tail].state
			tail++

			if w == 1 {
				// 如果在节点u处，即使Cry操作，也是失败的，那么它所有的前继节点, dp2[v] = false
				for _, v := range adj[u] {
					if dp2[v] == true {
						dp2[v] = false
						que[head] = data{v, 2}
						head++
					}
				}
			}
			if w == 2 {
				for _, v := range adj[u] {
					cnt[v]++
					// 如果它的所有后继节点都是失败节点
					if cnt[v] == deg[v] && dp1[v] == true {
						dp1[v] = false
						que[head] = data{v, 1}
						head++
					}
				}
			}
		}
	}

	var ans []bool

	for _, cur := range queries {
		if cur[0] == 1 {
			flip(cur[1] - 1)
		} else {
			u := cur[1] - 1
			ans = append(ans, dp1[u])
		}
	}

	return ans
}
