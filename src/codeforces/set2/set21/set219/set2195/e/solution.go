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

	for _, row := range drive(reader) {
		s := fmt.Sprintf("%v", row)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

const mod = 1e9 + 7

func drive(reader *bufio.Reader) [][]int {
	var tc int
	fmt.Fscan(reader, &tc)
	res := make([][]int, tc)
	for i := range tc {
		var n int
		fmt.Fscan(reader, &n)
		children := make([][]int, n)
		for j := range n {
			children[j] = make([]int, 2)
			fmt.Fscan(reader, &children[j][0], &children[j][1])
		}
		res[i] = solve(children)
	}
	return res
}

func solve(children [][]int) []int {
	// TODO
	n := len(children)
	fa := make([]int, n+1)
	for i := range n {
		if children[i][0] != 0 {
			fa[children[i][0]] = i + 1
		}
		if children[i][1] != 0 {
			fa[children[i][1]] = i + 1
		}
	}

	fp := make([]int, n+1)

	var dfs func(v int)
	dfs = func(v int) {
		if children[v-1][0] == 0 {
			fp[v] = 0
			return
		}
		l, r := children[v-1][0], children[v-1][1]
		dfs(l)
		dfs(r)
		fp[v] = 4 + fp[l] + fp[r]
	}

	dfs(1)

	ans := make([]int, n+1)

	var dfs2 func(v int)

	dfs2 = func(v int) {
		if v == 1 {
			// 还需要返回0
			ans[v] = fp[v] + 1
		} else {
			ans[v] = fp[v] + 1 + ans[fa[v]]
		}

		ans[v] %= mod

		l, r := children[v-1][0], children[v-1][1]
		if l != 0 {
			dfs2(l)
			dfs2(r)
		}
	}

	dfs2(1)

	return ans[1:]
}
