package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) (k int, c []int, res int) {
	var n int
	fmt.Fscan(reader, &n, &k)
	c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}
	res = solve(k, c)
	return
}

func solve(k int, c []int) int {
	n := len(c)
	sort.Ints(c)

	var sum int
	for _, v := range c {
		sum += v
	}

	check := func(x int) bool {
		// 将所有的数变成 >= x, 是否能够在k次完成
		var pref int
		var i int
		for i < n && c[i] < x {
			pref += c[i]
			i++
		}
		need := i*x - pref

		has := sum - pref - (n-i)*x

		return need <= min(k, has)
	}

	l, r := c[0], c[n-1]+1
	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	x := r - 1
	// 然后再找到diff

	check2 := func(diff int) bool {
		var pref int
		for i := 0; i < n && c[i] < x+diff; i++ {
			pref += x + diff - c[i]
		}
		var suf int

		for i := n - 1; i >= 0 && c[i] > x+diff; i-- {
			suf += c[i] - (x + diff)
		}

		return suf <= min(k, pref)
	}

	l = 0
	r = c[n-1] - x + 1

	for l < r {
		mid := (l + r) >> 1
		if check2(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func bruteForce(k int, c []int) int {
	n := len(c)
	for k > 0 {
		sort.Ints(c)
		if c[0] == c[n-1] {
			break
		}
		c[0]++
		c[n-1]--
		k--
	}
	sort.Ints(c)
	return c[n-1] - c[0]
}
