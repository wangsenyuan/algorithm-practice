package main

import (
	"bufio"
	"fmt"
	"os"
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
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	var res int
	var carry int
	for i := range n {
		// 一个carry可以消耗两个
		x := min(carry, a[i]/2)
		carry -= x
		res += x
		a[i] -= x * 2
		// 剩下的3个一组
		res += a[i] / 3
		carry += a[i] % 3
	}

	return res
}
