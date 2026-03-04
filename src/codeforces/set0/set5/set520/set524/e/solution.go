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
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range res {
		if v {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func drive(reader *bufio.Reader) []bool {
	var n, m, k, q int
	fmt.Fscan(reader, &n, &m, &k, &q)
	rocks := make([][]int, k)
	for i := range k {
		rocks[i] = make([]int, 2)
		fmt.Fscan(reader, &rocks[i][0], &rocks[i][1])
	}
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 4)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2], &queries[i][3])
	}
	return solve(n, m, rocks, queries)
}

func solve(n int, m int, rocks [][]int, queries [][]int) []bool {
	slices.SortFunc(rocks, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	atRow := make([][]int, n+1)
	atCol := make([][]int, m+1)
	for i, cur := range rocks {
		r, c := cur[0], cur[1]
		atRow[r] = append(atRow[r], i)
		atCol[c] = append(atCol[c], i)
	}

	cols := make([]*node, m+1)
	for c := 1; c <= m; c++ {
		tr := cols[c-1]
		for _, i := range atCol[c] {
			r := rocks[i][0]
			tr = set(tr, 1, n, r, c)
		}
		cols[c] = tr
	}
	rows := make([]*node, n+1)
	for r := 1; r <= n; r++ {
		tr := rows[r-1]
		for _, i := range atRow[r] {
			c := rocks[i][1]
			tr = set(tr, 1, m, c, r)
		}
		rows[r] = tr
	}

	check := func(x1 int, y1 int, x2 int, y2 int) bool {
		v := get(cols[y2], 1, n, x1, x2)
		if v >= y1 {
			return true
		}

		w := get(rows[x2], 1, m, y1, y2)
		if w >= x1 {
			return true
		}

		return false
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		ans[i] = check(cur[0], cur[1], cur[2], cur[3])
	}

	return ans
}

type node struct {
	left, right *node
	val         int
}

func set(n *node, lo, hi, p, v int) *node {
	if lo == hi {
		return &node{val: v}
	}
	mid := (lo + hi) / 2
	res := &node{}
	if n != nil {
		*res = *n
	}
	if p <= mid {
		res.left = set(res.left, lo, mid, p, v)
	} else {
		res.right = set(res.right, mid+1, hi, p, v)
	}
	res.val = nodeVal(res.left, res.right)
	return res
}

func get(n *node, lo, hi, l, r int) int {
	if n == nil {
		return -1
	}
	if lo == l && hi == r {
		return n.val
	}
	mid := (lo + hi) / 2
	if r <= mid {
		return get(n.left, lo, mid, l, r)
	}
	if mid < l {
		return get(n.right, mid+1, hi, l, r)
	}
	return min(get(n.left, lo, mid, l, mid), get(n.right, mid+1, hi, mid+1, r))
}

func nodeVal(left, right *node) int {
	a, b := -1, -1
	if left != nil {
		a = left.val
	}
	if right != nil {
		b = right.val
	}
	return min(a, b)
}
