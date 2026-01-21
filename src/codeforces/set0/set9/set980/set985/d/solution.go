package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, H int
	fmt.Fscan(reader, &n, &H)
	return solve(n, H)
}

func compare(x, y int, n int) int {
	if x == 0 || y == 0 {
		if n == 0 {
			return 0
		}
		return -1
	}
	// x * y >= n
	if x > n/y || y > n/x || x*y > n {
		return 1
	}
	if x*y < n {
		return -1
	}
	return 0
}

func solve(n int, H int) int {

	x := sort.Search(n+1, func(x int) bool {
		if x%2 == 0 {
			return x/2 > n/(x+1) || x/2*(x+1) >= n
		}
		return (x+1)/2 > n/x || (x+1)/2*x >= n
	})

	if x <= H {
		// x, x - 1, .... 1, 这个是最短的
		return x
	}
	if compare(x, x+1, n) == 1 {
		x--
	}
	// x > H

	check := func(w int) bool {
		x1 := min((H+w)/2, x)
		// x1是最大值
		s1 := x1 * (x1 + 1) / 2
		// s2 := n - s1
		// H, H + 1, .... x1, x1, .... x1 - 1, ..... 1
		if x1-H+x1 > w {
			return false
		}

		// s2 - x1 >= n
		if compare(H+x1, x1-H+1, n+x1) >= 0 {
			return true
		}

		s2 := (H+x1)*(x1-H+1)/2 - x1

		if s1+s2 >= n {
			return true
		}

		v := w - x1 - (x1 - H)

		return compare(v, x1, n-s1-s2) >= 0
	}
	l, r := x, n
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
