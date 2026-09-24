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
	var buf bytes.Buffer

	for _, x := range drive(reader) {
		fmt.Fprintf(&buf, "%d\n", x)
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	ops := make([][]int, m)
	for i := range m {
		ops[i] = make([]int, 4)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1], &ops[i][2], &ops[i][3])
	}
	return solve(n, ops)
}

func solve(n int, ops [][]int) []int {
	todo := make([]int, len(ops))
	var nums []int
	for i, op := range ops {
		todo[i] = i
		if op[0] == 1 {
			nums = append(nums, op[3])
		}
	}
	slices.Sort(nums)
	nums = slices.Compact(nums)

	tr := NewTree(n)

	var play func(todo []int, lo int, hi int)
	play = func(todo []int, lo int, hi int) {
		for _, i := range todo {
			if ops[i][0] == 2 {
				goto next
			}
		}
		return
	next:
		if lo+1 == hi {
			for _, i := range todo {
				if ops[i][0] == 2 {
					ops[i][3] = nums[lo]
				}
			}
			return
		}
		mid := (lo + hi) / 2
		x := nums[mid]
		var b, c []int
		for _, i := range todo {
			l, r, v := ops[i][1]-1, ops[i][2]-1, ops[i][3]
			if ops[i][0] == 1 {
				// insert l r v
				if v >= x {
					tr.Update(l, r, 1)
					b = append(b, i)
				} else {
					c = append(c, i)
				}
			} else {
				// query l, r, k
				cnt := tr.Query(l, r)
				if cnt >= v {
					b = append(b, i)
				} else {
					ops[i][3] -= cnt
					c = append(c, i)
				}
			}
		}
		for _, i := range todo {
			if ops[i][0] == 1 && ops[i][3] >= x {
				l, r := ops[i][1]-1, ops[i][2]-1
				tr.Update(l, r, -1)
			}
		}
		play(c, lo, mid)
		play(b, mid, hi)
	}

	play(todo, 0, len(nums))

	var res []int
	for _, cur := range ops {
		if cur[0] == 2 {
			res = append(res, cur[3])
		}
	}
	return res
}

type Tree struct {
	sum  []int
	lazy []int
}

func NewTree(n int) *Tree {
	sum := make([]int, n*4)
	lazy := make([]int, n*4)

	return &Tree{sum, lazy}
}

func (tr *Tree) apply(i int, l int, r int, v int) {
	tr.sum[i] += (r - l + 1) * v
	tr.lazy[i] += v
}

func (tr *Tree) push(i int, l int, r int) {
	if tr.lazy[i] != 0 {
		mid := (l + r) / 2
		tr.apply(i*2+1, l, mid, tr.lazy[i])
		tr.apply(i*2+2, mid+1, r, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) pull(i int) {
	tr.sum[i] = tr.sum[i*2+1] + tr.sum[i*2+2]
}

func (tr *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, l, r, v)
			return
		}
		tr.push(i, l, r)
		mid := (l + r) / 2
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(L, mid+1), R)
		}
		tr.pull(i)
	}
	n := len(tr.sum) / 4
	f(0, 0, n-1, L, R)
}

func (tr *Tree) Query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr.sum[i]
		}
		tr.push(i, l, r)
		mid := (l + r) / 2
		var res int
		if L <= mid {
			res += f(i*2+1, l, mid, L, min(R, mid))
		}
		if mid < R {
			res += f(i*2+2, mid+1, r, max(L, mid+1), R)
		}
		return res
	}
	n := len(tr.sum) / 4
	return f(0, 0, n-1, L, R)
}
