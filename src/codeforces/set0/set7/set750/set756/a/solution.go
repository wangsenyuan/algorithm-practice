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
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(p, b)
}

func solve(p []int, b []int) int {
	// 要将p变成一个cycle
	// 1 2 3 4
	n := len(p)
	for i := range n {
		p[i]--
	}

	marked := make([]bool, n)
	var cnt int
	for i := range n {
		if !marked[i] {
			j := i
			for !marked[j] {
				marked[j] = true
				j = p[j]
			}
			cnt++
		}
	}
	// 1 3 2
	// 3 1 2
	var res int
	if cnt > 1 {
		res += cnt
	}

	var sum int
	for _, v := range b {
		sum += v
	}

	if sum%2 == 0 {
		res++
	}

	return res
}
