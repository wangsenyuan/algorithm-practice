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

func drive(reader *bufio.Reader) int64 {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int64 {
	seen := make(map[int]bool)
	var sub int64
	for _, v := range a {
		if v <= k && !seen[v] {
			seen[v] = true
			sub += int64(v)
		}
	}
	return int64(k)*int64(k+1)/2 - sub
}
