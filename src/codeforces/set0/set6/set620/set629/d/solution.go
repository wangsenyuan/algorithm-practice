package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	cakes := make([][]int, n)
	for i := 0; i < n; i++ {
		var r, h int
		fmt.Fscan(reader, &r, &h)
		cakes[i] = []int{r, h}
	}
	return solve(n, cakes)
}

type pair struct {
	first  int
	second int
}

func solve(n int, cakes [][]int) float64 {
	arr := make([]pair, n)

	for i, cake := range cakes {
		r, h := cake[0], cake[1]
		arr[i] = pair{r * r * h, i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, b.second-a.second)
	})

	s := NewSegTree(n)

	for _, cur := range arr {
		i := cur.second
		v := s.Get(0, i)
		r, h := cakes[i][0], cakes[i][1]
		v += r * r * h
		s.Update(i, v)
	}

	best := s.Get(0, n)

	return float64(best) * math.Pi
}

type SegTree []int

func NewSegTree(n int) SegTree {
	return make(SegTree, 2*n)
}

func (s SegTree) Update(i int, v int) {
	n := len(s) >> 1
	i += n
	s[i] = v
	for i > 1 {
		s[i>>1] = max(s[i], s[i^1])
		i >>= 1
	}
}

func (s SegTree) Get(l int, r int) int {
	n := len(s) >> 1
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, s[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, s[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
