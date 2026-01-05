package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.10f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	dancers := make([][]int, n)
	for i := range n {
		dancers[i] = make([]int, 3)
		fmt.Fscan(reader, &dancers[i][0], &dancers[i][1], &dancers[i][2])
	}
	return solve(dancers)
}

func solve(dancers [][]int) float64 {
	n := len(dancers)

	slices.SortFunc(dancers, func(a, b []int) int {
		return a[2] - b[2]
	})

	adj := make([][]int, n)

	cover := func(i int, j int) bool {
		if dancers[i][2] <= dancers[j][2] {
			return false
		}
		dx := dancers[i][0] - dancers[j][0]
		dy := dancers[i][1] - dancers[j][1]
		return dx*dx+dy*dy <= dancers[i][2]*dancers[i][2]
	}

	fa := make([]int, n)

	for i := range n {
		fa[i] = -1
		for j := i + 1; j < n; j++ {
			if cover(j, i) {
				fa[i] = j
				adj[j] = append(adj[j], i)
				break
			}
		}
	}

	return math.Pi * float64(play(dancers, adj, fa))
}

func play(dancers [][]int, adj [][]int, fa []int) int {
	n := len(dancers)
	dp := make([][2][2]int, n)

	sign := func(v int) int {
		if v == 0 {
			return 1
		}
		return -1
	}

	var dfs func(u int)
	dfs = func(u int) {
		var w [2][2]int
		for _, v := range adj[u] {
			dfs(v)
			for i := range 2 {
				for j := range 2 {
					w[i][j] += dp[v][i][j]
				}
			}
		}

		r2 := dancers[u][2] * dancers[u][2]
		for i := range 2 {
			for j := range 2 {
				dp[u][i][j] = max(w[i^1][j]+r2*sign(i), w[i][j^1]+r2*sign(j))
			}
		}
	}
	var ans int
	for u := range n {
		if fa[u] < 0 {
			dfs(u)
			ans += dp[u][0][0]
		}
	}
	return ans
}
