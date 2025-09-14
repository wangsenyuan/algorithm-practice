package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int

	fmt.Fscan(reader, &n)
	res := solve(n)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func solve(n int) []int {

	check := func(x int, d int) bool {
		// x是奇数
		if x-1 > 2*n/x {
			return true
		}

		sum := x * (x - 1) / 2
		if sum >= n {
			return true
		}

		if d == 0 {
			return false
		}
		// m = x << d 但是有可能会溢出

		for i := d; i > 0; i-- {
			if sum >= n-x {
				return true
			}
			sum += x
			if x > n/2 && i > 0 {
				return true
			}
			x *= 2
		}
		// x is odd
		return sum >= n
	}

	check2 := func(x int, d int) bool {
		if (x - 1) > 2*n/x {
			// 这个需要刚好相同
			return false
		}
		sum := x * (x - 1) / 2
		if sum == n {
			return d == 0
		}
		if sum > n {
			return false
		}
		for i := d; i > 0; i-- {
			if sum > n-x {
				return false
			}
			sum += x
			x *= 2
		}

		return sum == n
	}

	x := sort.Search(n+1, func(x int) bool {
		if x == 0 {
			return false
		}
		if (x - 1) < 2*n/x {
			return x*(x-1)/2 >= n
		}
		return true
	})

	x /= 2

	var res []int

	for d := range 63 {
		// 要迭代知道出现奇数
		l, r := 0, x+1
		for l < r {
			mid := (l + r) >> 1
			if check(2*mid+1, d) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if check2(2*r+1, d) {
			res = append(res, (2*r+1)<<d)
		}
	}

	return res
}
