package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	ans := make([]int, t)
	for i := range t {
		var n, c int
		var s string
		fmt.Fscan(reader, &n, &c, &s)
		ans[i] = solve(n, c, s)
	}
	return ans
}

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve(n, c int, s string) int {
	if s[0] == '0' || s[n-1] == '0' {
		return -1
	}
	buf := []byte(s)
	buf[n-1] = '1'
	ans := 1

	countPow2 := func(x int) int {
		var p int
		for x%2 == 0 {
			p++
			x /= 2
		}
		return p
	}
	var p0 int
	p1 := countPow2(c)

	for i := 1; i < n; i++ {
		var f int
		switch buf[i] {
		case '1':
			f = 2
		case '0':
			f = i
		default:
			f = min(i, 2)
		}
		p0 += countPow2(f)
		c /= gcd(c, f)
		ans = mul(ans, f)
	}
	if c != 1 {
		return ans
	}

	r := p0 - p1 + 1
	ans = 1
	for i := 1; i < n; i++ {
		var f int
		switch buf[i] {
		case '1':
			f = 2
		case '0':
			f = i
		default:
			if i == 1 || i%2 == 0 {
				f = min(i, 2)
			} else {
				if r > 0 {
					f = i
					r--
				} else {
					f = 2
				}
			}

		}
		ans = mul(ans, f)
	}

	if r > 0 {
		return -1
	}

	return ans
}
