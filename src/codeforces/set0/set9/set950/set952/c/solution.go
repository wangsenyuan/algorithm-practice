package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) bool {
	// 不能产生相邻的，相差超过1的值
	n := len(a)
	for n > 1 {
		var j int
		for i := range n {
			if a[i] > a[j] {
				j = i
			}
			if i > 0 && abs(a[i]-a[i-1]) > 1 {
				return false
			}
		}

		for i := j + 1; i < n; i++ {
			a[i-1] = a[i]
		}
		n--
	}
	return true
}

func abs(num int) int {
	return max(num, -num)
}
