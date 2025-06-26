package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const inf = 1 << 60

func solve(a []int) int {
	n := len(a)

	f := func(arr []int) []int {
		fa := make([]int, n+1)
		for i := range n + 1 {
			fa[i] = i
		}
		find := func(u int) int {
			rt := u
			for fa[rt] != rt {
				rt = fa[rt]
			}
			for fa[u] != rt {
				fa[u], u = rt, fa[u]
			}
			return rt
		}
		var res []int
		for i, x := range arr {
			v := find(x)
			if v > 0 {
				res = append(res, i)
				fa[v] = v - 1
			}
		}
		return res
	}

	pre := f(a)
	slices.Reverse(a)
	suf := f(a)

	var ans int

	for i := range pre {
		ans += max(0, n-1-suf[i]-pre[i])
	}

	return ans
}

func solve1(a []int) int {
	n := len(a)

	t1 := NewSegTree(n, -inf, func(x, y int) int {
		return max(x, y)
	})

	t2 := NewSegTree(n, -inf, func(x, y int) int {
		return max(x, y)
	})

	for i := range n {
		t1.Update(i, i)
		t2.Update(i, i)
	}

	var L []int
	for i := range n {
		x := a[i] - 1
		v := t1.Query(0, x+1)
		if v >= 0 {
			// v存在
			L = append(L, i)
			t1.Update(v, -inf)
		}
	}

	var R []int
	for i := n - 1; i >= 0; i-- {
		x := a[i] - 1
		v := t2.Query(0, x+1)
		if v >= 0 {
			R = append(R, i)
			t2.Update(v, -inf)
		}
	}

	var ans int

	for i, p := range L {
		ans += max(0, R[i]-p)
	}

	return ans
}

type SegTree struct {
	f   func(int, int) int
	arr []int
	iv  int
	n   int
}

func NewSegTree(n int, iv int, f func(int, int) int) *SegTree {
	arr := make([]int, n*2)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{f, arr, iv, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.n
	t.arr[p] = v

	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Query(l int, r int) int {
	res := t.iv
	l += t.n
	r += t.n
	for l < r {
		if l&1 == 1 {
			res = t.f(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.f(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
