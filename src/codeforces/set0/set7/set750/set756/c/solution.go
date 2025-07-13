package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	res := process(reader)
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

func readString(reader *bufio.Reader) string {
	bs, _ := reader.ReadBytes('\n')
	return strings.TrimSpace(string(bs))
}

func process(reader *bufio.Reader) []int {
	m := readNum(reader)
	ops := make([]string, m)
	for i := range m {
		ops[i] = readString(reader)
	}
	return solve(ops)
}

func solve(ops []string) []int {
	m := len(ops)
	tr := NewTree(m)

	val := make([]int, m)

	ans := make([]int, m)

	for i := range m {
		buf := []byte(ops[i])
		var j, t int
		pos := readInt(buf, 0, &j) + 1
		pos = readInt(buf, pos, &t) + 1
		j--
		if t == 1 {
			var x int
			readInt(buf, pos, &x)
			val[j] = x
			tr.Update(j, 1)
		} else {
			tr.Update(j, -1)
		}
		active := tr.arr[0].open
		if active == 0 {
			ans[i] = -1
		} else {
			k := tr.GetTop()
			ans[i] = val[k]
		}
	}

	return ans
}

type data struct {
	close int
	open  int
}

func merge(l data, r data) data {
	if l.open >= r.close {
		return data{l.close, l.open - r.close + r.open}
	}
	// l.open < r.close
	return data{l.close + r.close - l.open, r.open}
}

type Tree struct {
	arr []data
}

func NewTree(n int) *Tree {
	arr := make([]data, 4*n)
	return &Tree{arr}
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			if v < 0 {
				tr.arr[i].close++
			} else {
				tr.arr[i].open++
			}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.arr[i] = merge(tr.arr[2*i+1], tr.arr[i*2+2])
	}

	n := len(tr.arr) / 4
	loop(0, 0, n-1)
}

func (tr *Tree) GetTop() int {
	var loop func(i int, l int, r int, dismissed int) int
	loop = func(i int, l int, r int, dismissed int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if tr.arr[2*i+2].open > dismissed {
			return loop(2*i+2, mid+1, r, dismissed)
		}

		tmp := tr.arr[2*i+2].close + dismissed - tr.arr[2*i+2].open
		return loop(2*i+1, l, mid, tmp)
	}

	return loop(0, 0, len(tr.arr)/4-1, 0)
}
