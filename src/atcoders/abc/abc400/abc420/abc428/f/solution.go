package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	ws := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &ws[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		var typ int
		fmt.Fscan(reader, &typ)
		if typ == 3 {
			var x int
			fmt.Fscan(reader, &x)
			queries[i] = []int{typ, x}
		} else {
			var v int
			fmt.Fscan(reader, &v)
			queries[i] = []int{typ, v}
		}
	}
	return solve(ws, queries)
}

func solve(ws []int, queries [][]int) []int {
	// n := len(ws)

	type state struct {
		v    int
		dir  int
		l, r int
	}

	var arr []state
	var stack []int

	var ans []int

	n := len(ws)

	tr := make(SegTree, 2*n)
	for i := range 2 * n {
		tr[i] = -1
	}

	play := func(x int) int {
		// 找到最小的v, 包含x的区间
		check := func(v int) bool {
			i := tr.Query(v, n)
			if i < 0 {
				return x < ws[v]
			}
			cur := arr[i]
			l, r := cur.l, cur.r
			if cur.dir == 1 {
				r = l + ws[v]
			} else {
				l = r - ws[v]
			}

			return l <= x && x < r
		}

		return n - sort.Search(n, check)
	}

	for _, cur := range queries {
		if cur[0] == 3 {
			ans = append(ans, play(cur[1]))
			continue
		}
		v := cur[1] - 1
		for len(stack) > 0 && arr[stack[len(stack)-1]].v < v {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			arr = append(arr, state{v: v, dir: cur[0], l: 0, r: ws[v]})
		} else {
			prev := arr[stack[len(stack)-1]]
			l, r := prev.l, prev.r
			if prev.dir == 1 {
				// 当前也是左端靠齐的
				r = l + ws[v]
			} else {
				l = r - ws[v]
			}
			arr = append(arr, state{v: v, dir: cur[0], l: l, r: r})
		}
		tr.Update(v, len(arr)-1)
		stack = append(stack, len(arr)-1)
	}

	return ans
}

type SegTree []int

func (s SegTree) Update(p int, v int) {
	n := len(s) / 2
	p += n
	s[p] = v
	for p > 1 {
		s[p>>1] = max(s[p], s[p^1])
		p >>= 1
	}
}

func (s SegTree) Query(l, r int) int {
	n := len(s) / 2
	l += n
	r += n
	ans := -1
	for l < r {
		if l&1 == 1 {
			ans = max(ans, s[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = max(ans, s[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
