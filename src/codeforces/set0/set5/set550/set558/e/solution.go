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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func drive(reader *bufio.Reader) string {
	_, m := readTwoNums(reader)
	s := readString(reader)
	ops := make([][]int, m)
	for i := range m {
		ops[i] = readNNums(reader, 3)
	}
	return solve(s, ops)
}

func solve(s string, ops [][]int) string {
	trs := make([]*Tree, 26)
	n := len(s)

	for i := range 26 {
		trs[i] = NewTree(n)
	}

	for i := range n {
		x := int(s[i] - 'a')
		trs[x].Update(i, i, 1)
	}

	cnt := make([]int, 26)
	for _, op := range ops {
		l, r, k := op[0]-1, op[1]-1, op[2]
		for i := range 26 {
			cnt[i] = trs[i].Get(l, r)
			trs[i].Update(l, r, 0)
		}
		if k == 1 {
			// 升序
			pos := l
			for i := range 26 {
				if cnt[i] > 0 {
					trs[i].Update(pos, pos+cnt[i]-1, 1)
					pos += cnt[i]
				}
			}
		} else {
			// 降序
			pos := l
			for i := 25; i >= 0; i-- {
				if cnt[i] > 0 {
					trs[i].Update(pos, pos+cnt[i]-1, 1)
					pos += cnt[i]
				}
			}
		}
	}

	res := make([]byte, n)

	for i := range n {
		for j := range 26 {
			cnt := trs[j].Get(i, i)
			if cnt == 1 {
				res[i] = byte(j + 'a')
				break
			}
		}
	}

	return string(res)
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	for i := range lazy {
		lazy[i] = -1
	}
	return &Tree{val, lazy, n}
}

func (t *Tree) apply(i int, l int, r int, v int) {
	t.val[i] = v * (r - l + 1)
	t.lazy[i] = v
}

func (t *Tree) push(i int, l int, r int) {

	if l < r && t.lazy[i] != -1 {
		mid := (l + r) >> 1
		t.apply(i*2+1, l, mid, t.lazy[i])
		t.apply(i*2+2, mid+1, r, t.lazy[i])
		t.lazy[i] = -1
	}
}

func (t *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if L == l && R == r {
			t.apply(i, l, r, v)
			return
		}
		t.push(i, l, r)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}

		t.val[i] = t.val[i*2+1] + t.val[i*2+2]
	}

	f(0, 0, t.sz-1, L, R)
}

func (t *Tree) Get(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int

	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return t.val[i]
		}
		t.push(i, l, r)
		mid := (l + r) >> 1
		var res int
		if L <= mid {
			res = f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			res += f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		return res
	}
	return f(0, 0, t.sz-1, L, R)
}
