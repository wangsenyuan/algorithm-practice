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
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, a, b int
	fmt.Fscan(reader, &n, &a, &b)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}
	return solve(n, a, b, x)
}

func solve(n int, a int, b int, x []int) []int {
	ans := make([]int, n)

	calc := func(v int, w int) int {
		return sort.Search(w, func(i int) bool {
			return (i*a)/b >= v
		})
	}

	for i, w := range x {
		// 能得到这么多的钱
		v := (w * a) / b
		// 得到v这么多的钱，需要 v * b + a
		// w1 * a / b = v
		w1 := calc(v, w)
		ans[i] = w - w1
	}

	return ans
}
