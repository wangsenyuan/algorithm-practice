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
	found, chest, stomach, heads, tails, _, _, _, _ := process(reader)
	if !found {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	buf.WriteString(fmt.Sprintf("%d %d\n", chest, stomach))
	for _, u := range heads {
		buf.WriteString(fmt.Sprintf("%d ", u))
	}
	buf.WriteByte('\n')
	for _, u := range tails {
		buf.WriteString(fmt.Sprintf("%d ", u))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (found bool, chest int, stomach int, heads []int, tails []int, n int, h int, t int, edges [][]int) {
	line := readNNums(reader, 4)
	n, m, h, t := line[0], line[1], line[2], line[3]
	edges = make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}
	found, chest, stomach, heads, tails = solve(n, m, h, t, edges)
	return
}

type pair struct {
	first  int
	second int
}

func solve(n int, m int, h int, t int, edges [][]int) (found bool, chest int, stomach int, heads []int, tails []int) {
	g := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	for i := range n {
		sort.Ints(g[i])
	}

	// x这些可以算到一边去
	// a + b = x
	// u - a >= h
	// v - b >= t
	// len(g[u]) >= h + a
	// len(g[v]) >= t + b

	check := func(u int, v int) (ok bool, hs []int, ts []int, x []int) {
		if len(g[u]) <= h || len(g[v]) <= t {
			return false, nil, nil, nil
		}
		var i, j int
		for i < len(g[u]) || j < len(g[v]) {
			if i < len(g[u]) && g[u][i] == v {
				i++
				continue
			}
			if j < len(g[v]) && g[v][j] == u {
				j++
				continue
			}
			if i < len(g[u]) && j < len(g[v]) && g[u][i] == g[v][j] {
				x = append(x, g[u][i])
				i++
				j++
			} else if j == len(g[v]) || i < len(g[u]) && g[u][i] < g[v][j] {
				hs = append(hs, g[u][i])
				i++
			} else {
				ts = append(ts, g[v][j])
				j++
			}
			// hs 是只和u相连的部分， ts只是和v相连的部分， x是两边都相连的部分
			du := max(0, h-len(hs))
			dv := max(0, t-len(ts))
			if du+dv <= len(x) {
				return true, hs, ts, x
			}
		}
		return false, nil, nil, nil
	}

	prepare := func(u int, v int, hs []int, ts []int, x []int) (bool, int, int, []int, []int) {
		for len(hs) < h {
			hs = append(hs, x[0])
			x = x[1:]
		}
		for len(ts) < t {
			ts = append(ts, x[0])
			x = x[1:]
		}
		for i := range h {
			hs[i]++
		}
		for i := range t {
			ts[i]++
		}
		return true, u + 1, v + 1, hs[:h], ts[:t]
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		ok, hs, ts, x := check(u, v)
		if ok {
			return prepare(u, v, hs, ts, x)
		}
		ok, hs, ts, x = check(v, u)
		if ok {
			return prepare(v, u, hs, ts, x)
		}
	}
	return false, 0, 0, nil, nil
}
