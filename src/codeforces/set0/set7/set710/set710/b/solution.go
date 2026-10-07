package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	x := make([]int, n)
	for i := range x {
		fmt.Fscan(reader, &x[i])
	}
	return solve(x)
}

func solve(x []int) int {
	slices.Sort(x)
	mid := len(x) / 2
	if 2*mid == len(x) {
		return x[mid-1]
	}
	return x[mid]
}
