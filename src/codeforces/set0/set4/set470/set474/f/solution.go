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
	fmt.Print(buf.String())
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
	q := readNum(reader)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	tr := NewTree(a)
	res := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0]-1, cur[1]-1
		res[i] = r - l + 1 - tr.Query(l, r).second
	}
	return res
}

type pair struct {
	first  int
	second int
}

func merge(a, b pair) pair {
	if a.first == b.first {
		return pair{a.first, a.second + b.second}
	}
	if a.first%b.first == 0 {
		return b
	}
	if b.first%a.first == 0 {
		return a
	}
	return pair{gcd(a.first, b.first), 0}
}

type Tree struct {
	val []pair
	sz  int
}

func NewTree(arr []int) *Tree {
	n := len(arr)
	val := make([]pair, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = pair{arr[l], 1}
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		val[i] = merge(val[2*i+1], val[2*i+2])
	}
	build(0, 0, n-1)

	return &Tree{val, n}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func (t *Tree) Query(L int, R int) pair {
	var loop func(i int, l int, r int, L int, R int) pair
	loop = func(i int, l int, r int, L int, R int) pair {
		if l == L && r == R {
			return t.val[i]
		}
		mid := (l + r) / 2
		if mid < L {
			return loop(2*i+2, mid+1, r, L, R)
		}
		if R <= mid {
			return loop(2*i+1, l, mid, L, R)
		}
		a := loop(2*i+1, l, mid, L, mid)
		b := loop(2*i+2, mid+1, r, mid+1, R)
		return merge(a, b)
	}
	return loop(0, 0, t.sz-1, L, R)
}
