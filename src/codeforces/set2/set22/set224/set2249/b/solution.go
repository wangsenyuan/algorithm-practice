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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n-1)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const mod = 998244353

func mul(a int, b int) int {
	return a * b % mod
}

func solve(a []int) int {
	n := len(a) + 1
	count := make([]int, n)

	// The values on the left of n are prefix maxima, and the values on the
	// right are suffix maxima. Thus, after equal runs are compressed, a must
	// strictly rise once and then strictly fall; a value cannot occur in two
	// separate runs.
	seen := make([]bool, n)
	falling := false
	previous := 0
	for _, v := range a {
		if v <= 0 || v >= n {
			return 0
		}
		count[v]++
		if v == previous {
			continue
		}
		if seen[v] || (falling && v > previous) {
			return 0
		}
		seen[v] = true
		if previous != 0 && v < previous {
			falling = true
		}
		previous = v
	}

	// For every distinct value, its first occurrence is forced to be that
	// value. At each later occurrence, choose any still-unused smaller value.
	ans, used := 2, 0 // n and the first n-1 can be swapped.
	for value := 1; value < n; value++ {
		for occurrence := 0; occurrence < count[value]; occurrence++ {
			if occurrence > 0 {
				ans = mul(ans, value-used)
			}
			used++
		}
	}

	return ans
}
