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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}

}

func drive(reader *bufio.Reader) int {
	var n, h int
	fmt.Fscan(reader, &n, &h)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(h, a)
}

type pair struct {
	first  int
	second int
}

func solve(h int, a []int) int {
	n := len(a)

	// dp[i] 是从i开始的左边的阶梯（递减）
	play := func(a []int) []int {
		dp := make([]int, n)
		var stack []pair
		for i := range n {
			for len(stack) > 0 && a[top(stack).first] <= a[i] {
				stack = stack[:len(stack)-1]
			}

			w := (h - a[i]) * (i + 1)
			if len(stack) > 0 {
				p := top(stack)
				w = p.second + (h-a[i])*(i-p.first)
			}
			stack = append(stack, pair{i, w})
			dp[i] = w
		}
		return dp
	}

	dp := play(a)
	slices.Reverse(a)
	fp := play(a)
	slices.Reverse(a)
	slices.Reverse(fp)

	var best int

	for l := range n {
		hi := l
		for r := l; r < n; r++ {
			if l == r {
				tmp := dp[l] + fp[r] - (h - a[l])
				best = max(best, tmp)
			} else {
				if a[hi] < a[r] {
					hi = r
				}
				tmp := dp[l] + fp[r]

				if hi != l {
					tmp += fp[l] - fp[hi] - (h - a[l])
				}
				if hi != r {
					tmp += dp[r] - dp[hi] - (h - a[r])
				}

				if hi != l && hi != r {
					tmp += (h - a[hi])
				}

				best = max(best, tmp)
			}
		}
	}

	return best
}

func top[T any](stack []T) T {
	return stack[len(stack)-1]
}
