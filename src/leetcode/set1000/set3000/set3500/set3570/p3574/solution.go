package p3574

import "math/bits"

func maxGCDScore(nums []int, k int) int64 {
	n := len(nums)
	freq := make([]int, 32)
	var best int
	for l := 0; l < n; l++ {
		clear(freq)
		var g int
		for r := l; r < n; r++ {
			g = gcd(g, nums[r])
			tmp := g * (r - l + 1)
			freq[getPowerOf2(nums[r])]++
			f := getPowerOf2(g)
			w := freq[f]
			if w <= k {
				tmp *= 2
			}
			best = max(best, tmp)
		}
	}

	return int64(best)
}

func getPowerOf2(num int) int {
	return bits.TrailingZeros(uint(num))
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
