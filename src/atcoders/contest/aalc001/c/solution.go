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

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
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
		var t, pos int
		fmt.Fscan(reader, &t, &pos)
		if t == 1 {
			var x int
			fmt.Fscan(reader, &x)
			queries[i] = []int{t, pos, x}
			continue
		}
		var threshold int
		fmt.Fscan(reader, &threshold)
		queries[i] = []int{t, pos, threshold}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	t := build(a)
	var ans []int
	n := len(a)
	for _, cur := range queries {
		l, r := cur[1]-1, cur[2]
		switch cur[0] {
		case 1:
			t.update(l, r)
			a[l] = r
		case 2:
			if a[l] >= r {
				ans = append(ans, l+1)
			} else {
				w := t.firstPosAfter(l, r)
				if w != -1 {
					ans = append(ans, w+1)
				} else {
					ans = append(ans, n+1)
				}
			}
		case 3:
			if a[l] >= r {
				ans = append(ans, l+1)
			} else {
				w := t.lastPosBefore(l, r)
				if w != -1 {
					ans = append(ans, w+1)
				} else {
					ans = append(ans, 0)
				}
			}
		}
	}
	return ans
}

type node struct {
	sum      int
	bestPref int
	bestSuf  int
}

func merge(a node, b node) node {
	return node{
		sum:      a.sum + b.sum,
		bestPref: max(a.bestPref, a.sum+b.bestPref),
		bestSuf:  max(b.bestSuf, b.sum+a.bestSuf),
	}
}

type tree []node

func build(a []int) tree {
	n := len(a)
	t := make(tree, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			t[i] = node{sum: a[l], bestPref: a[l], bestSuf: a[l]}
			return
		}
		mid := (l + r) >> 1
		f(i*2+1, l, mid)
		f(i*2+2, mid, r)
		t[i] = merge(t[i*2+1], t[i*2+2])
	}
	f(0, 0, n)
	return t
}

func (t tree) update(pos int, x int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l+1 == r {
			t[i] = node{sum: x, bestPref: x, bestSuf: x}
			return
		}
		mid := (l + r) >> 1
		if pos < mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid, r)
		}
		t[i] = merge(t[i*2+1], t[i*2+2])
	}
	f(0, 0, len(t)/4)
}

func (t tree) firstPosAfter(pos int, threshold int) int {
	var sum int

	var g func(i int, l int, r int) int
	g = func(i int, l int, r int) int {
		if l+1 == r {
			return l
		}
		mid := (l + r) >> 1
		if t[i*2+1].bestPref+sum >= threshold {
			return g(i*2+1, l, mid)
		}
		sum += t[i*2+1].sum
		return g(i*2+2, mid, r)
	}
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l+1 == r {
			sum += t[i].sum
			return -1
		}
		mid := (l + r) >> 1
		if mid <= pos {
			return f(i*2+2, mid, r)
		}
		// pos < mid
		ans := f(i*2+1, l, mid)
		if ans != -1 {
			return ans
		}
		if t[i*2+2].bestPref+sum >= threshold {
			return g(i*2+2, mid, r)
		}
		sum += t[i*2+2].sum
		return -1
	}
	return f(0, 0, len(t)/4)
}

func (t tree) lastPosBefore(pos int, threshold int) int {
	var sum int
	var g func(i int, l int, r int) int
	g = func(i int, l int, r int) int {
		if l+1 == r {
			return l
		}
		mid := (l + r) >> 1
		if t[i*2+2].bestSuf+sum >= threshold {
			return g(i*2+2, mid, r)
		}
		sum += t[i*2+2].sum
		return g(i*2+1, l, mid)
	}
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l+1 == r {
			sum += t[i].sum
			return -1
		}
		mid := (l + r) >> 1
		if pos < mid {
			return f(i*2+1, l, mid)
		}
		// mid <= pos
		ans := f(i*2+2, mid, r)
		if ans != -1 {
			return ans
		}
		if t[i*2+1].bestSuf+sum >= threshold {
			return g(i*2+1, l, mid)
		}
		sum += t[i*2+1].sum
		return -1
	}
	return f(0, 0, len(t)/4)
}
