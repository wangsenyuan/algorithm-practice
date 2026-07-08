package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
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
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	arr := slices.Clone(a)
	for _, cur := range queries {
		if cur[0] == 1 {
			arr = append(arr, cur[2])
		} else {
			arr = append(arr, cur[1], cur[2])
		}
	}

	slices.Sort(arr)
	arr = slices.Compact(arr)
	m := len(arr)

	sum := make(SegTree, 2*m)
	cnt := make(SegTree, 2*m)

	for _, v := range a {
		i := sort.SearchInts(arr, v)
		sum.Update(i, v)
		cnt.Update(i, 1)
	}

	var res []int
	for _, cur := range queries {
		if cur[0] == 1 {
			i := cur[1] - 1
			v := cur[2]
			j := sort.SearchInts(arr, a[i])
			sum.Update(j, -a[i])
			cnt.Update(j, -1)
			j = sort.SearchInts(arr, v)
			sum.Update(j, v)
			cnt.Update(j, 1)
			a[i] = v
		} else {
			l := sort.SearchInts(arr, cur[1])
			r := sort.SearchInts(arr, cur[2])
			if l <= r {
				c1 := cnt.Query(0, l)
				s2 := sum.Query(l, r)
				c2 := cnt.Query(r, m)
				res = append(res, c1*cur[1]+s2+c2*cur[2])
			} else {
				// l > r, max(l, min(r, a[i]))
				// 如果 a[i] > r => r => l
				// 如果 a[i] < r => a[i] => l
				res = append(res, len(a)*cur[1])
			}
		}
	}

	return res
}

type SegTree []int

func (st SegTree) Update(i int, v int) {
	n := len(st) / 2
	i += n
	st[i] += v
	for i > 1 {
		st[i>>1] += v
		i >>= 1
	}
}

func (st SegTree) Query(l int, r int) int {
	n := len(st) / 2
	l += n
	r += n
	ans := 0
	for l < r {
		if l&1 == 1 {
			ans += st[l]
			l++
		}
		if r&1 == 1 {
			r--
			ans += st[r]
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
