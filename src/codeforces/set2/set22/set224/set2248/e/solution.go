package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n, d int64
	var m int
	fmt.Fscan(reader, &n, &m, &d)
	p := make([]int64, m)
	r := make([]int64, m)
	for i := range m {
		fmt.Fscan(reader, &p[i], &r[i])
	}
	return solve(n, d, p, r)
}

func solve(n, d int64, p []int64, r []int64) string {
	m := len(p)
	if m == 0 {
		return "NO"
	}

	pref := make([]int64, m)
	for i, v := range r {
		pref[i] = v
		if i > 0 {
			pref[i] += pref[i-1]
		}
	}
	total := n*d + pref[m-1]

	score := func(x int64) int64 {
		q, rem := x/n, x%n
		res := q * total
		res += rem * d
		i := sort.Search(len(p), func(i int) bool {
			return p[i] > rem
		})
		if i > 0 {
			res += pref[i-1]
		}
		return res
	}

	for i := 0; i < m; i++ {
		x := p[i]
		sx := x*d + pref[i]
		for j := 0; j < m; j++ {
			y := p[j]
			if sx+y*d+pref[j] > score(x+y+1) {
				return "YES"
			}
		}
	}
	return "NO"
}
