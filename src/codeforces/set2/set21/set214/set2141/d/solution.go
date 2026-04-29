package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	n := len(a)
	slices.Sort(a)
	var sum int
	for _, v := range a {
		sum += v
	}
	avg := (sum + k) / n
	if avg < a[n-1] {
		return -1
	}
	// avg >= a[n-1]
	var score int
	for i := 1; i < n; i++ {
		diff := avg - a[i]
		if a[i] == a[0] {
			diff--
		}
		score += max(0, diff)
	}
	return score
}
