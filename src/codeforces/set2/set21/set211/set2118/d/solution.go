package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		ans := drive(reader)
		for _, b := range ans {
			if b {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []bool {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	qs := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &qs[i])
	}
	return solve(k, x, d, qs)
}

func solve(k int, x []int, d []int, qs []int) []bool {
	n := len(x)
	to := make([]int, 2*n)
	for i := range to {
		to[i] = -1
	}
	idx := map[int][]int{}
	last := map[int]int{}

	for i, c := range d {
		v := (x[i]%k - c + k) % k
		if j, ok := last[v]; ok {
			to[n+j] = i
		}
		last[v] = i
		idx[v] = append(idx[v], i)
	}

	clear(last)

	for i := n - 1; i >= 0; i-- {
		v := (x[i] + d[i]) % k
		if j, ok := last[v]; ok {
			to[j] = n + i
		}
		last[v] = i
	}

	vis := make([]int8, 2*n)

	var dfs func(u int) bool
	dfs = func(u int) bool {
		if vis[u] != 0 {
			return vis[u] > 0
		}
		vis[u] = -1
		if v := to[u]; v < 0 || dfs(v) {
			vis[u] = 1
			return true
		}
		return false
	}

	ans := make([]bool, len(qs))

	for i, v := range qs {
		id := idx[v%k]
		j := sort.Search(len(id), func(j int) bool {
			return x[id[j]] >= v
		})
		if j == len(id) || dfs(id[j]) {
			ans[i] = true
		}
	}
	return ans
}
