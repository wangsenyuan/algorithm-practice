package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res < 0 {
		fmt.Println("inf")
	} else {
		fmt.Println(res)
	}
}

func drive(reader *bufio.Reader) int {
	var b1, q, l, m int
	fmt.Fscan(reader, &b1, &q, &l, &m)
	a := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &a[i])
	}
	return solve(b1, q, l, a)
}

func solve(b1 int, q int, l int, a []int) int {
	slices.Sort(a)

	found := func(num int) bool {
		i := sort.SearchInts(a, num)
		return i < len(a) && a[i] == num
	}

	var res int
	if abs(b1) > l {
		return 0
	}
	if !found(b1) {
		res++
	}

	if b1 == 0 || q == 0 {
		if found(0) {
			return res
		}
		// 0, 可以被写出来，那么就可以写无数次
		return -1
	}
	if q == 1 {
		if res > 0 {
			return -1
		}
		return 0
	}
	if q == -1 {
		b2 := b1 * q
		if abs(b2) <= l && !found(b2) {
			res++
		}
		if res > 0 {
			return -1
		}
		return 0
	}
	// try all until l
	for num := b1 * q; abs(num) <= l; num *= q {
		if !found(num) {
			res++
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
