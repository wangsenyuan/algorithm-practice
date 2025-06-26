package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		ans := process(reader)
		s := fmt.Sprintf("%v", ans)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

type pair struct {
	first  int
	second int
}

type data struct {
	d int
	u int
	v int
}

func solve(n int, edges [][]int) []int {

	g := make([][]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	que := make([]int, n)
	dist := make([]int, n)
	track := make([]int, n)
	for i := range n {
		dist[i] = -1
		track[i] = -1
	}
	marked := make([]bool, n)

	reset := func(p int) {
		for i := range p {
			dist[que[i]] = -1
		}
	}

	bfs := func(s int) int {
		dist[s] = 0
		var head, tail int
		que[head] = s
		head++
		for tail < head {
			u := que[tail]
			tail++
			for _, v := range g[u] {
				if dist[v] < 0 && !marked[v] {
					dist[v] = dist[u] + 1
					track[v] = u
					que[head] = v
					head++
				}
			}
		}
		return head
	}

	var res []data

	var loop func(s int)

	loop = func(s int) {
		p := bfs(s)

		if p == 1 {
			res = append(res, data{1, s + 1, s + 1})
			return
		}

		for i := range p {
			x := que[i]
			if dist[x] > dist[s] || dist[x] == dist[s] && x > s {
				s = x
			}
		}

		reset(p)
		p = bfs(s)

		t := s
		for i := range p {
			x := que[i]
			if dist[x] > dist[t] || dist[x] == dist[t] && x > t {
				t = x
			}
		}

		if p <= 3 {
			res = append(res, data{p, max(s, t) + 1, min(s, t) + 1})
			return
		}

		res = append(res, data{dist[t] + 1, max(s, t) + 1, min(s, t) + 1})

		reset(p)
		track[s] = -1

		for u := t; u != -1; u = track[u] {
			marked[u] = true
			for _, v := range g[u] {
				if marked[v] || v == track[u] {
					continue
				}
				loop(v)
			}
		}
	}

	loop(0)

	slices.SortFunc(res, func(a, b data) int {
		return cmp.Or(b.d-a.d, b.u-a.u)
	})

	var ans []int
	for _, cur := range res {
		ans = append(ans, cur.d, cur.u, cur.v)
	}
	return ans
}
