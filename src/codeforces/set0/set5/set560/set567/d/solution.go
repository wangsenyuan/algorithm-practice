package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k, a, m int
	fmt.Fscan(reader, &n, &k, &a, &m)
	x := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &x[i])
	}
	return solve(n, k, a, x)
}

func solve(n int, k int, a int, x []int) int {

	marked := make([]bool, n+2)

	check := func(mid int) bool {
		for i := 0; i < mid; i++ {
			marked[x[i]] = true
		}
		marked[0] = true
		marked[n+1] = true
		var sum int

		prev := 0
		for i := 1; i <= n+1; i++ {
			if marked[i] {
				sum += (i - prev) / (a + 1)
				prev = i
			}
		}

		for i := range mid {
			marked[x[i]] = false
		}

		return sum < k
	}

	m := len(x)
	res := sort.Search(m, check)

	if res == m && !check(res) {
		return -1
	}
	return res
}

func solve1(n int, k int, a int, x []int) int {
	set := NewSet(n + 2)
	set.Set(0)
	set.Set(n + 1)
	// 可以放置cnt个ship, 必须保证 cnt >= k
	cnt := (n + 1) / (a + 1)

	for i, v := range x {
		l := set.QueryMax(0, v)
		r := set.QueryMin(v, n+1)
		tmp := (r - l) / (a + 1)
		cnt -= tmp
		set.Set(v)
		cnt += (v - l) / (a + 1)
		cnt += (r - v) / (a + 1)
		if cnt < k {
			return i + 1
		}
	}
	return -1
}

type Set struct {
	tr1 *SegTree
	tr2 *SegTree
}

func NewSet(n int) *Set {
	tr1 := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})
	tr2 := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})
	return &Set{tr1, tr2}
}

func (set *Set) QueryMin(L int, R int) int {
	return set.tr2.Find(L, R+1)
}

func (set *Set) QueryMax(L int, R int) int {
	return set.tr1.Find(L, R+1)
}

func (set *Set) Clear(pos int) {
	set.tr1.Update(pos, -1)
	set.tr2.Update(pos, set.tr2.initValue)
}

func (set *Set) Set(pos int) {
	set.tr1.Update(pos, pos)
	set.tr2.Update(pos, pos)
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func NewSegTree(n int, iv int, fn func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{arr, n, iv, fn}
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = v
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := tree.initValue

	for l < r {
		if l&1 == 1 {
			ans = tree.fn(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = tree.fn(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
