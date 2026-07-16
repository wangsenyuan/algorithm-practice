package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := drive(reader)

	fmt.Println(len(res))

	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int) []int {
	marked := make([]int, n+1)

	for i := 1; i*i < n; i++ {
		for j := i + 1; i*i+j*j <= n; j++ {
			marked[i*i+j*j]++
		}
	}

	var res []int
	for i := 2; i <= n; i++ {
		if marked[i] == 1 {
			res = append(res, i)
		}
	}

	return res
}
