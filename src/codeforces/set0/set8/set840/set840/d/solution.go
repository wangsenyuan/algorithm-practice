package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	debug.SetGCPercent(-1)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var l, r, k int
		fmt.Fscan(reader, &l, &r, &k)
		queries[i] = []int{l, r, k}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)

	ans := make([]int, len(queries))

	arr := make([][]int, n+1)
	for i, cur := range queries {
		r := cur[1]
		arr[r] = append(arr[r], i)
	}

	trs := make([]*node, n+1)
	trs[0] = build(0, n)

	for i := 1; i <= n; i++ {
		trs[i] = trs[i-1].update(a[i-1])

		for _, id := range arr[i] {
			l, k := queries[id][0], queries[id][2]
			l--
			d := (i-l)/k + 1

			ans[id] = -1

			for j := d; j <= i-l; j += d {
				w := query(trs[l], trs[i], j)
				if w.first >= d {
					ans[id] = w.second
					break
				}
			}
		}
	}

	return ans
}

type node struct {
	lf, rg *node
	l, r   int
	cnt    int
}

func build(l int, r int) *node {
	node := &node{l: l, r: r, cnt: 0}
	if l < r {
		mid := (l + r) >> 1
		node.lf = build(l, mid)
		node.rg = build(mid+1, r)
	}
	return node
}

func (n node) update(p int) *node {
	if n.l == n.r {
		n.cnt++
		return &n
	}
	mid := (n.l + n.r) >> 1
	if p <= mid {
		n.lf = n.lf.update(p)
	} else {
		n.rg = n.rg.update(p)
	}
	n.cnt = n.lf.cnt + n.rg.cnt
	return &n
}

type pair struct {
	first  int
	second int
}

// 在区间[a...b]找第k大的数
func query(a *node, b *node, k int) pair {
	if a.l == a.r {
		return pair{b.cnt - a.cnt, a.l}
	}

	tmp := b.lf.cnt - a.lf.cnt

	if tmp >= k {
		return query(a.lf, b.lf, k)
	}
	return query(a.rg, b.rg, k-tmp)
}
