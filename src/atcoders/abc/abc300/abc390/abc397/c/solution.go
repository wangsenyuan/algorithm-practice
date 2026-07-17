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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	suf := make([]int, n+1)
	var sufCount int
	for _, v := range a {
		suf[v]++
		if suf[v] == 1 {
			sufCount++
		}
	}

	pref := make([]int, n+1)
	var prefCount int

	var res int

	for _, v := range a {
		suf[v]--
		if suf[v] == 0 {
			sufCount--
		}

		pref[v]++
		if pref[v] == 1 {
			prefCount++
		}

		res = max(res, prefCount+sufCount)
	}

	return res
}
