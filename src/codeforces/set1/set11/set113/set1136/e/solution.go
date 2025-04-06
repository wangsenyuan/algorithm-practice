package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	k := readNNums(reader, n-1)
	m := readNum(reader)
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		queries[i] = readString(reader)
	}
	return solve(n, a, k, queries)
}

func solve(n int, a []int, k []int, queries []string) []int {
	// t[i] = sum k[0....i)
	pk := make([]int, n)
	b := make([]int, n)

	for i := 1; i < n; i++ {
		pk[i] = pk[i-1] + k[i-1]
	}

	for i := 0; i < n; i++ {
		b[i] = a[i] - pk[i]
	}
	// b[i] is sorted always
	tr := NewTree(b)
	tr2 := NewTree(pk)
	var ans []int

	for _, cur := range queries {
		s := []byte(cur)
		if cur[0] == '+' {
			var i, x int
			pos := readInt(s, 2, &i) + 1
			readInt(s, pos, &x)
			i--
			x += tr.Get(i, i)
			j := tr.LowerBound(x)
			tr.Update(i, j-1, x)
		} else {
			var l, r int
			pos := readInt(s, 2, &l) + 1
			readInt(s, pos, &r)
			l--
			r--
			tmp := tr.Get(l, r)
			tmp += tr2.Get(l, r)
			ans = append(ans, tmp)
		}
	}

	return ans
}

type Tree struct {
	val  []int // max value
	sum  []int
	lazy []int
	sz   int
}

const inf = 1e18

func NewTree(a []int) *Tree {
	n := len(a)
	val := make([]int, 4*n)
	sum := make([]int, 4*n)
	lazy := make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		lazy[i] = -inf
		if l == r {
			val[i] = a[l]
			sum[i] = a[l]
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		val[i] = val[2*i+2]
		sum[i] = sum[2*i+1] + sum[2*i+2]
	}
	build(0, 0, n-1)
	return &Tree{val, sum, lazy, n}
}

func (t *Tree) update(i int, l int, r int, v int) {
	t.lazy[i] = v
	t.val[i] = v
	t.sum[i] = v * (r - l + 1)
}

func (t *Tree) push(i int, l int, r int) {
	if t.lazy[i] == -inf {
		return
	}
	mid := (l + r) / 2
	t.update(2*i+1, l, mid, t.lazy[i])
	t.update(2*i+2, mid+1, r, t.lazy[i])
	t.lazy[i] = -inf
}

func (t *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if r < L || R < l {
			return
		}
		if L <= l && r <= R {
			t.update(i, l, r, v)
			return
		}
		t.push(i, l, r)
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
		t.val[i] = max(t.val[2*i+1], t.val[2*i+2])
		t.sum[i] = t.sum[2*i+1] + t.sum[2*i+2]
	}
	loop(0, 0, t.sz-1)
}

func (t *Tree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int

	loop = func(i int, l int, r int) int {
		if r < L || R < l {
			return 0
		}
		if L <= l && r <= R {
			return t.sum[i]
		}
		t.push(i, l, r)
		mid := (l + r) / 2
		return loop(2*i+1, l, mid) + loop(2*i+2, mid+1, r)
	}
	return loop(0, 0, t.sz-1)
}

func (t *Tree) LowerBound(v int) int {
	if t.val[0] < v {
		return t.sz
	}
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		t.push(i, l, r)

		mid := (l + r) / 2
		if t.val[2*i+1] >= v {
			return loop(2*i+1, l, mid)
		}
		return loop(2*i+2, mid+1, r)
	}
	return loop(0, 0, t.sz-1)
}
