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
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

const H = 20

func solve(a []int, queries [][]int) []int {
	n := len(a)
	trs := make([]*Tree, H)
	for i := range H {
		trs[i] = NewTree(a, i)
	}
	var sum int

	for i := 0; i < n; i++ {
		for d := range H {
			if (a[i]>>d)&1 == 1 {
				j := trs[d].GetMin(i, n-1)
				sum += (j - i) * (1 << d)
			}
		}
	}

	update := func(d int, p int, v int) {
		if (a[p]>>d)&1 == (v>>d)&1 {
			return
		}
		if (a[p]>>d)&1 == 1 {
			// 原来是连起来的
			// l和r是离的最近的空的位置
			l := trs[d].GetMax(0, p)
			r := trs[d].GetMin(p, n-1)
			sum -= count(r-l-1) * (1 << d)
			// v = 0
			if l < p {
				sum += count(p-l-1) * (1 << d)
			}
			if p < r {
				sum += count(r-p-1) * (1 << d)
			}
			trs[d].Update(p, 0)
		} else {
			trs[d].Update(p, 1)
			l := trs[d].GetMax(0, p)
			r := trs[d].GetMin(p, n-1)
			if l < p {
				sum -= count(p-l-1) * (1 << d)
			}
			if p < r {
				sum -= count(r-p-1) * (1 << d)
			}
			sum += count(r-l-1) * (1 << d)
		}
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		p, x := cur[0], cur[1]
		p--
		for d := range H {
			update(d, p, x)
		}
		ans[i] = sum
		a[p] = x
	}

	return ans
}

func count(n int) int {
	return (n + 1) * n / 2
}

const inf = 1 << 60

type Tree struct {
	val [][2]int
}

func NewTree(a []int, d int) *Tree {
	n := len(a)
	val := make([][2]int, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			if (a[l]>>d)&1 == 0 {
				val[i][0] = l
				val[i][1] = l
			} else {
				val[i][0] = -1
				val[i][1] = n
			}
			return
		}

		mid := (l + r) / 2
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		val[i][0] = max(val[i*2+1][0], val[i*2+2][0])
		val[i][1] = min(val[i*2+1][1], val[i*2+2][1])
	}
	build(0, 0, n-1)
	return &Tree{val}
}

func (tr *Tree) Update(p int, v int) {
	n := len(tr.val) / 4

	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			if v == 0 {
				tr.val[i][0] = l
				tr.val[i][1] = l
			} else {
				tr.val[i][0] = -1
				tr.val[i][1] = n
			}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		tr.val[i][0] = max(tr.val[i*2+1][0], tr.val[i*2+2][0])
		tr.val[i][1] = min(tr.val[i*2+1][1], tr.val[i*2+2][1])
	}
	loop(0, 0, n-1)
}

func (tr *Tree) GetMax(L int, R int) int {
	var loop func(i int, l int, r int, L int, R int) int
	loop = func(i int, l int, r int, L int, R int) int {
		if L == l && R == r {
			return tr.val[i][0]
		}
		mid := (l + r) / 2
		if R <= mid {
			return loop(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return loop(i*2+2, mid+1, r, L, R)
		}
		return max(loop(i*2+1, l, mid, L, mid), loop(i*2+2, mid+1, r, mid+1, R))
	}

	n := len(tr.val) / 4

	return loop(0, 0, n-1, L, R)
}

func (tr *Tree) GetMin(L int, R int) int {
	var loop func(i int, l int, r int, L int, R int) int
	loop = func(i int, l int, r int, L int, R int) int {
		if L == l && R == r {
			return tr.val[i][1]
		}
		mid := (l + r) / 2
		if R <= mid {
			return loop(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return loop(i*2+2, mid+1, r, L, R)
		}
		return min(loop(i*2+1, l, mid, L, mid), loop(i*2+2, mid+1, r, mid+1, R))
	}

	n := len(tr.val) / 4

	return loop(0, 0, n-1, L, R)
}
