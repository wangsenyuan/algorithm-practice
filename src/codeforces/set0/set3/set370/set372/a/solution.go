package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	s := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &s[i])
	}
	return solve(s)
}

func solve(s []int) int {
	n := len(s)
	sort.Ints(s)

	slices.Reverse(s)

	var res int

	a := s[:n/2]
	b := s[n/2:]

	for i, j := 0, 0; i < len(a); i++ {
		for j < len(b) && b[j]*2 > a[i] {
			j++
		}
		if j < len(b) {
			res++
			j++
		}
	}

	return n - res
}
