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
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) bool {
	n := len(a)

	play := func(toRem int) bool {
		prev := -1

		for i := range n {
			if i == toRem {
				continue
			}
			if a[i] > prev {
				prev = a[i]
			} else if a[i] == prev {
				prev = a[i] + 1
			} else {
				// a[i] < prev
				return false
			}
		}
		return true
	}
	// a[i] <= a[i+1] holds
	if play(0) {
		return true
	}
	prev := a[0]
	for i := 1; i < n; i++ {
		if a[i] < prev {
			return play(i-1) || play(i)
		}
		if a[i] == prev {
			prev++
		} else {
			// a[i] > prev
			prev = a[i]
		}
	}
	return true
}
