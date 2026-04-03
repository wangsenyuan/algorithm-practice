package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ask := func(l int, r int) int {
		fmt.Println("?", l, r)
		var res int
		fmt.Fscan(reader, &res)
		return res
	}
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n, ask)
		fmt.Println("!", res)
	}
}

func solve(n int, ask func(l int, r int) int) int {

	// we know sure l...r has the maximum value
	var f func(l int, r int, sum int) int

	f = func(l int, r int, sum int) int {
		if l == r {
			return sum
		}
		sum /= 2
		// 寻找一中点mid, l...mid = mid+1...r
		mid := sort.Search(r-l+1, func(i int) bool {
			return ask(l+1, l+i+1) >= sum
		})
		mid += l
		if mid-l+1 < r-mid {
			return f(l, mid, sum)
		}
		return f(mid+1, r, sum)
	}

	s := ask(1, n)
	return f(0, n-1, s)
}
