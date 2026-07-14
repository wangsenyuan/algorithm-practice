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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
	for i := range t {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		res[i] = solve(n, k)
	}
	return res
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

const N = 1001
const K = 51

var answers [N][K]int

func init() {
	// ways[k][sum] counts sets of k distinct positive interval lengths
	// with total length sum.
	var ways [K][N]int
	ways[0][0] = 1
	for k := 1; k < K; k++ {
		for sum := k; sum < N; sum++ {
			ways[k][sum] = add(ways[k][sum-k], ways[k-1][sum-k])
		}
	}

	fact := 1
	for k := 1; k < K; k++ {
		fact = mul(fact, k)
		cur := ways[k]
		// Applying prefix sums k+1 times convolves ways[k] with
		// C(n-sum+k, k), the number of ways to distribute empty spots.
		for rep := 0; rep <= k; rep++ {
			for n := 1; n < N; n++ {
				cur[n] = add(cur[n], cur[n-1])
			}
		}
		for n := 1; n < N; n++ {
			answers[n][k] = mul(cur[n], fact)
		}
	}
}

func solve(n, k int) int {
	if k >= K {
		return 0
	}
	return answers[n][k]
}
