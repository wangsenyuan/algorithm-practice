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

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []string {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []string {
	s1 := NewSegTree(n, -inf, func(a, b int) int {
		return max(a, b)
	})
	s2 := NewSegTree(n, inf, func(a, b int) int {
		return min(a, b)
	})

	ans := make([]string, len(queries))

	for i, cur := range queries {
		l, r := cur[0]-1, cur[1]-1
		if l > r {
			l, r = r, l
		}
		ans[i] = "Yes"

		if l+1 < r {
			// 离l最近的点
			w := s1.Get(l+1, r)
			if w > r {
				ans[i] = "No"
				continue
			}
			w = s2.Get(l+1, r)
			if w < l {
				ans[i] = "No"
				continue
			}
		}

		s1.Update(l, r)
		s2.Update(r, l)
	}

	return ans
}

const inf = 1 << 60

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
