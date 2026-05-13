package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ask := func(k int) []int {
		fmt.Printf("? %d\n", k)
		var q int
		fmt.Fscan(reader, &q)
		if q == 0 {
			return nil
		}
		path := make([]int, q)
		for i := range q {
			fmt.Fscan(reader, &path[i])
		}
		return path
	}
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n, ask)
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("! %d\n", len(res)))
		for _, edge := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", edge[0], edge[1]))
		}
		fmt.Println(buf.String())
	}
}

func solve(n int, ask func(int) []int) [][]int {

	var res [][]int

	k := 2

	dp := make([]int, n+1)
	for i := range n + 1 {
		dp[i] = -1
	}

	adj := make([][]int, n+1)

	var play func(v int)

	play = func(v int) {
		if dp[v] != -1 {
			return
		}

		dp[v] = 1

		for _, u := range adj[v] {
			play(u)
			dp[v] += dp[u]
		}
	}

	prev := []int{1}
	for {
		path := ask(k)
		if len(path) == 0 {
			break
		}

		var common int
		for i := 0; i < min(len(prev), len(path)); i++ {
			if prev[i] != path[i] {
				break
			}
			common++
		}

		if len(prev) > common {
			play(prev[common])
		}

		if common > 0 {
			adj[path[common-1]] = append(adj[path[common-1]], path[common])
			res = append(res, []int{path[common-1], path[common]})
		}
		if dp[path[common]] == -1 {
			k++
		} else {
			k += dp[path[common]]
		}
		prev = path
	}

	return res
}
