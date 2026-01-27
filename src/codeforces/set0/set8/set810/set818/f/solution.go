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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n)
		fmt.Fprintln(writer, res)
	}
}

func solve(n int) int {
	if n == 1 {
		return 0
	}
	// x是桥的数量
	check := func(x int) int {
		// 如果有x座桥
		if n > 2*x {
			return 0
		}
		// n <= 2 * x
		y := n - x
		cnt := y * (y - 1) / 2
		// cnt >= x
		return x + min(cnt, x)
	}

	l, r := n/2, n-1
	for r-l+1 > 3 {
		m1 := l + (r-l)/3
		m2 := r - (r-l)/3
		c1 := check(m1)
		c2 := check(m2)
		if c1 >= c2 {
			r = m2
		} else {
			l = m1
		}
	}
	var res int
	for i := l; i <= r; i++ {
		res = max(res, check(i))
	}
	return res
}
