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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(n, k, p)
}

type pair struct {
	first  int
	second int
}

func solve(n int, k int, p []int) int {
	// 在经过k轮后，dp[i]表示i，i存活，切i是最小存活的下标
	mx := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		mx[i] = max(mx[i+1], p[i])
	}
	if n == 1 || mx[0] == 0 {
		return 1
	}
	dist := make([][]int, n+3)
	for i := range n + 3 {
		dist[i] = make([]int, n+3)
	}

	var que []pair

	var ans int

	add := func(a, b int, v int) {
		if dist[a][b] == 0 {
			dist[a][b] = v
			que = append(que, pair{a, b})
			ans++
		}
	}

	add(0, 1, 1)

	for len(que) > 0 {
		cur := que[0]

		que = que[1:]
		a, b := cur.first, cur.second

		if a >= n || b >= n || dist[a][b] > k {
			continue
		}
		// a和b都被消灭了
		if p[a] > 0 && mx[b] > 0 {
			add(b+1, b+2, dist[a][b]+1)
		}

		// a幸存, b消灭
		if p[a] > 0 && mx[b] < 100 {
			add(a, b+1, dist[a][b]+1)
		}

		// a消灭, b幸存
		if p[a] < 100 && mx[b] > 0 {
			add(b, b+1, dist[a][b]+1)
		}
	}

	return ans
}
