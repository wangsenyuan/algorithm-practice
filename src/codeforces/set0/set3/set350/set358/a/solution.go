package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	return solve(x)
}

const inf = 1 << 60

func solve(x []int) string {
	arr := slices.Clone(x)
	slices.Sort(arr)
	arr = slices.Compact(arr)
	n := len(arr)
	t1 := build(n, -1, func(a int, b int) int {
		return max(a, b)
	})
	t2 := build(n, inf, func(a int, b int) int {
		return min(a, b)
	})

	for i := 0; i+1 < len(x); i++ {
		l, _ := slices.BinarySearch(arr, x[i])
		r, _ := slices.BinarySearch(arr, x[i+1])
		l, r = min(l, r), max(l, r)
		if l+1 <= r-1 {
			w := t2.query(l+1, r)
			if w < l {
				return "yes"
			}
			v := t1.query(l+1, r)
			if r < v {
				return "yes"
			}
		}
		w := t1.query(l, l+1)
		if w < r {
			t1.update(l, r)
		}
		v := t2.query(r, r+1)
		if l < v {
			t2.update(r, l)
		}
	}
	return "no"
}

type segtree struct {
	arr       []int
	initValue int
	f         func(int, int) int
	sz        int
}

func build(n int, iv int, f func(int, int) int) *segtree {
	arr := make([]int, 2*n)
	for i := range 2 * n {
		arr[i] = iv
	}
	return &segtree{arr, iv, f, n}
}

func (s *segtree) update(i int, v int) {
	i += s.sz
	s.arr[i] = v
	for i > 1 {
		s.arr[i>>1] = s.f(s.arr[i], s.arr[i^1])
		i >>= 1
	}
}

func (s *segtree) query(l int, r int) int {
	res := s.initValue
	l += s.sz
	r += s.sz
	for l < r {
		if l&1 == 1 {
			res = s.f(res, s.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = s.f(res, s.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
