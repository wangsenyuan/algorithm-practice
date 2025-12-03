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
	res := drive(reader)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1:len(s)-1])
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) []int {
	var n, m, p int
	fmt.Fscan(reader, &n, &m, &p)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	return solve(p, a, b)
}

func solve(p int, a []int, b []int) []int {
	b1 := slices.Clone(b)
	slices.Sort(b1)
	b1 = slices.Compact(b1)
	k := len(b1)

	f1 := make([]int, k+1)
	m := len(b)

	for i := range m {
		j := sort.SearchInts(b1, b[i])
		f1[j]++
	}

	f2 := make([]int, k+1)

	get := func(x int) int {
		j := sort.SearchInts(b1, x)
		if j < k && b1[j] == x {
			return j
		}
		return k
	}

	n := len(a)

	e := min(n, p)

	var res []int
	for i := range e {
		l, r := i, i
		for r < n {
			j := get(a[r])
			r += p

			f2[j]++
			for f2[j] > f1[j] {
				j1 := get(a[l])
				f2[j1]--
				l += p
			}
			if r-l == m*p {
				res = append(res, l+1)
			}
		}
		for l < r {
			j1 := get(a[l])
			f2[j1]--
			l += p
		}
	}

	slices.Sort(res)

	return res
}
