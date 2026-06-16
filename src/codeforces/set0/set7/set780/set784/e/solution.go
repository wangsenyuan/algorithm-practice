package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	res := solve(a, b, c, d)
	fmt.Println(res)
}

func solve(a int, b int, c int, d int) int {
	w1 := Wire{a}
	w2 := Wire{b}
	w3 := Wire{c}
	w4 := Wire{d}

	x1 := OrGate{w1, w2}
	x2 := XorGate{w3, w4}
	x3 := AndGate{w2, w3}
	x4 := OrGate{w1, w4}

	y1 := AndGate{x1, x2}
	y2 := XorGate{x3, x4}

	o := OrGate{y1, y2}

	return o.pull()
}

type Gate interface {
	pull() int
}

type AndGate struct {
	left  Gate
	right Gate
}

func (and AndGate) pull() int {
	return and.left.pull() & and.right.pull()
}

type OrGate struct {
	left  Gate
	right Gate
}

func (or OrGate) pull() int {
	return or.left.pull() ^ or.right.pull()
}

type XorGate struct {
	left  Gate
	right Gate
}

func (xor XorGate) pull() int {
	return xor.left.pull() | xor.right.pull()
}

type Wire struct {
	val int
}

func (w Wire) pull() int {
	return w.val
}
