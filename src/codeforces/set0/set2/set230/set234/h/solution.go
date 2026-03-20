package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd, _ := os.Open("input.txt")
	defer rd.Close()
	wr, _ := os.Create("output.txt")
	defer wr.Close()
	reader := bufio.NewReader(rd)
	_, _, order, op := drive(reader)
	s := fmt.Sprintf("%v", order)

	fmt.Fprintln(wr, s[1:len(s)-1])
	fmt.Fprintln(wr, len(op))
	if len(op) > 0 {
		s = fmt.Sprintf("%v", op)
		fmt.Fprintln(wr, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (a []int, b []int, order []int, op []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	b = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	order, op = solve(a, b)
	return
}

func solve(a []int, b []int) (order []int, op []int) {
	n := len(a)
	m := len(b)

	// state = 0, face down, state = 1, face up
	play := func(a []int, b []int, state int) (order []int, op []int) {
		for i, j := 0, 0; i < n || j < m; {
			if (i == n || a[i] != state) && (j == m || b[j] != state) {
				// 必须反转一次
				op = append(op, i+j)
				state ^= 1
			}
			if i < n && a[i] == state {
				order = append(order, i+1)
				i++
			}
			if j < m && b[j] == state {
				order = append(order, n+j+1)
				j++
			}
		}
		if state == 1 {
			op = append(op, n+m)
		}
		return
	}
	o1, p1 := play(a, b, a[0])
	if b[0] != a[0] {
		o2, p2 := play(a, b, b[0])
		if len(p1) > len(p2) {
			o1, p1 = o2, p2
		}
	}
	return o1, p1
}
