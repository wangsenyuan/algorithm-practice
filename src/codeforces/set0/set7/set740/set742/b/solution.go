package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, x)
}

func solve(a []int, x int) int {
	freq := make(map[int]int)
	var res int
	for _, v := range a {
		res += freq[v^x]
		freq[v]++
	}
	return res
}
