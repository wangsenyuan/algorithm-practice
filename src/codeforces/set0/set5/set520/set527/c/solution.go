package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
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

func drive(reader *bufio.Reader) []int {
	w, h, n := readThreeNums(reader)
	splits := make([][]int, n)
	for i := range n {
		s := readString(reader)
		var v int
		readInt([]byte(s), 2, &v)
		if s[0] == 'H' {
			splits[i] = []int{1, v}
		} else {
			splits[i] = []int{2, v}
		}
	}
	return solve(w, h, splits)
}

func solve(W int, H int, splits [][]int) []int {
	n := len(splits)

	t1 := NewTree(W + 1)
	t2 := NewTree(H + 1)

	res := make([]int, n)

	for i, cur := range splits {
		if cur[0] == 1 {
			// horizontal splits
			h := t2.Update(cur[1])
			w := t1[0].v
			res[i] = w * h
		} else {
			w := t1.Update(cur[1])
			h := t2[0].v
			res[i] = h * w
		}
	}
	return res
}

type data struct {
	l   int
	r   int
	v   int
	set bool
}

func merge(l int, r int, a data, b data) data {
	v := max(a.v, b.v)
	l1, r1 := l, r
	if a.set {
		l1 = a.r
	}
	if b.set {
		r1 = b.l
	}
	v = max(v, r1-l1)

	if a.set && b.set {
		l, r = a.l, b.r
	} else if a.set {
		l, r = a.l, a.r
	} else {
		l, r = b.l, b.r
	}

	return data{l: l, r: r, v: v, set: a.set || b.set}
}

type Tree []data

func NewTree(n int) Tree {
	res := make(Tree, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			res[i] = data{l: l, r: r, v: 0, set: false}
			return
		}
		mid := (l + r) / 2
		f(i*2+1, l, mid)
		f(i*2+2, mid+1, r)
		res[i] = data{l: l, r: r, v: r - l, set: false}
	}
	f(0, 0, n-1)
	return res
}

func (t Tree) Update(p int) int {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			t[i].set = true
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		t[i] = merge(l, r, t[i*2+1], t[i*2+2])
	}
	n := len(t) / 4
	f(0, 0, n-1)
	return t[0].v
}
