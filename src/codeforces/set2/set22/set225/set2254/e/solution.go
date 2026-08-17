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
		res := drive(reader)
		if len(res) == 0 {
			fmt.Fprintln(writer, -1)
			continue
		}
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(b)
}

func solve(b []int) []int {
	slices.Sort(b)
	n := len(b)
	t := make(tree, 4*n)
	for i := range n {
		t.update(i, b[i])
	}
	first := t.upperBound(1)
	if first[0] < 0 {
		return nil
	}
	a := make([]int, n)
	a[0] = first[1]
	t.update(first[0], -inf)

	for i := 1; i < n; i++ {
		// a[i] = a[i-1] + x > 0
		// x > -a[i-1]
		tmp := t.upperBound(-a[i-1] + 1)
		if tmp[0] < 0 {
			return nil
		}
		a[i] = a[i-1] + tmp[1]
		t.update(tmp[0], -inf)
	}

	return a
}

const inf = 1 << 60

type tree []int

func (t tree) update(pos int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			t[i] = v
			return
		}
		mid := (l + r) >> 1
		if pos < mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid, r)
		}
		t[i] = max(t[i*2+1], t[i*2+2])
	}
	n := len(t) / 4
	f(0, 0, n)
}

func (t tree) upperBound(num int) []int {
	if t[0] < num {
		return []int{-1, num - 1}
	}
	var f func(i int, l int, r int) []int
	f = func(i int, l int, r int) []int {
		if l+1 == r {
			return []int{l, t[i]}
		}
		mid := (l + r) >> 1
		if t[i*2+1] >= num {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid, r)
	}
	n := len(t) / 4
	return f(0, 0, n)
}
