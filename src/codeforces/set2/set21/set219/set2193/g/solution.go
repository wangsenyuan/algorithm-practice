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

	for range tc {
		_, _, ans := drive(reader)
		fmt.Printf("! %d\n", ans)
	}
}

func drive(reader *bufio.Reader) (n int, edges [][]int, res int) {
	ask := func(a int, b int) int {
		fmt.Printf("? %d %d\n", a, b)
		var res int
		fmt.Fscan(reader, &res)
		return res
	}

	fmt.Fscan(reader, &n)
	edges = make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	res = solve(n, edges, ask)
	return
}

func solve(n int, edges [][]int, ask func(a int, b int) int) int {
	// 找到一个叶子节点作为root
	if n == 2 {
		res := ask(1, 1)
		if res == 1 {
			return 1
		}
		return 2
	}

	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	root := -1
	for i := 1; i <= n; i++ {
		if len(adj[i]) == 1 {
			root = i
			break
		}
	}

	var ans int

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		var children []int
		for _, v := range adj[u] {
			if p != v {
				w := dfs(u, v)
				if ans > 0 {
					// 已经找到了
					return 0
				}
				if w > 0 {
					children = append(children, w)
				}
			}
		}
		// 子节点都被处理了
		if len(children) == 0 {
			return u
		}
		// u is checked or not
		var checked bool
		for len(children) >= 2 {
			a := children[0]
			b := children[1]
			children = children[2:]
			res := ask(a, b)
			if res == 0 {
				checked = true
				continue
			}

			// 只能是 a, u, b
			arr := []int{a, b}

			if !checked {
				arr = append(arr, u)
			}

			ans = arr[len(arr)-1]
			for i := 0; i+1 < len(arr); i++ {
				res = ask(arr[i], arr[i])
				if res == 1 {
					ans = arr[i]
					break
				}
			}

			return 0
		}
		if len(children) == 0 {
			return 0
		}
		// len(children) == 1
		v := children[0]
		if checked {
			// u确定不是了，但是v不确定
			return v
		}
		res := ask(u, v)
		if res == 1 {
			res = ask(u, u)
			if res == 1 {
				ans = u
			} else {
				ans = v
			}
		}

		return 0
	}

	dfs(-1, root)

	if ans == 0 {
		// 只可能是root
		ans = root
	}

	return ans
}
