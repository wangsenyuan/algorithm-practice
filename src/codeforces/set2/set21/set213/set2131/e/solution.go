package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		if res {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) bool {
	n := len(a)

	check := func(i int) bool {
		if a[i] == b[i] {
			return true
		}
		if i < n-1 && b[i] == a[i]^a[i+1] {
			return true
		}
		if i < n-2 && b[i] == a[i]^b[i+1] {
			return true
		}
		return false
	}

	for i := range n {
		if !check(i) {
			return false
		}
	}
	return true
}
