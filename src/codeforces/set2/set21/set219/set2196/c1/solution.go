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
	lo := make([]int, n+1)
	hi := make([]int, n+1)

	l := 1
	for i := 1; i <= n; i++ {
		lo[i] = l
		r := 1 << 30
		for l < r {
			mid := (l + r) >> 1
			path := ask(mid)
			if len(path) == 0 || path[0] != i {
				r = mid
			} else {
				l = mid + 1
			}
		}
		hi[i] = r
	}

	var res [][]int
	marked := make([]bool, n+1)

	var dfs func(u int)

	dfs = func(u int) {
		marked[u] = true
		// lo[u] 肯定是[u]
		k := lo[u] + 1
		for k < hi[u] {
			path := ask(k)
			// len(path) > 1
			v := path[1]
			res = append(res, []int{u, v})
			if !marked[v] {
				dfs(v)
			}
			k += hi[v] - lo[v]
		}
	}

	for u := 1; u <= n; u++ {
		if !marked[u] {
			dfs(u)
		}
	}

	return res
}
