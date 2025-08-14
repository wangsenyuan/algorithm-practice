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
	var n, k, S int
	fmt.Fscan(reader, &n, &k, &S)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k, S)
}

const inf int = 1e16 + 10

func solve(a []int, k int, S int) int {
	n := len(a)
	h := n / 2
	dp1 := get(a[:h], k, S)
	dp2 := get(a[h:], k, S)
	var ans int

	for i := range k + 1 {
		for j := 0; j+i <= k; j++ {
			for s1, v1 := range dp1[i] {
				s2 := S - s1
				ans += v1 * dp2[j][s2]
			}
		}
	}

	return ans
}

func fact(x int) int {
	res := x
	for x > 1 {
		x--
		if res > inf/x {
			return inf
		}
		res *= x
	}
	return res
}

func get(a []int, k int, S int) []map[int]int {
	n := len(a)
	// n <= 8

	F := make([]int, n)
	for i := range n {
		F[i] = fact(a[i])
	}

	N := 1 << n

	create := func() []map[int]int {
		fp := make([]map[int]int, k+1)
		for i := range k + 1 {
			fp[i] = make(map[int]int)
		}
		return fp
	}

	dp := create()
	dp[0][0] = 1

	for state := 1; state < N; state++ {
		fp := create()
		fp[0][0] = 1

		for i := range n {
			if (state>>i)&1 == 1 {
				nfp := create()
				for j := 0; j <= k; j++ {
					for s, v := range fp[j] {
						if s+a[i] <= S {
							nfp[j][s+a[i]] += v
						}
						if s+F[i] <= S && j+1 <= k {
							nfp[j+1][s+F[i]] += v
						}
					}
				}
				fp = nfp
			}
		}

		for j, vs := range fp {
			for s, v := range vs {
				dp[j][s] += v
			}
		}
	}

	return dp
}
