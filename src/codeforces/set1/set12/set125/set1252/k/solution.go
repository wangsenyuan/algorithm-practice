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

	for _, pair := range drive(reader) {
		fmt.Fprintln(writer, pair[0], pair[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	var s string
	fmt.Fscan(reader, &s)
	ops := make([][]int, q)
	for i := range q {
		var typ int
		fmt.Fscan(reader, &typ)
		if typ == 1 {
			ops[i] = make([]int, 3)
			ops[i][0] = 1
			fmt.Fscan(reader, &ops[i][1], &ops[i][2])
		} else {
			ops[i] = make([]int, 5)
			ops[i][0] = 2
			fmt.Fscan(reader, &ops[i][1], &ops[i][2], &ops[i][3], &ops[i][4])
		}
	}
	return solve(n, s, ops)
}

func solve(n int, s string, ops [][]int) [][]int {
	tr := build(s)

	var ans [][]int
	for _, op := range ops {
		if op[0] == 1 {
			tr.update(op[1]-1, op[2])
		} else {
			l, r, a, b := op[1]-1, op[2], op[3], op[4]
			tmp := tr.query(l, r)
			tmp = mat{{a, b}}.mul(tmp)
			ans = append(ans, []int{tmp[0][0], tmp[0][1]})
		}
	}
	return ans
}

const mod = 1_000_000_007

func add(nums ...int) int {
	var res int
	for _, num := range nums {
		res = (res + num) % mod
	}
	return res
}

func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res = (res * num) % mod
	}
	return res
}

type mat [][]int

func (a mat) mul(b mat) mat {
	n := len(a)
	m := len(a[0])
	k := len(b[0])
	c := make(mat, n)
	for i := range c {
		c[i] = make([]int, k)
	}
	for i := range n {
		for j := range k {
			for l := range m {
				c[i][j] = add(c[i][j], mul(a[i][l], b[l][j]))
			}
		}
	}
	return c
}

func convert(c byte) mat {
	if c == 'A' {
		return mat{
			{1, 0},
			{1, 1},
		}
	}
	return mat{
		{1, 1},
		{0, 1},
	}
}

type tree struct {
	val  []mat
	lazy []int
}

func build(s string) *tree {
	n := len(s)
	val := make([]mat, 4*n)
	lazy := make([]int, 4*n)
	var f func(i, l, r int)
	f = func(i, l, r int) {
		if l+1 == r {
			val[i] = convert(s[l])
			return
		}
		mid := (l + r) >> 1
		f(i<<1, l, mid)
		f(i<<1|1, mid, r)
		val[i] = val[i<<1].mul(val[i<<1|1])
	}
	f(1, 0, n)
	return &tree{val, lazy}
}

func (t *tree) toggle(i int) {
	t.lazy[i] ^= 1
	// [[a, b], [c, d]] => 	[[d, c], [b, a]]
	t.val[i] = mat{
		{t.val[i][1][1], t.val[i][1][0]},
		{t.val[i][0][1], t.val[i][0][0]},
	}
}

func (t *tree) push(i int) {
	if t.lazy[i] != 0 {
		t.toggle(i << 1)
		t.toggle(i<<1 | 1)
		t.lazy[i] = 0
	}
}

func (t *tree) update(L int, R int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.toggle(i)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		if L < mid {
			f(i<<1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i<<1|1, mid, r, max(mid, L), R)
		}
		t.val[i] = t.val[i<<1].mul(t.val[i<<1|1])
	}
	n := len(t.val) / 4
	f(1, 0, n, L, R)
}

func (t *tree) query(L int, R int) mat {
	var f func(i int, l int, r int, L int, R int) mat
	f = func(i int, l int, r int, L int, R int) mat {
		if l == L && r == R {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			return f(i<<1, l, mid, L, R)
		}
		if L >= mid {
			return f(i<<1|1, mid, r, L, R)
		}
		return f(i<<1, l, mid, L, mid).mul(f(i<<1|1, mid, r, mid, R))
	}
	n := len(t.val) / 4
	return f(1, 0, n, L, R)
}
