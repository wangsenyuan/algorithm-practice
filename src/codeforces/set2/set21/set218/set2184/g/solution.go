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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		for _, ans := range res {
			fmt.Fprintln(writer, ans)
		}
	}
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
		var t, l, r int
		fmt.Fscan(reader, &t, &l, &r)
		queries[i] = []int{t, l, r}
	}
	return solve(a, queries)
}

const inf = 1 << 60

func solve(a []int, queries [][]int) []int {
	n := len(a)

	tr := make([]int, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			tr[i] = a[l]
		} else {
			mid := (l + r) >> 1
			build(i*2+1, l, mid)
			build(i*2+2, mid+1, r)
			tr[i] = min(tr[i*2+1], tr[i*2+2])
		}
	}
	build(0, 0, n-1)

	var update func(i int, l int, r int, p int, v int)
	update = func(i int, l int, r int, p int, v int) {
		if l == r {
			tr[i] = v
		} else {
			mid := (l + r) >> 1
			if p <= mid {
				update(i*2+1, l, mid, p, v)
			} else {
				update(i*2+2, mid+1, r, p, v)
			}
			// 这个最小值是最右边的那个
			tr[i] = min(tr[i*2+1], tr[i*2+2])
		}
	}

	var query func(i int, l int, r int, L int, R int) int
	query = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr[i]
		}
		mid := (l + r) >> 1
		if R <= mid {
			return query(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return query(i*2+2, mid+1, r, L, R)
		}
		return min(query(i*2+1, l, mid, L, mid), query(i*2+2, mid+1, r, mid+1, R))
	}

	play := func(l int, r int) int {
		lo, hi := l, n
		for lo < hi {
			mid := (lo + hi) >> 1
			v := query(0, 0, n-1, l, mid)

			if v <= mid-l {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		if hi <= r && query(0, 0, n-1, l, hi) == hi-l {
			return 1
		}
		return 0
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			i, x := cur[1]-1, cur[2]
			update(0, 0, n-1, i, x)
		} else {
			l, r := cur[1]-1, cur[2]-1
			ans = append(ans, play(l, r))
		}
	}
	return ans
}
