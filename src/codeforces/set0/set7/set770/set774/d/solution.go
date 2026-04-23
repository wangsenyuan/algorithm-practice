package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("TRUTH")
	} else {
		fmt.Println("LIE")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, l, r int
	fmt.Fscan(reader, &n, &l, &r)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(l, r, a, b)
}

func solve(l int, r int, a []int, b []int) bool {
	l--
	r--

	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	for i := r + 1; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	n := len(a)
	freq := make([]int, n+1)
	for i := l; i <= r; i++ {
		freq[a[i]]++
		freq[b[i]]--
	}
	for i := 0; i <= n; i++ {
		if freq[i] != 0 {
			return false
		}
	}
	return true
}
