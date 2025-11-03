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

	var n, d, h int
	fmt.Fscan(reader, &n, &d, &h)
	res := solve(n, d, h)

	if len(res) == 0 {
		fmt.Fprintln(writer, -1)
		return
	}
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func solve(n int, d int, h int) [][]int {
	if n == 2 {
		if d == 1 && h == 1 {
			return [][]int{{1, 2}}
		}
		return nil
	}
	if h > d {
		return nil
	}
	if h == 1 {
		if d == 2 {
			var res [][]int
			for i := 2; i <= n; i++ {
				res = append(res, []int{1, i})
			}
			return res
		}
		// d == 1 and n > 2
		return nil
	}
	h++
	d++
	if h == d {
		var res [][]int
		for i := 1; i+1 <= d; i++ {
			res = append(res, []int{i, i + 1})
		}
		for i := d + 1; i <= n; i++ {
			res = append(res, []int{2, i})
		}
		return res
	}

	if d == n {
		var res [][]int
		for i := 1; i+1 <= n; i++ {
			res = append(res, []int{i, i + 1})
		}
		// h < d
		ok := false
		for i := 2; i < n; i++ {
			if max(i, n-i+1) == h {
				res[i-1][0] = 1
				res[i-2][1] = 1
				res[0][0] = i
				ok = true
				break
			}
		}
		if !ok {
			return nil
		}
		return res
	}

	// h < d
	var res [][]int
	// 2, 3,... d, d + 1
	root := -1
	for i := 2; i <= d; i++ {
		res = append(res, []int{i, i + 1})
		if max(i-2+1, d+1-i+1) == h {
			root = i
		}
	}

	if root == -1 {
		return nil
	}

	// change 1 with root
	res[root-2][0] = 1
	if root-3 >= 0 {
		res[root-3][1] = 1
	}

	for i := d + 2; i <= n; i++ {
		res = append(res, []int{1, i})
	}

	res = append(res, []int{1, root})

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
