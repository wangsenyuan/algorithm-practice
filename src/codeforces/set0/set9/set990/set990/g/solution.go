package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) [][]int {
	var n int32
	fmt.Fscan(reader, &n)
	a := make([]int32, n)
	for i := int32(0); i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int32, n-1)
	for i := int32(0); i < n-1; i++ {
		edges[i] = make([]int32, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(a, edges)
}

type edge struct {
	u int
	v int
}

func solve(a []int32, es [][]int32) [][]int {
	// n := int32(len(a))

	x := slices.Max(a)
	edges := make([][]edge, x+1)

	for _, e := range es {
		u, v := int(e[0]-1), int(e[1]-1)
		w := gcd(a[u], a[v])
		edges[w] = append(edges[w], edge{u: u, v: v})
	}
	n := len(a)
	time := make([]int, n)
	fa := make([]int, n)
	sz := make([]int, n)
	var now int

	find := func(x int) int {
		y := x
		for {
			if time[y] != now {
				time[y] = now
				fa[y] = y
				sz[y] = 1
			}
			if fa[y] == y {
				break
			}
			y = fa[y]
		}
		for fa[x] != x {
			fa[x], x = y, fa[x]
		}
		return y
	}

	union := func(u int, v int) bool {
		if sz[u] < sz[v] {
			u, v = v, u
		}
		sz[u] += sz[v]
		fa[v] = u
		return true
	}

	f := make([]int, x+1)

	for i := x; i > 0; i-- {
		now++
		for j := i; j <= x; j += i {
			for _, e := range edges[j] {
				u, v := e.u, e.v
				pu, pv := find(u), find(v)
				if pu != pv {
					f[i] += sz[pu] * sz[pv]
					union(pu, pv)
				}
			}
		}
		for j := 2 * i; j <= x; j += i {
			f[i] -= f[j]
		}
	}

	for _, v := range a {
		f[v]++
	}

	var res [][]int

	for i := 1; i <= int(x); i++ {
		if f[i] > 0 {
			res = append(res, []int{i, f[i]})
		}
	}
	return res
}
func gcd(a, b int32) int32 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func last[T any](arr []T) T {
	return arr[len(arr)-1]
}
