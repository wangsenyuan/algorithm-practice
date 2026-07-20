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
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	h := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &h[i])
	}
	return solve(h)
}

func solve(h []int) []int {
	n := len(h)
	H := make([]int, n*2)
	copy(H, h)
	copy(H[n:], h)

	mh := slices.Max(H)
	first := slices.Index(H, mh)
	dp := make([]int, 2*n)
	stack := make([]int, n)
	var top int
	stack[top] = first
	top++
	for i := first + 1; i < first+n; i++ {
		for top > 0 && H[stack[top-1]] < H[i] {
			top--
		}
		dp[i] = dp[stack[top-1]] + H[i]*(i-stack[top-1])
		stack[top] = i
		top++
	}

	last := first + n
	top = 0
	stack[top] = last
	top++
	fp := make([]int, n*2)
	for i := last - 1; i > first; i-- {
		for top > 0 && H[stack[top-1]] < H[i] {
			top--
		}
		fp[i] = fp[stack[top-1]] + H[i]*(stack[top-1]-i)
		stack[top] = i
		top++
	}

	ans := make([]int, n)

	for i := first; i <= last; i++ {
		if i > first {
			ans[i%n] += dp[i-1]
		}
		ans[i%n] += fp[i]
	}

	return ans
}
