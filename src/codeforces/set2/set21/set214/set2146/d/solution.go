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
		_, _, sum, res := drive(reader)
		fmt.Fprintln(writer, sum)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (l int, r int, sum int, res []int) {
	fmt.Fscan(reader, &l, &r)
	sum, res = solve(l, r)
	return
}

func solve(l int, r int) (sum int, a []int) {

	var tr [][2]int
	var cnt []int

	expand := func() int {
		tr = append(tr, [2]int{0, 0})
		cnt = append(cnt, 0)
		return len(tr) - 1
	}
	expand()

	for i := l; i <= r; i++ {
		var node int
		for d := range 30 {
			w := (i >> d) & 1
			if tr[node][w] == 0 {
				tr[node][w] = expand()
			}
			node = tr[node][w]
			cnt[node]++
		}
	}

	a = make([]int, r-l+1)

	for i := l; i <= r; i++ {
		var node int
		for d := range 30 {
			w := (i >> d) & 1
			w ^= 1
			if tr[node][w] == 0 || cnt[tr[node][w]] == 0 {
				w ^= 1
			}
			a[i-l] |= w << d
			node = tr[node][w]
			cnt[node]--
		}
		sum += a[i-l] | i
	}

	return
}

func solve1(l int, r int) (sum int, a []int) {
	n := r - l + 1
	a = make([]int, n)

	var f func(lo int, hi int, j int)

	f = func(lo int, hi int, j int) {
		if lo > hi {
			return
		}
		if lo == hi {
			a[lo-l] = hi
			return
		}
		mid := lo
		for mid+1 <= hi && ((mid+1)>>j)&1 == (lo>>j)&1 {
			mid++
		}
		if mid == hi {
			f(lo, hi, j-1)
			return
		}
		tl := mid + 1
		tr := mid

		for tl-1 >= lo && tr+1 <= hi {
			tl--
			tr++
			a[tl-l] = tr
			a[tr-l] = tl
		}
		f(lo, tl-1, j-1)
		f(tr+1, hi, j-1)
	}

	f(l, r, 29)

	for i, v := range a {
		sum += v | (l + i)
	}

	return
}
