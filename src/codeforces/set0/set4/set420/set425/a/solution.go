package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) int {
	x := slices.Max(a)
	if x <= 0 {
		return x
	}
	res := x

	n := len(a)
	for r := range n {
		for l := range r {
			var pos []int
			for i := range l {
				if a[i] > 0 {
					pos = append(pos, a[i])
				}
			}
			for i := r + 1; i < n; i++ {
				if a[i] > 0 {
					pos = append(pos, a[i])
				}
			}
			slices.Sort(pos)
			slices.Reverse(pos)
			if len(pos) > k {
				pos = pos[:k]
			}

			dp := make([]int, len(pos)+1)
			ndp := make([]int, len(pos)+1)
			for i := l; i <= r; i++ {
				for j := range dp {
					ndp[j] = dp[j] + a[i]
				}
				if a[i] < 0 {
					for j := len(pos); j > 0; j-- {
						ndp[j] = max(ndp[j], dp[j-1]+pos[j-1])
					}
				}
				for j := range ndp {
					res = max(res, ndp[j])
					if ndp[j] < 0 {
						ndp[j] = 0
					}
				}
				copy(dp, ndp)
			}

		}
	}

	return res
}
