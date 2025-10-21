package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func readNum(reader *bufio.Reader) (res int) {
	fmt.Fscan(reader, &res)
	return
}

func drive(reader *bufio.Reader) int {
	n := readNum(reader)
	r := readNNums(reader, n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, r, edges)
}

func solve(n int, r []int, edges [][]int) int {
	if n == 1 {
		return 1
	}
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var best int

	update := func(a []int, mid int, b []int) {
		// a是升序的, b是降序的
		j := len(b)
		for i, v := range a {
			for j > 0 && b[j-1] <= v {
				j--
			}
			best = max(best, i+1+j)
		}
		l := sort.SearchInts(a, mid)
		// l == len(a) or a[l] >= mid
		// l = len(a) or a[l] > mid
		r := sort.Search(len(b), func(i int) bool {
			return b[i] <= mid
		})

		best = max(best, l+1+r)
	}

	merge1 := func(a []int, b []int) []int {
		// a和b都是升序的, 如果可以在b中同样的位置，得到一个更小的值，应该使用它
		if len(a) > len(b) {
			a, b = b, a
		}
		for i := 0; i < len(a); i++ {
			b[i] = min(a[i], b[i])
		}
		return b
	}

	merge2 := func(a []int, b []int) []int {
		// a和b都是降序的
		if len(a) > len(b) {
			a, b = b, a
		}
		for i := 0; i < len(a); i++ {
			b[i] = max(a[i], b[i])
		}
		return b
	}

	var dfs func(p int, u int) (arr1 []int, arr2 []int)

	dfs = func(p int, u int) (arr1 []int, arr2 []int) {
		for _, v := range adj[u] {
			if p != v {
				x, y := dfs(u, v)
				update(arr1, r[u], y)
				update(x, r[u], arr2)
				arr1 = merge1(arr1, x)
				arr2 = merge2(arr2, y)
			}
		}
		i := sort.SearchInts(arr1, r[u])
		if i == len(arr1) {
			arr1 = append(arr1, r[u])
		} else {
			arr1[i] = r[u]
		}
		i = sort.Search(len(arr2), func(i int) bool {
			return r[u] >= arr2[i]
		})
		if i == len(arr2) {
			arr2 = append(arr2, r[u])
		} else {
			arr2[i] = r[u]
		}
		best = max(best, len(arr1), len(arr2))
		return
	}

	dfs(-1, 0)

	return best
}
