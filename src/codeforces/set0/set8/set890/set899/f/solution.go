package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) string {
	_, m := readTwoNums(reader)
	s := readString(reader)
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		queries[i] = readString(reader)
	}
	return solve(s, queries)
}

func get(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a')
	}
	if c >= 'A' && c <= 'Z' {
		return 26 + int(c-'A')
	}
	return 52 + int(c-'0')
}

func solve(s string, queries []string) string {
	n := len(s)

	tr := NewTree(s)

	for _, q := range queries {
		buf := []byte(q)
		var l, r int
		pos := readInt(buf, 0, &l) + 1
		pos = readInt(buf, pos, &r) + 1
		x := get(buf[pos])
		i := tr.FindPos(int32(l))
		j := tr.FindPos(int32(r))
		tr.Update(i, j, x)
	}

	var buf []byte
	for i := range n {
		if tr.Get(i) == 1 {
			buf = append(buf, s[i])
		}
	}
	return string(buf)
}

type Tree struct {
	cnt  [][62]int32
	sum  []int32
	lazy []int
	n    int
}

func NewTree(s string) *Tree {
	n := len(s)
	tr := new(Tree)
	tr.n = n
	tr.cnt = make([][62]int32, 4*n)
	tr.sum = make([]int32, 4*n)
	tr.lazy = make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			tr.cnt[i][get(s[l])]++
			tr.sum[i] = 1
			return
		}
		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		tr.pull(i)
	}
	build(0, 0, n-1)
	return tr
}

func (t *Tree) pull(i int) {
	for j := range 62 {
		t.cnt[i][j] = t.cnt[i*2+1][j] + t.cnt[i*2+2][j]
	}
	t.sum[i] = t.sum[i*2+1] + t.sum[i*2+2]
}

func (t *Tree) update(i int, v int) {
	if t.sum[i] == 0 {
		return
	}
	for j := range 62 {
		if (v>>j)&1 == 1 {
			t.sum[i] -= t.cnt[i][j]
			t.cnt[i][j] = 0
		}
	}
	t.lazy[i] |= v
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.update(i*2+1, t.lazy[i])
		t.update(i*2+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if L == l && R == r {
			t.update(i, 1<<v)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(L, mid+1), R)
		}
		t.pull(i)
	}
	f(0, 0, t.n-1, L, R)
}

func (t *Tree) FindPos(k int32) int {
	var f func(i int, l int, r int, k int32) int
	f = func(i int, l int, r int, k int32) int {
		if l == r {
			return l
		}
		t.push(i)
		mid := (l + r) >> 1
		if t.sum[i*2+1] >= k {
			return f(i*2+1, l, mid, k)
		}
		return f(i*2+2, mid+1, r, k-t.sum[i*2+1])
	}
	return f(0, 0, t.n-1, k)
}

func (t *Tree) Get(k int) int32 {
	var f func(i int, l int, r int) int32
	f = func(i int, l int, r int) int32 {
		if l == r || t.sum[i] == 0 {
			return t.sum[i]
		}
		t.push(i)
		mid := (l + r) >> 1
		if k <= mid {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid+1, r)
	}
	return f(0, 0, t.n-1)
}
