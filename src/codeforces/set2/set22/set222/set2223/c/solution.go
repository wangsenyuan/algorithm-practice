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
	var n, q int
	fmt.Fscan(reader, &n, &q)
	f := make([]int, n-1)
	for u := range n - 1 {
		fmt.Fscan(reader, &f[u])
	}
	l := make([]int, n-1)
	for u := range n - 1 {
		fmt.Fscan(reader, &l[u])
	}
	ms := make([]int, q)
	for i := range q {
		fmt.Fscan(reader, &ms[i])
	}
	return formatInts(solve(f, l, ms))
}

func formatInts(a []int) string {
	s := fmt.Sprintf("%v", a)
	return s[1 : len(s)-1]
}

const inf = 1e18

func solve(f, l []int, ms []int) []int {
	n := len(f) + 1
	g := make([][]int, n)
	dist := make([]int, n)

	for i := 1; i < n; i++ {
		g[f[i-1]-1] = append(g[f[i-1]-1], i)
		dist[i] = dist[f[i-1]-1] + l[i-1]
	}

	ans := make([]int, len(ms))

	var dfs func(u int, todo []int, rem int, M int)

	dfs = func(u int, todo []int, rem int, M int) {
		if len(g[u]) == 0 {
			for _, i := range todo {
				ans[i] = u + 1
			}
			return
		}
		if M > inf || M%len(g[u]) == 0 {
			j := (rem + dist[u]) % len(g[u])
			v := g[u][j]
			dfs(v, todo, rem, M)
			return
		}

		next := make([][]int, len(g[u]))
		for _, i := range todo {
			j := (ms[i] + dist[u]) % len(g[u])
			next[j] = append(next[j], i)
		}

		M = lcm(M, len(g[u]))

		for j, newTodo := range next {
			if newTodo != nil {
				dfs(g[u][j], newTodo, ms[newTodo[0]]%M, M)
			}
		}
	}

	todo := make([]int, len(ms))
	for i := range todo {
		todo[i] = i
	}

	dfs(0, todo, 0, 1)

	return ans
}

func gcd(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	c := gcd(a, b)
	a /= c
	if a > inf/b {
		return inf + 1
	}
	return a * b
}
