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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)

	for tc > 0 {
		tc--
		var n, k int
		fmt.Fscan(reader, &n, &k)
		res := solve(n, k)
		fmt.Fprintln(writer, res)
	}

}

var dp [1 << 10][2]int
var ndp [1 << 10][2]int

const inf = 11111111111111

func solve1(n int, k int) int {
	var digits []int
	for i := n; i > 0; i /= 10 {
		digits = append(digits, i%10)
	}

	slices.Reverse(digits)

	for mask := range 1 << 10 {
		for eq := range 2 {
			dp[mask][eq] = inf
			ndp[mask][eq] = inf
		}
	}

	var states []int
	for mask := range 1 << 10 {
		if bits.OnesCount(uint(mask)) <= k {
			states = append(states, mask)
		}
	}

	dp[0][1] = 0
	for _, d := range digits {
		for _, mask := range states {
			cur := dp[mask]
			for eq, v := range cur {
				if v < inf {
					for x := range 10 {
						if eq == 1 && x < d {
							continue
						}
						newMask := mask | (1 << x)
						nv := v*10 + x
						neq := eq
						if eq == 1 && x > d {
							neq = 0
						}
						if bits.OnesCount(uint(newMask)) > k {
							continue
						}
						ndp[newMask][neq] = min(ndp[newMask][neq], nv)
					}
				}
			}
		}

		for _, mask := range states {
			for eq := range 2 {
				dp[mask][eq] = ndp[mask][eq]
				ndp[mask][eq] = inf
			}
		}
	}
	best := inf
	for _, mask := range states {
		for _, v := range dp[mask] {
			if v < inf {
				best = min(best, v)
			}
		}
	}
	if best < inf {
		return best
	}
	return -1
}

func solve(n int, k int) int {
	if countDigits(n) <= k {
		return n
	}
	var arr []int
	for num := n; num > 0; num /= 10 {
		arr = append(arr, num%10)
	}

	reverse(arr)
	s := make(map[int]int)

	for {
		clear(s)
		for i := 0; i < len(arr); i++ {
			s[arr[i]]++
			if len(s) > k {
				for arr[i] == 9 {
					i--
				}
				arr[i]++
				for j := i + 1; j < len(arr); j++ {
					arr[j] = 0
				}
				break
			}
		}
		if len(s) <= k {
			return toNum(arr)
		}
	}
}

func countDigits(num int) int {
	res := make(map[int]int)
	for num > 0 {
		res[num%10]++
		num /= 10
	}
	return len(res)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func toNum(arr []int) int {
	var res int
	for _, num := range arr {
		res = res*10 + num
	}
	return res
}
