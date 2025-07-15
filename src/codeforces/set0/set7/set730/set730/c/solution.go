package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	roads := make([][]int, m)
	for i := range m {
		roads[i] = readNNums(reader, 2)
	}
	k := readNum(reader)
	stores := make([][]int, k)
	for i := range k {
		stores[i] = readNNums(reader, 3)
	}
	q := readNum(reader)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = readNNums(reader, 3)
	}
	return solve(n, roads, stores, queries)
}

func solve(n int, roads [][]int, stores [][]int, queries [][]int) []int {
	g := make([][]int, n)
	for _, road := range roads {
		u, v := road[0], road[1]
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	que := make([]int, n)

	bfs := func(s int) []int {
		dist := make([]int, n)
		for i := range n {
			dist[i] = -1
		}
		dist[s] = 0
		var head, tail int
		que[head] = s
		head++

		for tail < head {
			u := que[tail]
			tail++
			for _, v := range g[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}

		return dist
	}

	dist := make([][]int, n)

	for i := range n {
		dist[i] = bfs(i)
	}

	slices.SortFunc(stores, func(a, b []int) int {
		return a[2] - b[2]
	})

	check := func(s int, a int, r int, t int) bool {
		var sum int
		var cost int
		for _, cur := range stores {
			c, k, p := cur[0]-1, cur[1], cur[2]
			if dist[s][c] >= 0 && dist[s][c] <= t {
				x := min(k, r-sum)
				sum += x
				cost += x * p
				if sum == r || cost > a {
					break
				}
			}
		}
		return sum == r && cost <= a
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		s, r, a := cur[0], cur[1], cur[2]
		s--
		if !check(s, a, r, n) {
			ans[i] = -1
			continue
		}
		ans[i] = sort.Search(n, func(t int) bool {
			return check(s, a, r, t)
		})
	}

	return ans
}
