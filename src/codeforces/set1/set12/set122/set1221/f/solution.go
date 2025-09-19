package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, best, rect := drive(reader)
	fmt.Println(best)
	fmt.Println(rect[0], rect[1], rect[2], rect[3])
}

func drive(reader *bufio.Reader) (points [][]int, best int, rect []int) {
	var n int
	fmt.Fscan(reader, &n)
	points = make([][]int, n)
	for i := 0; i < n; i++ {
		var x, y, c int
		fmt.Fscan(reader, &x, &y, &c)
		points[i] = []int{x, y, c}
	}
	best, rect = solve(points)
	return
}

type point struct {
	x int
	y int
	c int
}

func solve(points [][]int) (best int, rect []int) {
	var xs []int

	n := len(points)
	arr := make([]point, n)

	for i, p := range points {
		x, y, c := p[0], p[1], p[2]
		x, y = min(x, y), max(x, y)
		// x <= y
		xs = append(xs, x, y)
		arr[i] = point{x, y, c}
	}
	slices.Sort(xs)
	xs = slices.Compact(xs)
	xs = append(xs, 2e9)

	slices.SortFunc(arr, func(a, b point) int {
		return cmp.Or(a.y-b.y, a.x-b.x)
	})

	tr := NewTree(xs)

	best = -1 << 60

	for i, j := 0, 0; i < len(xs); i++ {
		for j < n && arr[j].y == xs[i] {
			p := sort.SearchInts(xs, arr[j].x)
			tr.Update(0, p, arr[j].c)
			j++
		}
		tmp := tr.Get(0, i)
		if tmp[0]-xs[i] > best {
			best = tmp[0] - xs[i]
			lo := tmp[1]
			rect = []int{xs[lo], xs[lo], xs[i], xs[i]}
		}
	}

	return
}

type Tree struct {
	val  [][2]int
	lazy []int
	sz   int
}

func NewTree(arr []int) *Tree {
	n := len(arr)
	val := make([][2]int, 4*n)
	lazy := make([]int, 4*n)
	t := &Tree{val, lazy, n}
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			val[i][0] = arr[l]
			val[i][1] = l
			return
		}
		mid := (l + r) >> 1
		f(i*2+1, l, mid)
		f(i*2+2, mid+1, r)
		t.pull(i)
	}
	f(0, 0, n-1)
	return t
}

func (t *Tree) update(i int, v int) {
	t.val[i][0] += v
	t.lazy[i] += v
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.update(2*i+1, t.lazy[i])
		t.update(2*i+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func merge(a [2]int, b [2]int) [2]int {
	if a[0] >= b[0] {
		return a
	}
	return b
}

func (t *Tree) pull(i int) {
	x := merge(t.val[i*2+1], t.val[i*2+2])
	t.val[i][0] = x[0]
	t.val[i][1] = x[1]
}

func (t *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.update(i, v)
			return
		}

		t.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			f(i*2+1, l, mid, L, R)
		} else if mid < L {
			f(i*2+2, mid+1, r, L, R)
		} else {
			f(i*2+1, l, mid, L, mid)
			f(i*2+2, mid+1, r, mid+1, R)
		}
		t.pull(i)
	}

	f(0, 0, t.sz-1, L, R)
}

func (t *Tree) Get(L int, R int) [2]int {
	var f func(i int, l int, r int, L int, R int) [2]int
	f = func(i int, l int, r int, L int, R int) [2]int {
		if l == L && r == R {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return f(i*2+2, mid+1, r, L, R)
		}
		return merge(f(i*2+1, l, mid, L, mid), f(i*2+2, mid+1, r, mid+1, R))
	}
	return f(0, 0, t.sz-1, L, R)
}
