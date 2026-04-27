package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		res := solve(n, k)
		fmt.Fprintln(writer, res)
	}
}

const inf = 1 << 30

func solve(n int, k int) int {
	if k == 0 {
		return 0
	}
	w := bits.OnesCount(uint(n))

	if k >= 32 {
		return k - 1 + w
	}
	// m := bits.Len(uint(n))
	dp := make([][2]int, k+1)
	// dp[j][0/1] 进行操作j次后，后缀有carry 0/1的情况下，后缀中剩余1的个数
	for j := range k + 1 {
		dp[j][0] = inf
		dp[j][1] = inf
	}
	dp[0][0] = 0

	for num := n; num > 0; num >>= 1 {
		ndp := make([][2]int, k+1)
		for j := range k + 1 {
			ndp[j][0] = inf
			ndp[j][1] = inf
		}
		d := num & 1
		for j := range k + 1 {
			for c := range 2 {
				if dp[j][c] < inf {
					switch c + d {
					case 0:
						ndp[j][0] = min(ndp[j][0], dp[j][c])
					case 1:
						ndp[j][0] = min(ndp[j][0], dp[j][c]+1)
						if j+1 <= k {
							ndp[j+1][1] = min(ndp[j+1][1], dp[j][c])
						}
					default:
						// c + d == 2
						ndp[j][1] = min(ndp[j][1], dp[j][c])
					}
				}
			}
		}

		dp = ndp
	}

	best := inf
	for j := range k + 1 {
		best = min(best, min(dp[j][0], dp[j][1]+1))
	}

	return k + w - best
}
