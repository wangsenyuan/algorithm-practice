package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	modifications := make([][]int, m)
	for i := 0; i < m; i++ {
		modifications[i] = readNNums(reader, 3)
	}
	return solve(a, modifications)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func inv(a int) int {
	return pow(a, mod-2)
}

func solve(a []int, modifications [][]int) []int {
	n := len(a)
	I := make([]int, n+1)
	I[0] = 1
	for i := 1; i <= n; i++ {
		I[i] = inv(i)
	}

	tree := NewTree(a)

	for _, cur := range modifications {
		l, r, x := cur[0], cur[1], cur[2]
		l--
		r--
		m := r - l + 1
		tree.Update(l, r, pair{sub(1, I[m]), mul(I[m], x%mod)})
	}

	tree.GetResult(a)

	return a
}

type pair struct {
	first  int
	second int
}

type Tree struct {
	lazy []pair
	val  []int
	sz   int
}

func NewTree(arr []int) *Tree {
	n := len(arr)
	lazy := make([]pair, 4*n)
	val := make([]int, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		lazy[i] = zero
		if l == r {
			val[i] = arr[l] % mod
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
	}
	build(0, 0, n-1)
	return &Tree{lazy, val, n}
}

var zero = pair{1, 0}

func (t *Tree) update(i int, v pair) {
	t.val[i] = add(mul(t.val[i], v.first), v.second)
	f := mul(t.lazy[i].first, v.first)
	s := add(mul(t.lazy[i].second, v.first), v.second)
	t.lazy[i] = pair{f, s}
}

func (t *Tree) push(i int) {
	if t.lazy[i] != zero {
		t.update(2*i+1, t.lazy[i])
		t.update(2*i+2, t.lazy[i])
		t.lazy[i] = zero
	}
}

func (t *Tree) Update(L int, R int, v pair) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			t.update(i, v)
			return
		}
		t.push(i)
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
	}
	loop(0, 0, t.sz-1)
}

func (t *Tree) GetResult(res []int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			res[l] = t.val[i]
			return
		}
		t.push(i)
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
	}
	loop(0, 0, t.sz-1)
}
