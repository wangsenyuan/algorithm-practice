package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		for _, cur := range res {
			s := fmt.Sprintf("%v", cur)
			writer.WriteString(s[1 : len(s)-1])
			writer.WriteByte('\n')
		}
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(a, queries)
}

type query struct {
	id int
	l  int
	r  int
}

func solve1(a []int, queries [][]int) [][]int {
	// n := len(a)
	pos := make(map[int][]int)
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}

	check := func(x int, l int, r int) bool {
		cnt := (r - l + 1) / 3
		i := sort.SearchInts(pos[x], l)
		j := sort.SearchInts(pos[x], r+1)
		return j-i > cnt
	}

	mark := make([]int, len(a))
	var marker int
	play := func(l int, r int) []int {
		var res []int

		marker++
		for range 50 {
			i := rand.IntN(r-l+1) + l
			if mark[i] == marker {
				continue
			}
			mark[i] = marker
			if len(res) > 0 && res[0] == a[i] {
				continue
			}

			if check(a[i], l, r) {
				res = append(res, a[i])
			}
			if len(res) == 2 {
				break
			}
		}
		if len(res) > 0 {
			slices.Sort(res)
			return res
		}
		return []int{-1}
	}

	ans := make([][]int, len(queries))

	for i, cur := range queries {
		ans[i] = play(cur[0]-1, cur[1]-1)
	}

	return ans
}

func solve(a []int, queries [][]int) [][]int {
	pos := make(map[int][]int)
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}

	check := func(x int, l int, r int) bool {
		cnt := (r - l + 1) / 3
		i := sort.SearchInts(pos[x], l)
		j := sort.SearchInts(pos[x], r+1)
		return j-i > cnt
	}

	tr := NewTree(a, 2)

	find := func(l int, r int) []int {
		tmp := tr.Get(l, r)

		var res []int
		for _, x := range tmp {
			if check(x.first, l, r) {
				res = append(res, x.first)
			}
		}
		if len(res) == 0 {
			return []int{-1}
		}
		slices.Sort(res)
		return res
	}

	ans := make([][]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0]-1, cur[1]-1)
	}

	return ans
}

type pair struct {
	first  int
	second int
}

type Tree struct {
	arr [][]pair
	k   int
	sz  int
}

func merge(a []pair, b []pair, k int) []pair {
	// merge a & b to get a c, at most k size
	c := slices.Clone(a)
	// add b to a
	for _, x := range b {
		if x.second == 0 {
			continue
		}
		found := false
		for i := range len(c) {
			if c[i].first == x.first {
				c[i].second += x.second
				found = true
				break
			}
		}
		if found {
			continue
		}
		c = append(c, x)
		for len(c) > k {
			d := c[0].second
			for _, y := range c[1:] {
				d = min(d, y.second)
			}
			var j int
			for i := range len(c) {
				c[i].second -= d
				if c[i].second > 0 {
					c[j] = c[i]
					j++
				}
			}
			c = c[:j]
		}
	}
	return c
}

func NewTree(a []int, k int) *Tree {
	n := len(a)
	arr := make([][]pair, 2*n)

	for i := n; i < 2*n; i++ {
		arr[i] = []pair{{a[i-n], 1}}
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = merge(arr[i*2], arr[i*2+1], k)
	}
	return &Tree{arr, k, n}
}

func (tr *Tree) Get(l int, r int) []pair {
	l += tr.sz
	r += tr.sz + 1

	var res []pair
	for l < r {
		if l&1 == 1 {
			res = merge(res, tr.arr[l], tr.k)
			l++
		}
		if r&1 == 1 {
			r--
			res = merge(res, tr.arr[r], tr.k)
		}
		l >>= 1
		r >>= 1
	}
	return res
}
