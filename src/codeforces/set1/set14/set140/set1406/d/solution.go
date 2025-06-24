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
	qs := make([][]int, m)
	for i := range m {
		qs[i] = readNNums(reader, 3)
	}
	return solve(a, qs)
}

func solve(a []int, queries [][]int) []int {
	tr := NewTree(a)
	m := len(queries)
	ans := make([]int, m+1)
	n := len(a)

	calc := func(sum int) int {
		if sum >= 0 {
			return (sum + 1) / 2
		}
		if abs(sum)%2 == 0 {
			return sum / 2
		}
		return (sum + 1) / 2
	}

	ans[0] = calc(a[n-1] + tr.sum[0])

	for i, cur := range queries {
		l, r, v := cur[0]-1, cur[1]-1, cur[2]
		tr.Update(l, r, v)
		// a[0]和 a[n-1]似乎也有影响
		if r == n-1 {
			a[n-1] += v
		}
		ans[i+1] = calc(a[n-1] + tr.sum[0])
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}

type Tree struct {
	val  [][2]int
	sum  []int
	lazy []int
	sz   int
}

func (tr *Tree) merge(i int) {
	l := tr.val[2*i+1]
	r := tr.val[2*i+2]
	tr.sum[i] = tr.sum[2*i+1] + tr.sum[2*i+2] + max(l[1]-r[0], 0)
	tr.val[i][0] = l[0]
	tr.val[i][1] = r[1]
}

func NewTree(a []int) *Tree {
	n := len(a)
	tr := new(Tree)
	tr.val = make([][2]int, 4*n)
	tr.sum = make([]int, 4*n)
	tr.lazy = make([]int, 4*n)
	tr.sz = n
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			tr.val[i][0] = a[l]
			tr.val[i][1] = a[l]
			tr.sum[i] = 0
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		tr.merge(i)
	}
	build(0, 0, n-1)
	return tr
}

func (tr *Tree) update(i int, v int) {
	// sum no change, only value change
	tr.val[i][0] += v
	tr.val[i][1] += v
	tr.lazy[i] += v
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(2*i+1, tr.lazy[i])
		tr.update(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int, L int, R int)
	loop = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.update(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			loop(2*i+1, l, mid, L, R)
		} else if mid < L {
			loop(2*i+2, mid+1, r, L, R)
		} else {
			loop(2*i+1, l, mid, L, mid)
			loop(2*i+2, mid+1, r, mid+1, R)
		}
		tr.merge(i)
	}
	loop(0, 0, tr.sz-1, L, R)
}
