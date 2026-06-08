package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	d := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &d[i])
	}
	return solve(n, d)
}

const mod = 998244353

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, d []int) int {
	dp := NewTree(n)
	dp.set(n-1, 1)
	dp.set(n-2, 1)
	for i := n - 3; i >= 0; i-- {
		j := i + d[i]
		tmp := dp.get(j, j)
		// 如果当前位置作为最大值，那么只能使用j作为第二大的数
		if d[i] == d[i+1] {
			// 可以使用方案1，在位置i上放置*小*的数
			// 那么到目前为止共有n-i个数，最大值和第二大的数，位置不变，当前i这里可以放置的个数 = n - i - 2个数
			dp.update(i+1, n-1, n-i-2)
		} else {
			// 只能有两个情况，一种是i是最大值，一种是j是最大值
			dp.update(i+1, n-1, 0)
		}
		dp.set(i, tmp)
		// 但是如果当前位置作为第二大的数呢，怎么处理呢？
		dp.set(j, add(tmp, dp.get(j, j)))
	}

	return dp.get(0, n-1)
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	for i := range 4 * n {
		lazy[i] = 1
		val[i] = 0
	}
	sz := n
	return &Tree{val, lazy, sz}
}

func (tr *Tree) apply(i int, v int) {
	tr.val[i] = mul(tr.val[i], v)
	tr.lazy[i] = mul(tr.lazy[i], v)
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 1 {
		tr.apply(i*2+1, tr.lazy[i])
		tr.apply(i*2+2, tr.lazy[i])
		tr.lazy[i] = 1
	}
}

func (tr *Tree) pull(i int) {
	tr.val[i] = add(tr.val[i*2+1], tr.val[i*2+2])
}

func (tr *Tree) update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		tr.pull(i)
	}

	f(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) get(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr.val[i]
		}
		tr.push(i)
		var res int
		mid := (l + r) >> 1
		if L <= mid {
			res = add(res, f(i*2+1, l, mid, L, min(mid, R)))
		}
		if mid < R {
			res = add(res, f(i*2+2, mid+1, r, max(mid+1, L), R))
		}
		return res
	}
	return f(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) set(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = v
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.pull(i)
	}
	f(0, 0, tr.sz-1)
}
