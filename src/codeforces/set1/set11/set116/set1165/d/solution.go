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
		ans := drive(reader)
		fmt.Fprintln(writer, ans)
	}
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
	slices.Sort(a)
	n := len(a)
	x := a[0] * a[n-1]

	for i := range n {
		if x%a[i] != 0 {
			return -1
		}
	}

	var d []int

	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			d = append(d, i)
			if i != x/i {
				d = append(d, x/i)
			}
		}
	}
	slices.Sort(d)
	if len(d) == len(a) {
		return x
	}
	return -1
}
