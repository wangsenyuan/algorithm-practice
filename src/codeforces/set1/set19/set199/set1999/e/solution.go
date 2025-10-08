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
		fmt.Fprintln(writer, res)
	}
}

const N = 2e5 + 10

var pref [N]int
var dp [N]int

func init() {
	dp[1] = 1
	dp[2] = 1
	pref[1] = 1
	pref[2] = 2
	for i := 3; i < N; i++ {
		dp[i] = dp[i/3] + 1
		pref[i] = pref[i-1] + dp[i]
	}
}

func drive(reader *bufio.Reader) int {
	var l, r int
	fmt.Fscan(reader, &l, &r)
	return solve(l, r)
}

func solve(l int, r int) int {
	if l+1 == r {
		x := dp[l]
		y := dp[r]
		return min(2*x+dp[r], 2*y+dp[l])
	}
	a := dp[l]*2 + pref[r] - pref[l]
	b := dp[l+1]*2 + pref[r] - pref[l+1] + dp[l]
	return min(a, b)
}
