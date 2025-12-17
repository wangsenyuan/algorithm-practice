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
	var a, b, k int64
	fmt.Fscan(reader, &a, &b, &k)
	return int(solve(a, b, k))
}

func solve(a int64, b int64, k int64) int64 {
	if a <= b {
		return 0
	}

	// Calculate LCM of [2, k]
	mc := int64(1)
	for i := int64(2); i <= k; i++ {
		mc = (mc * i) / gcd(mc, i)
	}

	// Precompute DP for [0, mc-1]
	dp := make([]int64, mc)
	for i := int64(1); i < mc; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := int64(1); i < mc; i++ {
		dfs(i, k, dp)
	}

	ans := int64(0)

	// If we can reduce a by a%mc and still be > b, do it
	if a-a%mc > b {
		ans += dp[a%mc]
		a -= a % mc
	}

	// Calculate how many full periods we can skip
	df := a - b
	rep := df / mc
	a = a - rep*mc
	ans += rep * (1 + dp[mc-1])

	// Handle the remaining part
	df = a - b
	if df > 0 {
		if a%mc == 0 {
			ans += bfs((a-1)%mc, b%mc, k) + 1
		} else {
			ans += bfs(a%mc, b%mc, k)
		}
	}

	return ans
}

func gcd(x, y int64) int64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func dfs(x int64, k int64, dp []int64) int64 {
	if dp[x] != -1 {
		return dp[x]
	}

	tp := int64(1 << 30) // INF equivalent
	for i := int64(2); i <= k; i++ {
		if x%i != 0 {
			next := dfs(x-x%i, k, dp)
			if next < tp {
				tp = next
			}
		}
	}
	next := dfs(x-1, k, dp)
	if next < tp {
		tp = next
	}
	dp[x] = tp + 1
	return dp[x]
}

func bfs(st, ed int64, k int64) int64 {
	if st == ed {
		return 0
	}

	visited := make(map[int64]bool)
	queue := []int64{st}
	steps := []int64{0}
	visited[st] = true
	head := 0

	for head < len(queue) {
		current := queue[head]
		currentSteps := steps[head]
		head++

		if current == ed {
			return currentSteps
		}

		// Type 2: subtract (current mod x) for each x in [2, k]
		for i := int64(2); i <= k; i++ {
			next := current - current%i
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
				steps = append(steps, currentSteps+1)
			}
		}

		// Type 1: subtract 1
		next := current - 1
		if !visited[next] {
			visited[next] = true
			queue = append(queue, next)
			steps = append(steps, currentSteps+1)
		}
	}

	return -1
}
