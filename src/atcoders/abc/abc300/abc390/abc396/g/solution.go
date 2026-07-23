package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var h, w int
	fmt.Fscan(reader, &h, &w)
	grid := make([]string, h)
	for i := range h {
		fmt.Fscan(reader, &grid[i])
	}
	return solve(w, grid)
}

func solve(w int, grid []string) int {
	dp := make([][]int, w+1)
	for i := range w + 1 {
		dp[i] = make([]int, 1<<w)
	}
	for _, row := range grid {
		var mask int
		for _, c := range row {
			mask = mask<<1 | int(c-'0')
		}
		dp[0][mask]++
	}

	for i := range w {
		for j := i; j >= 0; j-- {
			for bit := range 1 << w {
				dp[j+1][bit^(1<<i)] += dp[j][bit]
			}
		}
	}
	ans := 1 << 60
	for bit := range 1 << w {
		var sum int
		for i := range w + 1 {
			sum += min(i, w-i) * dp[i][bit]
		}
		ans = min(ans, sum)
	}
	return ans
}

func solve1(w int, grid []string) int {
	n := 1 << w
	freq := make([]int, n)
	for _, row := range grid {
		var mask int
		for _, c := range row {
			mask = mask<<1 | int(c-'0')
		}
		freq[mask]++
	}

	cost := make([]int, n)
	for mask := range n {
		ones := bits.OnesCount(uint(mask))
		cost[mask] = min(ones, w-ones)
	}

	fwht(freq, false)
	fwht(cost, false)
	for i := range n {
		freq[i] *= cost[i]
	}
	fwht(freq, true)

	return slices.Min(freq)
}

func fwht(a []int, inverse bool) {
	for length := 1; length < len(a); length <<= 1 {
		for start := 0; start < len(a); start += 2 * length {
			for i := 0; i < length; i++ {
				x := a[start+i]
				y := a[start+length+i]
				a[start+i] = x + y
				a[start+length+i] = x - y
			}
		}
	}
	if inverse {
		for i := range a {
			a[i] /= len(a)
		}
	}
}
