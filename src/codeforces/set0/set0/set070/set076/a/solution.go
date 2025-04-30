package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	g, s := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 4)
	}
	return solve(n, edges, g, s)
}

type edge struct {
	u int
	v int
	g int
	s int
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 62

func solve(n int, edges [][]int, G int, S int) int {
	var arr []edge
	var mc int
	for _, cur := range edges {
		u, v, g, s := cur[0], cur[1], cur[2], cur[3]
		u--
		v--
		if u != v {
			arr = append(arr, edge{u, v, g, s})
			mc = max(mc, g*G+s*S)
		}
	}

	slices.SortFunc(arr, func(a, b edge) int {
		return cmp.Or(a.g-b.g, a.s-b.s)
	})

	ans := inf

	g := make([][]pair, n)

	var silver int

	que := make([]int, n)
	marked := make([]bool, n)
	fa := make([]pair, n)
	bfs := func(s int) int {
		clear(marked)
		var head, tail int
		que[head] = s
		head++
		marked[s] = true
		for tail < head {
			u := que[tail]
			tail++
			for _, it := range g[u] {
				v := it.first
				w := it.second
				silver = max(silver, w)
				if !marked[v] {
					fa[v] = pair{u, w}
					marked[v] = true
					que[head] = v
					head++
				}
			}
		}
		return head
	}

	remove := func(u int, v int) {
		for i := range g[u] {
			if g[u][i].first == v {
				copy(g[u][i:], g[u][i+1:])
				g[u] = g[u][:len(g[u])-1]
				return
			}
		}
	}

	disconnect := func(u int, v int) {
		remove(u, v)
		remove(v, u)
	}

	for _, cur := range arr {
		u, v, gold, s := cur.u, cur.v, cur.g, cur.s
		if u == v {
			continue
		}
		bfs(u)
		if !marked[v] {
			g[u] = append(g[u], pair{v, s})
			g[v] = append(g[v], pair{u, s})
		} else {
			// v is reachable from u
			v1 := v
			s1 := s
			v2 := v
			for v1 != u {
				// 把最大的替换了
				if fa[v1].second > s1 {
					s1 = fa[v1].second
					v2 = v1
				}
				v1 = fa[v1].first
			}
			if s1 != s {
				u2 := fa[v2].first
				disconnect(u2, v2)
				g[u] = append(g[u], pair{v, s})
				g[v] = append(g[v], pair{u, s})
			}
		}

		silver = 0

		sz := bfs(u)

		if sz == n {
			ans = min(ans, gold*G+silver*S)
		}
	}

	if ans == inf {
		return -1
	}

	return ans
}
