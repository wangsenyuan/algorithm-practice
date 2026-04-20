package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	var sum int
	for _, v := range a {
		sum += v
	}
	mx := (sum + n - 1) / n
	my := sum / n
	slices.Sort(a)
	slices.Reverse(a)
	// 前w个是mx,后面几个是my
	w := sum % n
	if w == 0 {
		w = n
	}
	var cnt int
	for i := range n {
		if i < w {
			cnt += abs(mx - a[i])
		} else {
			cnt += abs(my - a[i])
		}
	}
	return cnt / 2
}

func abs(num int) int {
	return max(num, -num)
}
