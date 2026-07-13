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

func drive(reader *bufio.Reader) string {
	var n int
	var p, q, r int64
	fmt.Fscan(reader, &n, &p, &q, &r)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(p, q, r, a)
}

func solve(p, q, r int64, a []int) string {

	n := len(a)
	dp := make([]int, n+1)
	for i := range n + 1 {
		dp[i] = 1
	}
	pref := make(map[int64]int)

	play := func(d int, w int64) {
		var sum int64
		pref[0] = 0
		for i, v := range a {
			sum += int64(v)
			if (dp[i+1]>>d)&1 == 1 {
				if j, ok := pref[sum-w]; ok {
					dp[j] |= 1 << (d + 1)
				}
			}
			pref[sum] = i + 1
		}
		clear(pref)
	}

	play(0, r)

	play(1, q)

	play(2, p)

	for i := range n {
		if (dp[i]>>3)&1 == 1 {
			return "Yes"
		}
	}

	return "No"
}
