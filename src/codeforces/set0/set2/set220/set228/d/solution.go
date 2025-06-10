package main

import (
	"bufio"
	"bytes"
	"fmt"
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
		var k int
		pos := readInt(s, 0, &k) + 1
		if k == 1 {
			queries[i] = make([]int, 3)
		} else {
			queries[i] = make([]int, 4)
		}
		queries[i][0] = k
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	trs := make([]*Tree, 7)
	for i := 2; i <= 6; i++ {
		trs[i] = NewTree(a, i)
	}

	change := func(i int, v int) {
		for z := 2; z <= 6; z++ {
			trs[z].Update(i, v)
		}
	}

	res := make([]int, 0, len(queries))

	for _, cur := range queries {
		if cur[0] == 1 {
			p, v := cur[1], cur[2]
			p--
			change(p, v)
		} else {
			l, r, z := cur[1], cur[2], cur[3]
			l--
			r--
			res = append(res, trs[z].Get(l, r))
		}
	}
	return res
}

type Tree struct {
	val [][]int
	sz  int
	z   int
}

func getValue(i int, w int, z int) int {
	// 下标从1开始
	if i%w == 0 {
		return 2
	}
	if i%w <= z {
		return i % w
	}
	return 2*z - i%w
}

func merge(c []int, a []int, b []int, z int, x int) {
	w := 2 * (z - 1)
	for j := 0; j < w; j++ {
		c[j] = a[j] + b[(j+x)%w]
	}
}

func NewTree(a []int, z int) *Tree {
	w := 2 * (z - 1)
	n := len(a)
	val := make([][]int, 4*n)

	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		val[i] = make([]int, w)
		if l == r {
			for j := 0; j < w; j++ {
				val[i][j] = getValue(j+1, w, z) * a[l]
			}
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		merge(val[i], val[2*i+1], val[2*i+2], z, mid-l+1)
	}
	build(0, 0, n-1)
	return &Tree{val, n, z}
}

func (tr *Tree) Update(p int, v int) {
	w := 2 * (tr.z - 1)
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			for j := 0; j < w; j++ {
				tr.val[i][j] = getValue(j+1, w, tr.z) * v
			}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		merge(tr.val[i], tr.val[2*i+1], tr.val[2*i+2], tr.z, mid-l+1)
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) Get(L int, R int) int {
	w := 2 * (tr.z - 1)
	var loop func(i int, l int, r int, L int, R int) []int
	loop = func(i int, l int, r int, L int, R int) []int {
		if l == L && r == R {
			return tr.val[i]
		}
		mid := (l + r) / 2
		if R <= mid {
			return loop(2*i+1, l, mid, L, R)
		}
		if mid < L {
			return loop(2*i+2, mid+1, r, L, R)
		}
		a := loop(2*i+1, l, mid, L, mid)
		b := loop(2*i+2, mid+1, r, mid+1, R)
		c := make([]int, w)
		merge(c, a, b, tr.z, mid-L+1)
		return c
	}
	res := loop(0, 0, tr.sz-1, L, R)
	return res[0]
}
