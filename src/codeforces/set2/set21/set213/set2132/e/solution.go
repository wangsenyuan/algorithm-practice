package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		ans := drive(reader)
		for _, v := range ans {
			fmt.Fprintln(writer, v)
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var x, y, z int
		fmt.Fscan(reader, &x, &y, &z)
		queries[i] = []int{x, y, z}
	}
	return solve(a, b, queries)
}

func solve(a []int, b []int, queries [][]int) []int {
	s1 := play(a)
	s2 := play(b)

	n := len(a)
	m := len(b)
	todo := make([][]int, n+m+1)

	for i, q := range queries {
		todo[q[2]] = append(todo[q[2]], i)
	}

	ans := make([]int, len(queries))

	get := func(arr []int, i int) int {
		if i < 0 {
			return 0
		}
		return arr[i]
	}

	for i, j := 0, 0; i < n || j < m; {
		if j == m || i < n && a[i] >= b[j] {
			i++
		} else {
			j++
		}
		k := i + j
		for _, id := range todo[k] {
			x, y := queries[id][0], queries[id][1]
			if x < i {
				ans[id] = get(s1, x-1) + get(s2, k-x-1)
			} else if y < j {
				ans[id] = get(s1, k-y-1) + get(s2, y-1)
			} else {
				ans[id] = get(s1, i-1) + get(s2, j-1)
			}
		}
	}

	return ans
}

func play(arr []int) []int {
	slices.Sort(arr)
	slices.Reverse(arr)
	sum := slices.Clone(arr)
	for i := 1; i < len(arr); i++ {
		sum[i] += sum[i-1]
	}
	return sum
}

func solve1(a []int, b []int, queries [][]int) []int {
	a = play(a)
	b = play(b)

	get := func(x int, y int) int {
		var res int
		if x > 0 {
			if x <= len(a) {
				res += a[x-1]
			} else {
				res += a[len(a)-1]
			}
		}
		if y > 0 {
			if y <= len(b) {
				res += b[y-1]
			} else {
				res += b[len(b)-1]
			}
		}
		return res
	}

	find := func(x int, y int, z int) int {
		// l + y <= z => l <= z - y
		l, r := max(0, z-y), min(x, z)
		for r-l >= 3 {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			s1 := get(m1, z-m1)
			s2 := get(m2, z-m2)
			if s1 > s2 {
				r = m2
			} else {
				l = m1
			}
		}
		var res int
		for i := l; i <= r; i++ {
			res = max(res, get(i, z-i))
		}
		return res
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = find(query[0], query[1], query[2])
	}
	return ans

}
