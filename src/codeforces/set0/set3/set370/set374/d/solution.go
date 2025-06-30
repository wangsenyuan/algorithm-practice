package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) string {
	n, m := readTwoNums(reader)
	a := readNNums(reader, m)
	events := make([]int, n)
	for i := range n {
		events[i] = readNum(reader)
	}
	return solve(a, events)
}

func solve(a []int, events []int) string {
	m := len(events)
	tr := NewTree(m)
	var buf []byte

	for _, v := range events {
		if v == -1 {
			// hit the table
			w := tr.cnt[0]
			k := sort.SearchInts(a, w)
			// a[i] >= w
			if k == len(a) || a[k] > w {
				k--
			}
			// a[k] <= w
			for i := k; i >= 0; i-- {
				j := tr.FindKthPosition(a[i])
				tr.Update(j, -1)
			}
		} else {
			buf = append(buf, byte('0'+v))
			tr.Update(len(buf)-1, 1)
		}
	}
	var res []byte

	for i, x := range buf {
		if tr.Get(i) == 1 {
			res = append(res, x)
		}
	}

	if len(res) == 0 {
		return "Poor stack!"
	}

	return string(res)
}

type Tree struct {
	cnt []int
}

func NewTree(n int) *Tree {
	return &Tree{cnt: make([]int, 4*n)}
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.cnt[i] += v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.cnt[i] = tr.cnt[2*i+1] + tr.cnt[2*i+2]
	}
	n := len(tr.cnt) / 4
	loop(0, 0, n-1)
}

func (tr *Tree) FindKthPosition(k int) int {
	var loop func(i int, l int, r int, k int) int
	loop = func(i int, l int, r int, k int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if k <= tr.cnt[2*i+1] {
			return loop(2*i+1, l, mid, k)
		}
		return loop(2*i+2, mid+1, r, k-tr.cnt[2*i+1])
	}
	n := len(tr.cnt) / 4
	return loop(0, 0, n-1, k)
}

func (tr *Tree) Get(p int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if l == r {
			return tr.cnt[i]
		}
		mid := (l + r) / 2
		if p <= mid {
			return loop(2*i+1, l, mid)
		}
		return loop(2*i+2, mid+1, r)
	}
	n := len(tr.cnt) / 4
	return loop(0, 0, n-1)
}
