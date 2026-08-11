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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	parent := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &parent[i])
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(parent, a)
}

func solve(parent []int, a []int) string {
	n := len(a)
	if n == 1 {
		return "YES"
	}
	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i-1] - 1
		adj[p] = append(adj[p], i)
	}

	var k int
	for _, v := range a {
		if v > 0 {
			k++
		}
	}

	type pair struct {
		first  int
		second int
	}

	dp := make([][]pair, n)

	var dfs func(u int) bool
	dfs = func(u int) bool {
		if len(adj[u]) == 0 {
			dp[u] = append(dp[u], pair{a[u], a[u]})
			return true
		}

		var arr []pair
		// 必须要能连接起来
		for _, v := range adj[u] {
			if !dfs(v) {
				return false
			}
			arr = append(arr, dp[v]...)
		}

		// 现在要合并, 先找到最小的first所在的位置
		var minAt int
		for i, w := range arr {
			if w.first < arr[minAt].first {
				minAt = i
			}
		}
		var tmp []pair
		tmp = append(tmp, arr[minAt:]...)
		tmp = append(tmp, arr[:minAt]...)

		for i := 0; i < len(tmp); i++ {
			if i == 0 || tmp[i].first > tmp[i-1].second+1 {
				dp[u] = append(dp[u], tmp[i])
			} else if tmp[i].first == tmp[i-1].second+1 {
				dp[u][len(dp[u])-1].second = tmp[i].second
			} else {
				return false
			}
		}
		return len(dp[u]) == 1
	}

	ok := dfs(0)

	if ok {
		return "YES"
	}
	return "NO"
}

func last[T any](arr []T) T {
	return arr[len(arr)-1]
}
