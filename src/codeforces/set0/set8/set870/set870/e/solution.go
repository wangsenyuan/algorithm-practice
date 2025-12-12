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
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	return solve(points)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

type point struct {
	id int
	x  int
	y  int
}

func solve(points [][]int) int {
	n := len(points)
	arr := make([]point, n)
	for i := range n {
		arr[i] = point{id: i, x: points[i][0], y: points[i][1]}
	}

	slices.SortFunc(arr, func(a, b point) int {
		return cmp.Or(a.x-b.x, a.y-b.y)
	})

	py := make(map[int]int)

	adj := make([][]int, n)

	connect := func(u int, v int) {
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	for i := 0; i < n; {
		j := i
		for i < n && arr[i].x == arr[j].x {
			if i > j {
				connect(arr[i-1].id, arr[i].id)
			}
			if v, ok := py[arr[i].y]; ok {
				connect(v, arr[i].id)
			}
			py[arr[i].y] = arr[i].id
			i++
		}
	}

	xs := make(map[int]bool)
	ys := make(map[int]bool)

	marked := make([]bool, n)

	var dfs func(p int, u int) bool
	dfs = func(p int, u int) bool {
		if marked[u] {
			// found cycle
			return true
		}
		xs[points[u][0]] = true
		ys[points[u][1]] = true

		marked[u] = true
		cycle := false
		for _, v := range adj[u] {
			if p != v {
				if dfs(u, v) {
					cycle = true
				}
			}
		}

		return cycle
	}

	res := 1

	for i := range n {
		if !marked[i] {
			cycle := dfs(-1, i)
			cur := pow(2, len(xs)+len(ys))
			if !cycle {
				cur = add(cur, mod-1)
			}
			res = mul(res, cur)
			xs = make(map[int]bool)
			ys = make(map[int]bool)
		}
	}

	return res
}
