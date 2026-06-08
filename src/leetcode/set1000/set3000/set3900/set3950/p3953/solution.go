package p3953

import (
	"math/bits"
	"slices"
)

func maxScore(nums []int, maxVal int) int {
	mv := slices.Max(nums)
	mv = max(maxVal, mv)

	lpf := make([]int, mv+1)
	var primes []int
	for i := 2; i <= mv; i++ {
		if lpf[i] == 0 {
			primes = append(primes, i)
			lpf[i] = i
		}
		for _, p := range primes {
			if i*p > mv {
				break
			}
			lpf[i*p] = p
			if i%p == 0 {
				break
			}
		}
	}
	freq := make([]int, mv+1)
	for _, num := range nums {
		freq[num]++
	}

	f1 := slices.Clone(freq)

	for x := 2; x <= mv; x++ {
		for y := 2 * x; y <= mv; y += x {
			freq[x] += freq[y]
		}
	}

	// 如果选择nums[i] = 1, 那么可以做到不修改
	var best int
	if freq[1] > 0 {
		best = 1
	}

	freq[1] = len(nums)

	// else 修改一个
	for x := 2; x <= mv; x++ {
		var arr []int
		for x1 := x; x1 > 1; {
			f := lpf[x1]
			for x1%f == 0 {
				x1 /= f
			}
			arr = append(arr, f)
		}
		k := len(arr)
		// 计算有多少个gcd(x, ?) > 1
		var cnt int
		for mask := 1; mask < 1<<k; mask++ {
			prod := 1
			for i := range k {
				if (mask>>i)&1 == 1 {
					prod *= arr[i]
				}
			}

			if bits.OnesCount(uint(mask))%2 == 1 {
				cnt += freq[prod]
			} else {
				cnt -= freq[prod]
			}
		}

		if x <= maxVal {
			if f1[x] == 0 {
				best = max(best, x-max(cnt, 1))
			} else {
				// x已经存在了，那么有一个不删除
				best = max(best, x-(cnt-1))
			}
		} else if f1[x] > 0 {
			// 必须要有一个x
			best = max(best, x-(cnt-1))
		}
		// else f1[x] > 1, no way to use x
	}

	return best
}
