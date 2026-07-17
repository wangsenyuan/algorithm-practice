package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	x, y, ok := drive(reader)
	if !ok {
		fmt.Println(-1)
		return
	}
	fmt.Println(x, y)
}

func drive(reader *bufio.Reader) (x, y int, ok bool) {
	var n int64
	fmt.Fscan(reader, &n)
	return solve(n)
}

func cubic(i int64) int64 {
	return i * i * i
}

func solve(n int64) (x, y int, ok bool) {

	for d := int64(1); cubic(d) <= n; d++ {
		// x - y = d
		if n%d == 0 {
			m := n / d
			k := play(3, 3*d, d*d-m)
			if k > 0 {
				return int(k + d), int(k), true
			}
		}
	}

	return 0, 0, false
}

func play(a int64, b int64, c int64) int64 {
	var l int64
	r := int64(600000001)
	for l+1 < r {
		mid := (l + r) >> 1
		if a*mid*mid+b*mid+c <= 0 {
			l = mid
		} else {
			r = mid
		}
	}

	if a*l*l+b*l+c == 0 {
		return l
	}
	return -1
}
