package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		if x {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}
func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) []bool {
	n, m := readTwoNums(reader)
	p := readNNums(reader, n-1)
	labels := readString(reader)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(n, p, labels, queries)
}

func solve(n int, p []int, labels string, queries [][]int) []bool {
	adj := make([][]int, n)

	for i := 1; i < n; i++ {
		adj[p[i-1]-1] = append(adj[p[i-1]-1], i)
	}

	buf_at_depth := make([][]byte, n)
	var id int32
	where := make([]int32, n)
	pos := make([]int32, n)
	sz := make([]int32, n)
	h := make([]int32, n)

	var dfs func(u int)
	dfs = func(u int) {
		where[u] = int32(len(buf_at_depth[h[u]]))
		buf_at_depth[h[u]] = append(buf_at_depth[h[u]], labels[u])
		pos[u] = id
		id++
		sz[u] = 1
		for _, v := range adj[u] {
			h[v] = h[u] + 1
			dfs(v)
			sz[u] += sz[v]
		}
	}

	dfs(0)

	at := make([][]int, n)
	for i, cur := range queries {
		v := cur[0] - 1
		at[v] = append(at[v], i)
	}

	type data struct {
		l int32
		r int32
	}

	at_depth := make([]int, n)

	isAnc := func(u int, v int) bool {
		return pos[u] < pos[v] && pos[v] < pos[u]+sz[u]
	}

	ans := make([]bool, len(queries))

	var dfs1 func(u int)

	records := make([]data, len(queries))

	dfs1 = func(u int) {
		at_depth[h[u]] = u
		for _, v := range adj[u] {
			dfs1(v)
		}
		for _, i := range at[u] {
			h := queries[i][1] - 1
			if u == at_depth[h] || !isAnc(u, at_depth[h]) {
				ans[i] = true
			} else {
				records[i].r = where[at_depth[h]]
			}
		}
	}

	dfs1(0)

	var dfs2 func(u int)
	dfs2 = func(u int) {
		at_depth[h[u]] = u

		for i := len(adj[u]) - 1; i >= 0; i-- {
			v := adj[u][i]
			dfs2(v)
		}

		for _, i := range at[u] {
			h := queries[i][1] - 1
			if u != at_depth[h] && isAnc(u, at_depth[h]) {
				records[i].l = where[at_depth[h]]
			}
		}
	}

	dfs2(0)

	pref := make([][]int32, n)
	for i := range n {
		pref[i] = make([]int32, len(buf_at_depth[i])+1)
		for j := 0; j < len(buf_at_depth[i]); j++ {
			x := int(buf_at_depth[i][j] - 'a')
			pref[i][j+1] = pref[i][j] ^ (1 << x)
		}
	}

	for i := range len(queries) {
		h := queries[i][1] - 1
		if !ans[i] {
			l, r := records[i].l, records[i].r
			mask := pref[h][r+1] ^ pref[h][l]
			ans[i] = bits.OnesCount(uint(mask)) <= 1
		}
	}
	return ans
}
