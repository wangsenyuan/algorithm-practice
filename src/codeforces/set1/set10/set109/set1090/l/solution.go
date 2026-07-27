package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var t, n, a, b, k int
	fmt.Fscan(reader, &t, &n, &a, &b, &k)
	return solve(t, n, a, b, k)
}

func solve(t, n, a, b, k int) int {
	if k > n {
		return 0
	}
	// 如果一个学生无法完成k个课程, 没有必要让他参与
	check := func(w int) bool {
		n1 := (n + 1) / 2
		n2 := n / 2
		seats := n1*min(a, w) + n2*min(b, w)
		return seats/k >= w
	}

	// upper bound is t (students), not n (lectures)
	res := sort.Search(t+1, func(w int) bool {
		return !check(w)
	})

	return res - 1
}
