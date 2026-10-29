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
		var t, l, r, x int
		fmt.Fscan(reader, &t, &l, &r, &x)
		queries[i] = []int{t, l, r, x}
	}
	return solve(a, queries)
}

const mod = 998244353

func add(nums ...int) int {
	var res int
	for _, num := range nums {
		res += num
		res %= mod
	}
	return res
}
func sub(a, b int) int {
	b %= mod
	return add(a, mod-b)
}
func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res = res * num % mod
	}
	return res
}

func pow(a int, b int) int {
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

func solve(a []int, queries [][]int) []int {
	var ans []int
	t := build(a)
	for _, cur := range queries {
		l, r, x := cur[1], cur[2], cur[3]
		if cur[0] == 1 {
			t.update(l-1, r, x)
		} else {
			ans = append(ans, t.query(l-1, r)[x])
		}
	}
	return ans
}

// (a + x) ^^ 2 - a ^^ 2 = x * (2 * a + x)
// (a + x) ^^ 2 - a ^^ 2 + (b + x) ^^ 2 - b ^^ 2 = x * (2 * a + x + 2 * b + x)
// (a + x) ^^ 3 - a ^^ 3 = x * (3 * a ^^ 2 + 3 * a * x + x ^^ 2)
// sum2 += 2 * sum1 * x + len * x ^^ 2
// sum3 += 3 * sum2 * x + 3 * sum1 * x ^^ 2 + len * x ^^ 3
// sum4 += 4 * sum3 * x + 6 * sum2 * x ^^ 2 + 4 * sum1 * x ^^ 3 + len * x ^^ 4

type node struct {
	sum [5]int
	tag int
	len int
}

func (n *node) add(x int) {
	n.sum[4] = add(n.sum[4], mul(4, n.sum[3], x), mul(6, n.sum[2], mul(x, x)), mul(4, n.sum[1], pow(x, 3)), mul(n.len, pow(x, 4)))
	n.sum[3] = add(n.sum[3], mul(3, n.sum[2], x), mul(3, n.sum[1], pow(x, 2)), mul(n.len, pow(x, 3)))
	n.sum[2] = add(n.sum[2], mul(2, n.sum[1], x), mul(n.len, pow(x, 2)))
	n.sum[1] = add(n.sum[1], mul(n.len, x))
	n.tag = add(n.tag, x)
}

type tree []node

func build(a []int) tree {
	n := len(a)
	t := make(tree, n*4)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		t[i].len = r - l
		if l+1 == r {
			t[i].add(a[l])
			return
		}
		mid := (l + r) >> 1
		f(i*2+1, l, mid)
		f(i*2+2, mid, r)
		for j := range 5 {
			t[i].sum[j] = add(t[i*2+1].sum[j], t[i*2+2].sum[j])
		}
	}
	f(0, 0, n)
	return t
}

func (t tree) push(i int) {
	if t[i].tag != 0 {
		t[i*2+1].add(t[i].tag)
		t[i*2+2].add(t[i].tag)
		t[i].tag = 0
	}
}

func (t tree) update(L int, R int, x int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t[i].add(x)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		if L < mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid, r, max(mid, L), R)
		}
		for j := range 5 {
			t[i].sum[j] = add(t[i*2+1].sum[j], t[i*2+2].sum[j])
		}
	}
	n := len(t) / 4
	f(0, 0, n, L, R)
}

func (t tree) query(L int, R int) [5]int {
	var f func(i int, l int, r int, L int, R int) [5]int
	f = func(i int, l int, r int, L int, R int) [5]int {
		if l == L && r == R {
			return t[i].sum
		}
		t.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid <= L {
			return f(i*2+2, mid, r, L, R)
		}
		a := f(i*2+1, l, mid, L, mid)
		b := f(i*2+2, mid, r, mid, R)
		var c [5]int
		for j := range 5 {
			c[j] = add(a[j], b[j])
		}
		return c
	}
	return f(0, 0, len(t)/4, L, R)
}
