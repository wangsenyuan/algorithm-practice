package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)
	ans := solve(n, k)
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

const N = 51

var F [N]int

const inf = 1e18

func init() {
	F[0] = 1
	F[1] = 1
	for i := 2; i < N; i++ {
		F[i] = min(inf, F[i-1]+F[i-2])
	}
}

func solve(n int, k int) []int {
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if F[n-i-1] >= k {
			ans[i] = i + 1
		} else {
			// F[n-i] < k
			ans[i] = i + 2
			ans[i+1] = i + 1
			k -= F[n-i-1]
			i++
		}
	}

	return ans
}
