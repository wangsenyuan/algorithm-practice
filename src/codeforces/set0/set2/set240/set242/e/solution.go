package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	queries := make([][]int, m)
	for i := range m {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			queries[i] = make([]int, 3)
		} else {
			queries[i] = make([]int, 4)
		}
		var pos int
		for j := range len(queries[i]) {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	return solve(a, queries)
}

const X = 1e6

func solve(a []int, queries [][]int) []int {
	h := bits.Len(uint(X))
	trs := make([]*Tree, h)
	for i := range h {
		trs[i] = NewTree(a, i)
	}

	find := func(l int, r int) int {
		var res int
		for i := range h {
			cnt := trs[i].Get(l, r)
			res += cnt * (1 << i)
		}
		return res
	}

	update := func(l int, r int, x int) {
		for i := range h {
			if (x>>i)&1 == 1 {
				trs[i].Xor(l, r)
			}
		}
	}

	var ans []int

	for _, cur := range queries {
		l, r := cur[1]-1, cur[2]-1
		if cur[0] == 1 {
			// sum
			ans = append(ans, find(l, r))
		} else {
			update(l, r, cur[3])
		}
	}
	return ans
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(arr []int, d int) *Tree {
	n := len(arr)
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = (arr[l] >> d) & 1
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		val[i] = val[2*i+1] + val[2*i+2]
	}
	build(0, 0, n-1)
	return &Tree{val, lazy, n}
}

func (tr *Tree) xor(i int, l int, r int) {
	tr.val[i] = r - l + 1 - tr.val[i]
	tr.lazy[i] ^= 1
}

func (tr *Tree) push(i int, l int, r int) {
	if tr.lazy[i] == 1 {
		mid := (l + r) / 2
		tr.xor(2*i+1, l, mid)
		tr.xor(2*i+2, mid+1, r)
		tr.lazy[i] = 0
	}
}

func (tr *Tree) pull(i int) {
	tr.val[i] = tr.val[2*i+1] + tr.val[2*i+2]
}

// 反转区间L...R
func (tr *Tree) Xor(L int, R int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			tr.xor(i, l, r)
			return
		}
		tr.push(i, l, r)
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
		tr.pull(i)
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) Get(L int, R int) int {
	var loop func(i int, l int, r int) int

	loop = func(i int, l int, r int) int {
		if R < l || r < L {
			return 0
		}
		if L <= l && r <= R {
			return tr.val[i]
		}
		tr.push(i, l, r)
		mid := (l + r) / 2
		return loop(2*i+1, l, mid) + loop(2*i+2, mid+1, r)
	}
	return loop(0, 0, tr.sz-1)
}
