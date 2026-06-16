package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var tc int
	fmt.Fscan(reader, &tc)

	res := make([]int, tc)
	for i := 0; i < tc; i++ {
		var n, r int
		fmt.Fscan(reader, &n, &r)
		seen := make([]int, n-1)
		for j := range seen {
			fmt.Fscan(reader, &seen[j])
		}
		res[i] = solve(n, r, seen)
	}
	return res
}

func solve(n int, r int, seen []int) int {
	sum := 0
	for _, x := range seen {
		sum += x
	}

	guess := (r - 1 - sum%n + n) % n
	if guess == 0 {
		return n
	}
	return guess
}
