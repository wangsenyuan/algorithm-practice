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
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	return solve(a)
}

func solve(a []int) int {
	var k int
	for _, v := range a {
		k += v
	}
	n := len(a)
	w := k / n
	// 每个里面需要w个
	var res int
	var pref int
	for i, v := range a {
		pref += v
		res += abs(pref - w*(i+1))
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
