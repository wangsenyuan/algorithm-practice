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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, d, b int
	fmt.Fscan(reader, &n, &d, &b)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, b, d, a)
}

func solve(n int, b int, d int, a []int) int {
	a1 := slices.Clone(a)
	slices.Reverse(a1)

	// 从左往右，不超过mid的，最多移动w个学生的情况下，有多少个寝室 == b
	f := func(arr []int, mid int, w int) int {
		var res int
		var sum int
		var j int
		for i := range mid {
			r := min(i+(i+1)*d, n-1)
			for j <= r && sum < b {
				v := min(w, arr[j])
				sum += v
				w -= v
				j++
			}
			if sum >= b {
				res++
				sum -= b
			}
		}
		return mid - res
	}

	mid := (n + 1) / 2

	check := func(w int) bool {
		return f(a, mid, w) < f(a1, n-mid, n*b-w)
	}

	w := sort.Search(n*b, check)

	return max(f(a, mid, w-1), f(a1, n-mid, n*b-(w-1)))
}
