package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

const mod = 1_000_000_007

func mul(a, b int) int {
	return a * b % mod
}

const N = 2e5 + 10

var F [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}
}

func solve(n int, edges [][]int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 不能交叉的情况下, 那么假设u得另外一头, 只能是有一个可以接回来
	// 它可以连接很多个, 但是只能是一个, 这边也是, 也就是 ^v^v 这样的结构才行
	// 去掉叶子节点后(只能一层, 剩下的部分, 必须连成一条线)
	leaf := make([]bool, n)
	root := -1
	for i := range n {
		if len(adj[i]) == 1 {
			leaf[i] = true
		} else {
			root = i
		}
	}

	if root < 0 {
		return F[n]
	}

	adj2 := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if leaf[u] || leaf[v] {
			continue
		}
		adj2[u] = append(adj2[u], v)
		adj2[v] = append(adj2[v], u)
	}

	var path []int
	for u := range n {
		if leaf[u] {
			continue
		}
		if len(adj2[u]) > 2 {
			return 0
		}
		path = append(path, u)
		if len(adj2[u]) == 1 {
			root = u
		}
	}

	marked := make([]bool, n)
	marked[root] = true
	markedCnt := 1
	for {
		next := -1
		for _, v := range adj2[root] {
			if !marked[v] {
				next = v
				break
			}
		}
		if next < 0 {
			break
		}
		root = next
		marked[root] = true
		markedCnt++
	}

	if markedCnt != len(path) || len(path) > 1 && len(adj2[root]) != 1 {
		// not connected
		return 0
	}

	res := 4
	if len(path) == 1 {
		res = 2
	}

	for _, u := range path {
		rem := len(adj[u]) - len(adj2[u])
		res = mul(res, F[rem])
	}

	return res
}
