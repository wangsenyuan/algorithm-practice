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
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 3)
		for j := range 3 {
			fmt.Fscan(reader, &queries[i][j])
		}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	arr := slices.Clone(a)
	slices.Sort(arr)
	arr = slices.Compact(arr)

	bit := make(BIT, len(a)+2)

	var play func(idx []int, l int, r int, todo []int)

	ans := make([]int, len(queries))

	play = func(idx []int, l int, r int, todo []int) {
		if l == r || len(todo) == 0 {
			return
		}
		if l+1 == r || len(idx) == 1 {
			for _, i := range todo {
				ans[i] = a[idx[0]]
			}
			return
		}

		var small []int
		var big []int
		mid := (l + r) / 2
		x := arr[mid]
		// 需要知道在区间[ql, qr]中间有多少个比 <= x
		for _, i := range idx {
			if a[i] < x {
				small = append(small, i)
				bit.update(i, 1)
			} else {
				big = append(big, i)
			}
		}

		var todoSmall []int
		var todoBig []int
		for _, i := range todo {
			l, r, k := queries[i][0]-1, queries[i][1]-1, queries[i][2]
			cnt := bit.query(l, r)
			if cnt >= k {
				todoSmall = append(todoSmall, i)
			} else {
				todoBig = append(todoBig, i)
				// 这个减去不能丢掉
				queries[i][2] -= cnt
			}
		}

		for _, i := range idx {
			if a[i] < x {
				bit.update(i, -1)
			}
		}
		play(small, l, mid, todoSmall)
		play(big, mid, r, todoBig)
	}

	idx := make([]int, len(a))
	for i := range len(a) {
		idx[i] = i
	}
	todo := make([]int, len(queries))
	for i := range len(queries) {
		todo[i] = i
	}
	play(idx, 0, len(arr), todo)

	return ans
}

type BIT []int

func (bit BIT) update(p int, v int) {
	for p++; p < len(bit); p += p & -p {
		bit[p] += v
	}
}

func (bit BIT) pre(p int) int {
	var res int
	for p++; p > 0; p -= p & -p {
		res += bit[p]
	}
	return res
}

func (bit BIT) query(l int, r int) int {
	return bit.pre(r) - bit.pre(l-1)
}
