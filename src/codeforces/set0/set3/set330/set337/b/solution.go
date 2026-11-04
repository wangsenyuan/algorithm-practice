package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var a, b, c, d int
	fmt.Fscan(reader, &a, &b, &c, &d)
	return solve(a, b, c, d)
}

type pair struct {
	first  int
	second int
}

func (p pair) lessEq(other pair) bool {
	return p.first*other.second <= p.second*other.first
}

func (p pair) mul(other pair) pair {
	a := p.first * other.first
	b := p.second * other.second
	g := gcd(a, b)
	return pair{a / g, b / g}
}

func (p pair) inv() pair {
	return pair{p.second, p.first}
}

func (p pair) div(other pair) pair {
	return p.mul(other.inv())
}

func (p pair) sub(other pair) pair {
	// a/c - b / d = (a*d - b*c) / (c*d)
	a := p.first*other.second - p.second*other.first
	if a == 0 {
		return pair{0, 1}
	}
	b := p.second * other.second
	g := gcd(a, b)
	return pair{a / g, b / g}
}

func solve(a int, b int, c int, d int) string {
	// 肯定有一个维度是完全匹配的
	// 如果水平维度匹配
	// a / d1 = c / d
	// a * d / c
	g := gcd(a*d, c)
	d1 := pair{a * d / g, c / g}
	var w pair
	if d1.lessEq(pair{b, 1}) {
		// a * b - c * d1
		w = pair{a, 1}.mul(d1)
	}
	// 如果垂直维度匹配
	// a1 / b = c / d
	a1 := pair{b, 1}.mul(pair{c, d})
	var v pair
	if a1.lessEq(pair{a, 1}) {
		v = pair{b, 1}.mul(a1)
	}

	var diff pair

	if v.second == 0 || w.second > 0 && v.lessEq(w) {
		diff = pair{a, 1}.mul(pair{b, 1}).sub(w)
	} else {
		diff = pair{a, 1}.mul(pair{b, 1}).sub(v)
	}

	diff = diff.div(pair{a, 1}.mul(pair{b, 1}))

	return fmt.Sprintf("%d/%d", diff.first, diff.second)
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
