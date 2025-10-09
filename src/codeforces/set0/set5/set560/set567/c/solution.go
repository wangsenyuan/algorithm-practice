package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	freq := make(map[int]int)
	for _, v := range a {
		freq[v]++
	}
	// n := len(a)

	var res int
	pref := make(map[int]int)

	for _, v := range a {
		freq[v]--

		if abs(v)%k == 0 {
			res += pref[v/k] * freq[v*k]
		}

		pref[v]++
	}

	return res
}

func abs(x int) int {
	return max(x, -x)
}
