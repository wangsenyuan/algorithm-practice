package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	for _, x := range drive(reader) {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	ops := make([][]int, m)
	for i := range m {
		var typ int
		fmt.Fscan(reader, &typ)
		if typ == 0 {
			ops[i] = make([]int, 4)
			fmt.Fscan(reader, &ops[i][1], &ops[i][2], &ops[i][3])
		} else {
			ops[i] = make([]int, 2)
			ops[i][0] = typ
			fmt.Fscan(reader, &ops[i][1])
		}
	}
	return solve(n, edges, ops)
}

func solve(n int, edges [][]int, ops [][]int) []int {
	g := make([][]int, n)
	for _, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	fa := make([][]int, n)
	dep := make([]int, n)
	H := bits.Len(uint(n))
	tin := make([]int, n)
	tout := make([]int, n)
	var timer int
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, H)
		fa[u][0] = p
		for i := 1; i < H; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		tin[u] = timer
		timer++
		for _, v := range g[u] {
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
		tout[u] = timer
	}

	dfs(0, 0)

	upToDep := func(v int, d int) int {
		for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
			v = fa[v][bits.TrailingZeros32(k)]
		}
		return v
	}

	getLca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		u = upToDep(u, dep[v])
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	type query struct {
		a int
		b int
		v int
	}
	qs := make([]query, len(ops))
	idx := make([]int, len(ops))
	var sorted []int
	for i, op := range ops {
		idx[i] = i
		if op[0] == 0 {
			qs[i] = query{op[1] - 1, op[2] - 1, op[3]}
			sorted = append(sorted, op[3])
		} else if op[0] == 1 {
			k := op[1] - 1
			qs[i] = qs[k]
			qs[i].v *= -1
		} else {
			// op[2]
			qs[i] = query{-(op[1] - 1), -1, -1}
		}
	}
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	tr := newFenwick(n + 1)

	var play func(lo int, hi int, todo []int)
	play = func(lo int, hi int, todo []int) {
		for _, i := range todo {
			if qs[i].b < 0 {
				goto next
			}
		}
		return
	next:
		if lo+1 == hi {
			if lo >= 0 {
				for _, i := range todo {
					if qs[i].b < 0 {
						qs[i].v = sorted[lo]
					}
				}
			}

			return
		}

		mid := (lo + hi) >> 1
		var pathCnt int
		var b, c []int
		for _, i := range todo {
			q := qs[i]
			if q.b >= 0 {
				// update
				x, y, v := q.a, q.b, q.v
				if abs(v) < sorted[mid] {
					b = append(b, i)
					continue
				}
				// abs(v) >= sorted(mid)
				v = sign(v)
				pathCnt += v
				tr.update(tin[x], v)
				tr.update(tin[y], v)
				p := getLca(x, y)
				tr.update(tin[p], -v)
				if fa[p][0] != p {
					tr.update(tin[fa[p][0]], -v)
				}
				c = append(c, i)
			} else {
				// query
				x := -q.a
				cnt := tr.query(tin[x], tout[x]-1)
				if cnt == pathCnt {
					b = append(b, i)
				} else {
					c = append(c, i)
				}
			}
		}

		tr.reset()
		play(lo, mid, b)
		play(mid, hi, c)
	}

	play(-1, len(sorted), idx)

	var res []int
	for _, q := range qs {
		if q.b < 0 {
			res = append(res, q.v)
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
func sign(num int) int {
	if num > 0 {
		return 1
	}
	if num < 0 {
		return -1
	}
	return 0
}

type fenwick struct {
	val  []int
	todo []int
}

func newFenwick(n int) *fenwick {
	val := make([]int, n+1)
	return &fenwick{val, nil}
}

func (t *fenwick) update(i int, v int) {
	t.todo = append(t.todo, i)
	for i++; i < len(t.val); i += i & -i {
		t.val[i] += v
	}
}

func (t *fenwick) pre(i int) int {
	var res int
	for i++; i > 0; i -= i & -i {
		res += t.val[i]
	}
	return res
}

func (t *fenwick) query(l int, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func (t *fenwick) reset() {
	for _, i := range t.todo {
		for i++; i < len(t.val); i += i & -i {
			t.val[i] = 0
		}
	}
	t.todo = nil
}
