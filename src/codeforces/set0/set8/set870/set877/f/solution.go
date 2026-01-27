package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, w := range res {
		fmt.Fprintln(writer, w)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	t := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &t[i])
	}
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(k, t, a, queries)
}

func solve(k int, t []int, a []int, queries [][]int) []int {
	n := len(a)
	var nums []int
	pref := make([]int, n+1)
	for i := range n {
		v := a[i]
		if t[i] == 2 {
			v *= -1
		}
		pref[i+1] = pref[i] + v
		nums = append(nums, pref[i+1], pref[i+1]-k, pref[i+1]+k)
	}
	nums = append(nums, 0, -k, k)
	slices.Sort(nums)
	nums = slices.Compact(nums)
	I := make([]int, n+1)
	L := make([]int, n+1)
	R := make([]int, n+1)
	for i := range n + 1 {
		I[i] = sort.SearchInts(nums, pref[i])
		L[i] = sort.SearchInts(nums, pref[i]-k)
		R[i] = sort.SearchInts(nums, pref[i]+k)
	}

	type query struct {
		l  int
		r  int
		id int
	}

	qs := make([]query, len(queries))
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		qs[i] = query{l: l - 1, r: r, id: i}
	}

	blk_sz := int(math.Sqrt(float64(n)))

	slices.SortFunc(qs, func(a query, b query) int {
		if a.r/blk_sz != b.r/blk_sz {
			return a.r - b.r
		}
		if (a.r/blk_sz)&1 == 0 {
			return a.l - b.l
		}
		return b.l - a.l
	})

	cnt := make([]int, len(nums))

	var ans int

	addR := func(i int) {
		ans += cnt[L[i]]
		cnt[I[i]]++
	}
	delR := func(i int) {
		cnt[I[i]]--
		ans -= cnt[L[i]]
	}

	addL := func(i int) {
		ans += cnt[R[i]]
		cnt[I[i]]++
	}
	delL := func(i int) {
		cnt[I[i]]--
		ans -= cnt[R[i]]
	}

	res := make([]int, len(queries))

	var l, r int
	addR(0)

	for _, cur := range qs {
		for r < cur.r {
			r++
			addR(r)
		}
		for l > cur.l {
			l--
			addL(l)
		}
		for l < cur.l {
			delL(l)
			l++
		}
		for r > cur.r {
			delR(r)
			r--
		}
		res[cur.id] = ans
	}

	return res
}
