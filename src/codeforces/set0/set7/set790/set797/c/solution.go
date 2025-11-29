package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) string {
	n := len(s)

	tr := NewSegTree(n)

	for i := range n {
		x := int(s[i] - 'a')
		tr.Update(i, x)
	}

	stack := make([]int, n)
	var top int
	var buf []byte

	hasBetter := func(j int) bool {
		v := tr.Get(j+1, n)
		return v < int(s[j]-'a')
	}

	for i := range n {
		stack[top] = i
		top++
		for top > 0 && !hasBetter(stack[top-1]) {
			j := stack[top-1]
			buf = append(buf, s[j])
			top--
			tr.Update(j, inf)
		}
	}
	return string(buf)
}

const inf = 1 << 60

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = inf
	}
	return SegTree(arr)
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v

	for p > 1 {
		tr[p>>1] = min(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	var res int = inf
	for l < r {
		if l&1 == 1 {
			res = min(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
