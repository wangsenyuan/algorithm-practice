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
	_, _, bus, res := drive(reader)
	fmt.Println(bus)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (a int, travelers [][]int, bus int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &a)
	travelers = make([][]int, n)
	for i := 0; i < n; i++ {
		var x, d int
		fmt.Fscan(reader, &x, &d)
		travelers[i] = []int{x, d}
	}
	bus, res = solve(a, travelers)
	return
}

type person struct {
	id int
	l  int
	r  int
}

const inf = 1 << 60

func solve(a int, travelers [][]int) (int, []int) {
	n := len(travelers)
	people := make([]person, n)
	for i, cur := range travelers {
		x, d := cur[0], cur[1]
		people[i] = person{i, x, x + d}
	}

	slices.SortFunc(people, func(a, b person) int {
		// 优先选择先下车的，同时下车的，优先选择先上车的, 后上车的还有机会
		return cmp.Or(a.r-b.r, a.id-b.id)
	})

	tr := NewSegTree(n)

	var res []int
	// k座的车能否满足条件
	check := func(k int) bool {
		tr.Reset()

		res = res[:0]

		var taken int

		for j, cur := range people {
			l := cur.l
			i := sort.Search(n, func(i int) bool {
				return people[i].r > l
			})
			i = tr.Get(0, i)
			if i >= 0 {
				// 让离l最近的人下车
				// 如果让其他人早下车了，貌似会造成不一样的结果
				taken--
				tr.Update(i, -1)
			}
			if taken < k {
				res = append(res, cur.id+1)
				taken++
				tr.Update(j, j)
			}
		}

		return len(res) >= a
	}

	l, r := 1, a
	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	check(r)
	return r, res[:a]
}

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = -1
	}
	return SegTree(arr)
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v

	for p > 1 {
		tr[p>>1] = max(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	var res int = -1
	for l < r {
		if l&1 == 1 {
			res = max(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func (tr SegTree) Reset() {
	for i := range tr {
		tr[i] = -1
	}
}
